package lib

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"strings"

	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

type Zip struct {
}

func (z *Zip) Default(zipPath string) error {
	err := z.convertZipToUTF8(zipPath)
	if err != nil {
		return err
	}
	return nil
}

func (z *Zip) convertZipToUTF8(zipPath string) error {

	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	tmpZipPath := zipPath + ".tmp"
	outFile, err := os.Create(tmpZipPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}

		content, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			return err
		}

		utf8Content, err := z.convertToUTF8(content)
		if err != nil {
			return err
		}

		newFile, err := zipWriter.Create(f.Name)
		if err != nil {
			return err
		}

		_, err = newFile.Write(utf8Content)
		if err != nil {
			return err
		}
	}

	if err := zipWriter.Close(); err != nil {
		return err
	}
	outFile.Close()

	if err := os.Remove(zipPath); err != nil {
		return err
	}
	if err := os.Rename(tmpZipPath, zipPath); err != nil {
		return err
	}

	return nil
}

func (z *Zip) convertToUTF8(input []byte) ([]byte, error) {
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(input)
	if err != nil {
		return nil, err
	}

	enc := z.getEncoding(result.Charset)

	//fmt.Print(result.Charset)

	if enc == nil {
		return input, nil
	}

	decoder := enc.NewDecoder()
	reader := transform.NewReader(bytes.NewReader(input), decoder)
	return io.ReadAll(reader)
}

func (z *Zip) getEncoding(charset string) encoding.Encoding {
	switch strings.ToLower(charset) {
	case "utf-8", "utf8":
		return nil
	case "iso-8859-1", "latin1":
		return charmap.ISO8859_1
	case "iso-8859-2", "latin2":
		return charmap.ISO8859_2
	case "iso-8859-3", "latin3":
		return charmap.ISO8859_3
	case "iso-8859-4", "latin4":
		return charmap.ISO8859_4
	case "iso-8859-5", "cyrillic":
		return charmap.ISO8859_5
	case "iso-8859-6", "arabic":
		return charmap.ISO8859_6
	case "iso-8859-7", "greek":
		return charmap.ISO8859_7
	case "iso-8859-8", "hebrew":
		return charmap.ISO8859_8
	case "iso-8859-9", "latin5":
		return charmap.ISO8859_9
	case "iso-8859-10", "latin6":
		return charmap.ISO8859_10
	case "iso-8859-13", "latin7":
		return charmap.ISO8859_13
	case "iso-8859-14", "latin8":
		return charmap.ISO8859_14
	case "iso-8859-15", "latin9":
		return charmap.ISO8859_15
	case "iso-8859-16", "latin10":
		return charmap.ISO8859_16
	case "windows-1250":
		return charmap.Windows1250
	case "windows-1251":
		return charmap.Windows1251
	case "windows-1252":
		return charmap.Windows1252
	case "windows-1253":
		return charmap.Windows1253
	case "windows-1254":
		return charmap.Windows1254
	case "windows-1255":
		return charmap.Windows1255
	case "windows-1256":
		return charmap.Windows1256
	case "windows-1257":
		return charmap.Windows1257
	case "windows-1258":
		return charmap.Windows1258
	case "macintosh", "mac-roman":
		return charmap.Macintosh
	case "koi8-r":
		return charmap.KOI8R
	case "koi8-u":
		return charmap.KOI8U
	case "koi8-ru":
		return charmap.KOI8R
	case "gbk", "cp936":
		return simplifiedchinese.GBK
	case "gb18030":
		return simplifiedchinese.GB18030
	case "big5", "cp950":
		return traditionalchinese.Big5
	case "shift_jis", "sjis", "cp932":
		return japanese.ShiftJIS
	case "euc-jp":
		return japanese.EUCJP
	case "euc-kr":
		return korean.EUCKR
	case "utf-16le":
		return unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	case "utf-16be":
		return unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	default:
		return nil
	}
}

package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"strings"
)

func Request(method string, query map[string]string) (map[string]any, error) {
	client := &http.Client{}
	rawJson, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	var data = strings.NewReader(string(rawJson))
	req, err := http.NewRequest("POST", "https://api.subsource.net/api/"+method, data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("authority", "api.subsource.net")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//fmt.Print(string(body))

	ret := make(map[string]any)
	err = json.Unmarshal(body, &ret)

	if err != nil {
		return nil, err
	}

	return ret, nil
}

func DownloadFile(path, url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	filename := getFilenameFromHeader(resp)
	filename = strings.Replace(filename, "-[SubSource]", "", -1)

	out, err := os.Create(path + filename)
	if err != nil {
		return "", err
	}

	defer out.Close()
	_, err = io.Copy(out, resp.Body)

	return filename, err
}

func getFilenameFromHeader(resp *http.Response) string {
	contentDisp := resp.Header.Get("Content-Disposition")
	if contentDisp == "" {
		return ""
	}

	_, params, err := mime.ParseMediaType(contentDisp)
	if err != nil {
		return ""
	}

	return params["filename"]
}

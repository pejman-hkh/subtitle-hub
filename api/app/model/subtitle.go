package model

import (
	"fmt"
	"strconv"
	"strings"
	"subtitle/gorn"
	"subtitle/lib"
	"time"
)

type Subtitle struct {
	BaseModel
	MovieId    uint   `gorm:"index" json:"movie_id"`
	UserId     uint   `gorm:"index" json:"user_id"`
	SubId      uint   `gorm:"index" json:"sub_id"`
	Rate       string `gorm:"size:100" json:"rate"`
	Lang       string `gorm:"index;size:100" json:"lang"`
	Type       string `gorm:"index;size:100" json:"type"`
	Title      string `gorm:"size:255" json:"title"`
	User       string `gorm:"size:255" json:"user"`
	FileName   string `gorm:"size:255" json:"file_name"`
	Link       string `gorm:"index;size:1000" json:"link"`
	LinkName   string `gorm:"index;size:255" json:"link_name"`
	Comment    string `json:"comment"`
	Error      string `json:"-"`
	Downloaded uint8  `gorm:"index" json:"downloaded"`
}

func (s *Subtitle) Sub(movie string, lang string, id string) (map[string]any, error) {
	return lib.Request("getSub", map[string]string{"movie": movie, "lang": lang, "id": id})
}

func (s *Subtitle) DaemonDownloadSubs() {
	path := "./public/subtitles/"
	for {

		subtitles := []Subtitle{}
		gorn.DB.Where("downloaded = 2 and updated_at < NOW() - INTERVAL 120 MINUTE order by id desc").Limit(100).Find(&subtitles)
		for _, subtitle := range subtitles {
			subtitle.Downloaded = 0
			subtitle.Error = ""
			subtitle.Save(&subtitle)
		}

		subtitles = []Subtitle{}
		gorn.DB.Where("downloaded = 0").Limit(100).Find(&subtitles)

		for _, subtitle := range subtitles {

			idStr := strconv.Itoa(int(subtitle.SubId))
			langStr := strings.Replace(strings.ToLower(subtitle.Lang), "/", "_", -1)
			detail, err := s.Sub(subtitle.LinkName, langStr, idStr)
			if err != nil {
				subtitle.Error = err.Error()
				subtitle.Downloaded = 2
				subtitle.Save(&subtitle)
			}

			sub, ok := detail["sub"].(map[string]any)
			if ok {
				token, ok := sub["downloadToken"].(string)
				if ok {

					filename, err := lib.DownloadFile(path, "https://api.subsource.net/api/downloadSub/"+token)

					zip := lib.Zip{}
					zip.Default(path + filename)

					if err == nil {
						subtitle.FileName = filename
						subtitle.Downloaded = 1
						subtitle.Save(&subtitle)
					} else {

						subtitle.Error = err.Error()
						subtitle.Downloaded = 2
						subtitle.Save(&subtitle)

						fmt.Print(err)
					}
				}
			}
			time.Sleep(2 * time.Second)
		}
		time.Sleep(2 * time.Second)
	}
}

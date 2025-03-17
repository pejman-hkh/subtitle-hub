package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"subtitle/gorn"
	"subtitle/lib"
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	BaseModel
	Name      string     `gorm:"size:255" json:"name"`
	ImdbCode  string     `gorm:"index;size:255" json:"imdb_code"`
	LinkName  string     `gorm:"index;size:255" json:"link_name"`
	Poster    string     `json:"poster"`
	Year      uint       `gorm:"index" json:"year"`
	Type      string     `gorm:"index;size:50" json:"type"`
	SubId     uint       `gorm:"index" json:"sub_id"`
	Detailed  uint8      `gorm:"index" json:"detail"`
	Data      string     `json:"data"`
	Subtitles []Subtitle `gorm:"foreignKey:MovieId;" json:"subtitles"`
	Seasons   []Season   `gorm:"foreignKey:MovieId;" json:"seasons"`
}

func (m *Movie) DaemonGetDetail() {
	for {
		movies := []Movie{}
		gorn.DB.Where("detailed = 0").Limit(100).Find(&movies)

		for _, movie := range movies {
			movie.GetDetail(movie.LinkName, "")
			movie.Detailed = 1
			movie.Save(&movie)

			time.Sleep(2 * time.Second)
		}
		time.Sleep(2 * time.Second)
	}
}

func (movie *Movie) GetSeasons() {
	data := make(map[string]any)
	json.Unmarshal([]byte(movie.Data), &data)

	seasons, ok := data["seasons"].([]any)
	if ok {
		//fmt.Print(seasons)
		for _, rawSeason := range seasons {
			seasonData, ok := rawSeason.(map[string]any)
			if ok {
				season := Season{}
				season.MovieId = movie.ID
				seasonNumber, ok := seasonData["number"].(float64)
				if ok && seasonNumber != 0 {
					gorn.DB.Where("movie_id = ? and season = ? ", movie.ID, uint(seasonNumber)).First(&season)
					season.Season = uint(seasonNumber)

					season.Save(&season)
				}
			}
		}
	}

	gorn.DB.Preload("Seasons", func(db *gorm.DB) *gorm.DB {
		return db.Order("season asc")
	}).First(&movie)
}

func (m *Movie) Search(title string) ([]Movie, error) {
	searched, err := lib.Request("searchMovie", map[string]string{"query": title})
	if err != nil {
		return nil, err
	}

	foundArray, ok := searched["found"].([]any)
	if !ok {
		return nil, errors.New("found index not found")
	}

	movies := []Movie{}
	for _, rawData := range foundArray {
		data, ok := rawData.(map[string]any)
		if !ok {
			break
		}

		movie := Movie{}

		imdbCode := ""

		//fmt.Print("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", title)

		if strings.HasPrefix(title, "tt") {
			imdbCode = title
		} else {
			imdbCode, _ = data["imdb"].(string)
		}

		gorn.DB.Where("imdb_code = ?", imdbCode).First(&movie)

		jsonData, err := json.Marshal(rawData)
		if err == nil {
			movie.Data = string(jsonData)
		}

		movie.ImdbCode = imdbCode

		linkName, ok := data["linkName"].(string)
		if ok {
			movie.LinkName = linkName
		}

		poster, ok := data["poster"].(string)
		if ok && movie.Poster == "" {
			movie.Poster = poster
		}

		title, ok := data["title"].(string)
		if ok {
			movie.Name = title
		}

		typed, ok := data["type"].(string)
		if ok {
			movie.Type = typed
		}

		subId, ok := data["id"].(float64)
		if ok {
			movie.SubId = uint(subId)
		}

		year, ok := data["releaseYear"].(float64)
		if ok {
			movie.Year = uint(year)
		}

		movie.Save(&movie)
		movies = append(movies, movie)
	}

	return movies, nil
}

func (m *Movie) GetDetail(name string, season string) (map[string]any, error) {

	param := make(map[string]string)
	param["movieName"] = name
	param["langs"] = "[]"
	if season != "" {
		param["season"] = season
	}

	subtitle, err := lib.Request("getMovie", param)

	fmt.Print("aaaaaaaaaaaaaaa", param, subtitle)

	if err != nil {
		return nil, err
	}

	movie := Movie{}
	if season == "" {

		mv, ok := subtitle["movie"].(map[string]any)

		if !ok {
			return nil, errors.New("movie index not found")
		}

		if m.ID == 0 {
			imdbCode, ok := mv["imdbLink"].(string)
			if ok {
				gorn.DB.Where("imdb_code = ? ", imdbCode).First(&movie)
			}

			if movie.ID == 0 {
				movie.ImdbCode = imdbCode
				title, ok := mv["fullName"].(string)
				if ok {
					movie.Name = title
				}

				year, ok := mv["year"].(float64)
				if ok {
					movie.Year = uint(year)
				}

				poster, ok := mv["poster"].(string)
				if ok {
					movie.Poster = poster
				}

				typed, ok := mv["type"].(string)
				if ok {
					movie.Type = typed
				}

				subId, ok := mv["id"].(float64)
				if ok {
					movie.SubId = uint(subId)
				}

				movie.Save(&movie)
			}
		}
	}

	if m.ID != 0 {
		movie = *m
	}

	subs, ok := subtitle["subs"].([]any)
	if !ok {
		return nil, errors.New("subs index not found")
	}

	for _, rawData := range subs {
		data, ok := rawData.(map[string]any)
		if !ok {
			break
		}

		subtitle := Subtitle{}

		if season != "" {
			seasonNumber := gorn.Atoi(strings.Replace(season, "season-", "", -1))
			subtitle.Season = uint(seasonNumber)
		}

		subtitle.MovieId = movie.ID

		subId, ok := data["subId"].(float64)
		if ok {
			subtitle.SubId = uint(subId)
		}

		gorn.DB.Where("sub_id = ?", subtitle.SubId).First(&subtitle)

		commentary, ok := data["commentary"].(string)
		if ok {
			subtitle.Comment = commentary
		}

		lang, ok := data["lang"].(string)
		if ok {
			if lang != "English" && lang != "Farsi/Persian" {
				continue
			}

			subtitle.Lang = lang
		}

		linkName, ok := data["linkName"].(string)
		if ok {
			if movie.LinkName == "" {
				movie.LinkName = linkName
				movie.Save(&movie)
			}
			subtitle.LinkName = linkName
		}

		link, ok := data["fullLink"].(string)
		if ok {
			subtitle.Link = link
		}

		releaseName, ok := data["releaseName"].(string)
		if ok {
			subtitle.Title = releaseName
		}

		uploadedBy, ok := data["uploadedBy"].(string)
		if ok {
			subtitle.User = uploadedBy
		}

		uploadedById, ok := data["uploadedById"].(string)
		if ok {
			subtitle.UserId = uint(gorn.Atoi(uploadedById))
		}

		subtitle.Save(&subtitle)
	}
	return subtitle, nil
}

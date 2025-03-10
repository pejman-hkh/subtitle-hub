package model

import (
	"errors"
	"subtitle/gorn"
	"subtitle/lib"
	"time"
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
	Subtitles []Subtitle `gorm:"foreignKey:MovieId;" json:"subtitles"`
}

func (m *Movie) DaemonGetDetail() {
	for {
		movies := []Movie{}
		gorn.DB.Where("detailed = 0").Limit(100).Find(&movies)

		for _, movie := range movies {
			movie.Detail(movie.LinkName)
			movie.Detailed = 1
			movie.Save(&movie)

			time.Sleep(1 * time.Second)
		}
		time.Sleep(2 * time.Second)
	}
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
		imdbCode, ok := data["imdb"].(string)
		if !ok {
			continue
		}
		gorn.DB.Where("imdb_code = ?", imdbCode).First(&movie)

		movie.ImdbCode = imdbCode

		linkName, ok := data["linkName"].(string)
		if ok {
			movie.LinkName = linkName
		}

		poster, ok := data["poster"].(string)
		if ok {
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

func (m *Movie) Detail(name string) (map[string]any, error) {
	subtitle, err := lib.Request("getMovie", map[string]string{"movieName": name})

	if err != nil {
		return nil, err
	}

	mv, ok := subtitle["movie"].(map[string]any)

	if !ok {
		return nil, errors.New("movie index not found")
	}

	movie := Movie{}
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
	} else {
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

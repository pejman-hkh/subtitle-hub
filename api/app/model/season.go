package model

type Season struct {
	BaseModel
	MovieId   uint       `gorm:"index" json:"movie_id"`
	Title     string     `gorm:"size:255" json:"title"`
	Season    uint       `gorm:"index" json:"season"`
	Poster    string     `json:"poster"`
	Detailed  uint8      `gorm:"index" json:"detail"`
	Subtitles []Subtitle `gorm:"foreignKey:MovieId;references:MovieId" json:"subtitles"`
}

func (s *Season) GetDetail() {

}

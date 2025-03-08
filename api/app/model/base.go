package model

import (
	"subtitle/gorn"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (b *BaseModel) Init() {

}

func (a *BaseModel) Save(model any) *gorm.DB {
	ret := gorn.DB.Save(model)
	return ret
}

func (a *BaseModel) Delete(model any) *gorm.DB {
	ret := gorn.DB.Delete(model)
	return ret
}

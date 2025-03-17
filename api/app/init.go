package app

import (
	"subtitle/app/controller"
	"subtitle/app/model"
	"subtitle/gorn"

	"github.com/gin-gonic/gin"
)

func Init(g *gin.RouterGroup) {
	controller.InitMovie(g)
	controller.InitSubtitle(g)
}

func Seeds() {

}

func Migrations() {
	gorn.DB.AutoMigrate(model.Movie{})
	gorn.DB.AutoMigrate(model.Subtitle{})
	gorn.DB.AutoMigrate(model.Season{})
}

package main

import (
	"flag"
	"subtitle/app"
	"subtitle/app/middle"
	"subtitle/docs"
	"subtitle/gorn"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Daemons() {
	// subtitle := model.Subtitle{}
	// movie := model.Movie{}
	// go subtitle.DaemonDownloadSubs()
	// go movie.DaemonGetDetail()

}

// @title           Subtitle API
// @version         1.0
// @description     Subtitle api
// @contact.name   Pejman Hkh
// @contact.url    https://www.peji.ir
// @contact.email  pejman.hkh@gmail.com
// @host
// @BasePath  /api/v1
func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	var productionFlag bool
	flag.BoolVar(&productionFlag, "p", false, "production")
	flag.Parse()
	gr := gorn.Gorn{}
	gr.Init("./", productionFlag)

	r.Use(middle.Cors())

	g := r.Group("api/v1")
	app.Migrations()
	app.Seeds()
	app.Init(g)

	r.Static("/api/files", "./public")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	go Daemons()
	r.Run(":8083")
}

package controller

import (
	"subtitle/app/model"
	"subtitle/gorn"

	"github.com/gin-gonic/gin"
)

type SubtitleController struct {
	BaseController
}

func InitSubtitle(r *gin.RouterGroup) {
	index := SubtitleController{}
	index.InitRoutes(r)
}

func (c *SubtitleController) InitRoutes(r *gin.RouterGroup) {

	g := r.Group("subtitles")
	g.GET("/:id/download", c.Download)

}

// DetailSubtitles godoc
// @Summary      Download Subtitle
// @Description  Download Subtitle
// @Tags         Subtitles
// @Param        id	path	string	true	"ID"
// @Accept       json
// @Produce      json
// @Router       /subtitles/{id}/download [get]
func (s *SubtitleController) Download(ctx *gin.Context) {
	subtitle := model.Subtitle{}
	gorn.DB.Where("id = ?", ctx.Param("id")).First(&subtitle)
	filename, err := subtitle.Download()

	if err != nil {
		subtitle.Error = err.Error()
		subtitle.Downloaded = 2
		subtitle.Save(&subtitle)
		s.FlashError(ctx, "failed", map[string]any{"err": err.Error()})
		return
	} else {
		subtitle.FileName = filename
		subtitle.Downloaded = 1
		subtitle.Save(&subtitle)
	}

	s.FlashSuccess(ctx, "ok", map[string]any{"subtitle": subtitle})
}

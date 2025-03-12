package controller

import (
	"archive/zip"
	"io"
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
	g.GET("/:id/json", c.Json)

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

func readZipFile(zipPath string) (map[string]string, error) {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	files := make(map[string]string)

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return nil, err
		}
		content, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			return nil, err
		}

		files[f.FileInfo().Name()] = string(content)
	}

	return files, nil
}

// DetailSubtitles godoc
// @Summary      Json Subtitle
// @Description  Json Subtitle
// @Tags         Subtitles
// @Param        id	path	string	true	"ID"
// @Accept       json
// @Produce      json
// @Router       /subtitles/{id}/json [get]
func (s *SubtitleController) Json(ctx *gin.Context) {
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

	path := "./public/subtitles/"

	files, err := readZipFile(path + subtitle.FileName)
	if err != nil {
		s.FlashError(ctx, err.Error(), nil)
		return
	}

	s.FlashSuccess(ctx, "ok", map[string]any{"files": files})
}

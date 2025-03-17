package controller

import (
	"fmt"
	"strings"
	"subtitle/app/model"
	"subtitle/gorn"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MovieController struct {
	BaseController
}

func InitMovie(r *gin.RouterGroup) {
	index := MovieController{}
	index.InitRoutes(r)
}

func (c *MovieController) InitRoutes(r *gin.RouterGroup) {

	r.GET("", c.Index)
	r.GET("/movie/:link", c.DetailByLink)
	r.GET("/movie/:link/:season", c.Season)

	g := r.Group("movies")

	g.GET("", c.Index)
	g.GET("/index", c.Index)
	g.GET("/search", c.SearchMovie)
	g.GET("/detail", c.Detail)

}

// ListMovies godoc
// @Summary      List movies
// @Description  Get list of all movies
// @Tags         movies
// @Param        search    query     string  false  "Search"
// @Param        page    query     number  false  "Pagination"
// @Param        name    query     string  false  "Name"
// @Param        year    query     number  false  "Year"
// @Param        imdb_code    query     string  false  "Imdb Code"
// @Param        type    query     string  false  "Type"
// @Accept       json
// @Produce      json
// @Router       /movies [get]
func (c *MovieController) Index(ctx *gin.Context) {
	var p gorn.Paginator
	list := []model.Movie{}
	advancedSearch := map[string]string{"type": "=", "imdb_code": "=", "name": "like", "year": "="}
	search := []string{"name", "imdb_code", "year"}
	gorn.DB.Scopes(c.AdvancedSearch(ctx, &list, advancedSearch)).Scopes(c.Search(ctx, &list, search)).Scopes(p.Paginate(ctx, &list)).Order("id desc").Find(&list)

	c.FlashSuccess(ctx, "ok", map[string]any{"list": list, "pagination": p})
}

// SearchMovies godoc
// @Summary      Search movies
// @Description  Search in movies
// @Tags         movies
// @Param        q    query     string  false  "Search"
// @Accept       json
// @Produce      json
// @Router       /movies/search [get]
func (s *MovieController) SearchMovie(ctx *gin.Context) {
	movie := model.Movie{}
	search, err := movie.Search(ctx.Query("q"))
	if err != nil {
		s.FlashError(ctx, err.Error(), nil)
		return
	}

	s.FlashSuccess(ctx, "ok", map[string]any{"list": search})
}

// DetailMovies godoc
// @Summary      Details movie
// @Description  Get movie details
// @Tags         movies
// @Param        id	query	int	false	"ID"
// @Param        imdb	query	string	false	"IMDB Code"
// @Accept       json
// @Produce      json
// @Router       /movies/detail [get]
func (s *MovieController) Detail(ctx *gin.Context) {
	movie := model.Movie{}
	if ctx.Query("id") != "" {
		gorn.DB.Preload("Subtitles").Preload("Seasons").Where("id = ?", ctx.Query("id")).First(&movie)
	} else if ctx.Query("imdb") != "" {
		gorn.DB.Preload("Subtitles").Preload("Seasons").Where("imdb_code = ?", ctx.Query("imdb")).First(&movie)

		if movie.ID == 0 {
			search, err := movie.Search(ctx.Query("imdb"))

			if err != nil {
				s.FlashError(ctx, err.Error(), nil)
				return
			}

			fmt.Print(search)

			gorn.DB.Preload("Subtitles").Preload("Seasons").Where("imdb_code = ?", ctx.Query("imdb")).First(&movie)

		}
	}

	// detail, err := movie.Detail(ctx.Query("name"))

	// if err != nil {
	// 	s.FlashError(ctx, err.Error(), nil)
	// 	return
	// }

	movie.GetSeasons()

	s.FlashSuccess(ctx, "ok", map[string]any{"movie": movie})
}

// Movies Season godoc
// @Summary      Get Season
// @Description  Get Season
// @Tags         movies
// @Param        link	path	string	true	"Link"
// @Param        season	path	string	true	"Season"
// @Accept       json
// @Produce      json
// @Router       /movie/{link}/{season} [get]
func (s *MovieController) Season(ctx *gin.Context) {
	movie := model.Movie{}
	gorn.DB.Preload("Seasons", func(db *gorm.DB) *gorm.DB {
		return db.Order("season asc")
	}).Where("link_name = ?", ctx.Param("link")).First(&movie)

	season := model.Season{}
	seasonNumber := strings.Replace(ctx.Param("season"), "season-", "", -1)
	fmt.Print(seasonNumber)

	gorn.DB.Preload("Subtitles", func(db *gorm.DB) *gorm.DB {
		return db.Where("season = ?", seasonNumber)
	}).Where("movie_id = ? and season = ?", movie.ID, seasonNumber).First(&season)

	if season.Detailed == 0 {
		movie.GetDetail(movie.LinkName, ctx.Param("season"))
		season.Detailed = 1
		season.Save(&season)

		gorn.DB.Preload("Subtitles", func(db *gorm.DB) *gorm.DB {
			return db.Where("season = ?", seasonNumber)
		}).First(&season)

	}

	s.FlashSuccess(ctx, "ok", map[string]any{"movie": movie, "season": season})
}

// DetailMovies godoc
// @Summary      Details movie
// @Description  Get movie details
// @Tags         movies
// @Param        link	path	string	true	"Link"
// @Accept       json
// @Produce      json
// @Router       /movie/{link} [get]
func (s *MovieController) DetailByLink(ctx *gin.Context) {
	movie := model.Movie{}
	gorn.DB.Preload("Subtitles").Preload("Seasons", func(db *gorm.DB) *gorm.DB {
		return db.Order("season asc")
	}).Where("link_name = ?", ctx.Param("link")).First(&movie)

	if len(movie.Subtitles) == 0 && movie.Data == "" {
		movie.Search(movie.ImdbCode)
	}

	if movie.Detailed == 0 {
		go func() {
			movie.GetDetail(movie.LinkName, "")
			if movie.ID != 0 {
				movie.Detailed = 1
				movie.Save(&movie)
			}
		}()
	}

	movie.GetSeasons()

	s.FlashSuccess(ctx, "ok", map[string]any{"movie": movie})
}

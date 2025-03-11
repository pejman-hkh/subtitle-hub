package controller

import (
	"fmt"
	"subtitle/app/model"
	"subtitle/gorn"

	"github.com/gin-gonic/gin"
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
		gorn.DB.Preload("Subtitles").Where("id = ?", ctx.Query("id")).First(&movie)
	} else if ctx.Query("imdb") != "" {
		gorn.DB.Preload("Subtitles").Where("imdb_code = ?", ctx.Query("imdb")).First(&movie)

		if movie.ID == 0 {
			search, err := movie.Search(ctx.Query("imdb"))

			if err != nil {
				s.FlashError(ctx, err.Error(), nil)
				return
			}

			fmt.Print(search)

			gorn.DB.Preload("Subtitles").Where("imdb_code = ?", ctx.Query("imdb")).First(&movie)

		}
	}

	// detail, err := movie.Detail(ctx.Query("name"))

	// if err != nil {
	// 	s.FlashError(ctx, err.Error(), nil)
	// 	return
	// }

	s.FlashSuccess(ctx, "ok", map[string]any{"movie": movie})
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
	gorn.DB.Preload("Subtitles").Where("link_name = ?", ctx.Param("link")).First(&movie)

	if len(movie.Subtitles) == 0 && movie.Data == "" {
		movie.Search(movie.ImdbCode)
	}

	if movie.Detailed == 0 {
		go func() {
			movie.Detail(movie.LinkName)
			movie.Detailed = 1
			movie.Save(&movie)
		}()
	}

	s.FlashSuccess(ctx, "ok", map[string]any{"movie": movie})
}

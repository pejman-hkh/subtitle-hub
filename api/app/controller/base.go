package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BaseController struct {
}

func (b *BaseController) Flash(ctx *gin.Context, status uint8, msg string, data map[string]any) {
	ctx.JSON(http.StatusOK, gin.H{"status": status, "msg": msg, "data": data})
}

func (b *BaseController) FlashSuccess(ctx *gin.Context, msg string, data map[string]any) {
	b.Flash(ctx, 1, msg, data)
}

func (b *BaseController) FlashError(ctx *gin.Context, msg string, data map[string]any) {
	b.Flash(ctx, 0, msg, data)
}

func (c *BaseController) Search(ctx *gin.Context, list any, search []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if ctx.Query("search") == "" {
			return db
		}

		sql := ""
		bind := make(map[string]any)
		pre := ""
		for _, v := range search {
			sql += pre + "" + v + " like @" + v
			bind[v] = "%" + ctx.Query("search") + "%"
			pre = " or "
		}

		if sql != "" {
			return db.Where("("+sql+")", bind)
		}
		return db
	}
}

func (c *BaseController) AdvancedSearch(ctx *gin.Context, list any, asearch map[string]string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		sql := ""
		bind := make(map[string]any)
		pre := ""

		if ctx.Query("id") != "" {
			sql += pre + " id = @id "
			bind["id"] = ctx.Query("id")
			pre = " and "
		} else {

			for k, v := range asearch {
				search := ctx.Query(k)
				if search == "" {
					continue
				}

				if v == "like" {
					sql += pre + k + " like @" + k
					bind[k] = "%" + search + "%"
				} else if v == "=" {
					sp := strings.Split(search, ",")
					if len(sp) > 1 {
						sql += "("
						for k1, v1 := range sp {
							sql += pre + k + " = @" + k + strconv.Itoa(k1)
							bind[k+strconv.Itoa(k1)] = v1
							pre = " or "
						}
						sql += ")"
						pre = " and "
					} else {
						sql += pre + k + " = @" + k
						bind[k] = search
					}
				} else if v == ">" {
					sql += pre + k + " > @" + k
					bind[k] = search
				} else if v == ">=" {
					sql += pre + k + " >= @" + k
					bind[k] = search
				}
				pre = " and "
			}
		}

		if sql != "" {
			return db.Where(sql, bind)
		}
		return db
	}
}

package gorn

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Paginator struct {
	Start uint  `json:"start"`
	End   uint  `json:"end"`
	Count int64 `json:"count"`
	Size  int   `json:"size"`
	Next  uint  `json:"next"`
	Prev  uint  `json:"prev"`
	Page  uint  `json:"page"`
	From  uint  `json:"from"`
	To    uint  `json:"to"`
}

func (p *Paginator) Paginate(r *gin.Context, model any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if r.Query("excel") != "" || r.Query("nopage") != "" {
			return db
		}

		var totalRows int64

		page, _ := strconv.Atoi(r.Query("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(r.Query("page_size"))
		if pageSize <= 0 {
			pageSize = 20
		}

		db.Session(&gorm.Session{}).Model(model).Count(&totalRows)

		p.Page = uint(page)
		p.Count = totalRows
		p.Size = pageSize
		p.From = uint((page*pageSize)-pageSize) + 1
		p.To = uint(page * pageSize)
		p.Start = 1
		end := uint(math.Ceil(float64(totalRows) / float64(pageSize)))
		if end <= 0 {
			end = 1
		}
		p.End = end
		next := uint(page) + 1
		if next > end {
			next = end
		}

		p.Next = next
		prev := page - 1
		if prev <= 0 {
			prev = 1
		}

		p.Prev = uint(prev)

		offset := (page - 1) * pageSize

		return db.Offset(offset).Limit(pageSize)
	}
}

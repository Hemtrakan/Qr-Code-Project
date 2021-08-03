package utility

import (
	"gorm.io/gorm"
	"math"
)

// Paginator

type Param struct {
	DB      *gorm.DB
	Page    int
	Limit   int
	OrderBy []string
}

type Paginator struct {
	TotalRecord int64 `json:"total_record"`
	TotalPage   int `json:"total_page"`
	Offset      int `json:"offset"`
	Limit       int `json:"limit"`
	Page        int `json:"page"`
	PrevPage    int `json:"prev_page"`
	NextPage    int `json:"next_page"`
}

func Paging(p *Param, result interface{}) *Paginator {
	db := p.DB

	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 10
	}
	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}

	var paginator Paginator
	var count int64
	var offset int

	countRecords(db, result, &count)

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	db.Limit(p.Limit).Offset(offset).Find(result)

	paginator.TotalRecord = count
	paginator.Page = p.Page

	paginator.Offset = offset
	paginator.Limit = p.Limit
	if totalPage := int(math.Ceil(float64(count) / float64(p.Limit))); totalPage > 0 {
		paginator.TotalPage = totalPage
	} else {
		paginator.TotalPage = 1
	}
	if p.Page > 1 {
		paginator.PrevPage = p.Page - 1
	} else {
		paginator.PrevPage = p.Page
	}

	if p.Page == paginator.TotalPage {
		paginator.NextPage = p.Page
	} else {
		paginator.NextPage = p.Page + 1
	}
	return &paginator
}

func countRecords(db *gorm.DB, anyType interface{}, count *int64) {
	db.Model(anyType).Count(count)
}

// --
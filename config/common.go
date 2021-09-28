package config

import (
	"math"
	"math/rand"
	"time"
)

var PageInfoVerify = Rules{"Page": {NotEmpty()}, "Limit": {NotEmpty()}}

type PageInfo struct {
	Page    int    `json:"page" form:"page"`
	Limit   int    `json:"limit" form:"limit"`
	Search  string `json:"search" form:"search"`
	OrderBy int    `json:"order_by" form:"order_by"`
}

type PaginHelper struct {
	TotalPage int `json:"total_page"`
	PrevPage  int `json:"prev_page"`
	NextPage  int `json:"next_page"`
}

func GetPaginHelper(pageInfo PageInfo, totalRecord int64) PaginHelper {
	var p PaginHelper
	p.TotalPage = int(math.Ceil(float64(totalRecord) / float64(pageInfo.Limit)))
	if pageInfo.Page > 1 {
		p.PrevPage = pageInfo.Page - 1
	} else {
		p.PrevPage = pageInfo.Page
	}

	if pageInfo.Page == p.TotalPage {
		p.NextPage = pageInfo.Page
	} else {
		p.NextPage = pageInfo.Page + 1
	}
	return p
}
func GetRandomNumber() (radomnumber int) {
	rand.Seed(time.Now().UnixNano())
	// min := 10000000000000
	// max := 99999999999999
	randn := rand.Int()
	return randn
}

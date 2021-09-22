package config

var PageInfoVerify = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}

type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

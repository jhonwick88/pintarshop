package controllers

import (
	"net/http"
	"strconv"

	"github.com/jhonwick88/pintarshop/config"
	"github.com/jhonwick88/pintarshop/models"

	"github.com/gin-gonic/gin"
)

type CreateItemInput struct {
	Name          string `json:"name"  binding:"required"`
	Price         uint64 `json:"price"  binding:"required"`
	PriceOriginal uint64 `json:"price_original"  binding:"required"`
	Description   string `json:"description"`
	Stock         uint8  `json:"stock"  binding:"required"`
}
type UpdateItemInput struct {
	Name          string `json:"name"`
	Price         uint64 `json:"price"`
	PriceOriginal uint64 `json:"price_original"`
	Description   string `json:"description"`
	Stock         uint8  `json:"stock"`
}

func CreateItem(c *gin.Context) {
	var createItem CreateItemInput
	if err := c.ShouldBindJSON(&createItem); err != nil {
		config.FailWithMessage(err.Error(), c)
		return
	}
	item := models.Item{
		Name:          createItem.Name,
		Price:         createItem.Price,
		PriceOriginal: createItem.PriceOriginal,
		Description:   createItem.Description,
		Stock:         createItem.Stock}
	models.DB.Create(&item)
	config.OkWithData(item, c)
}

func FindCustomItems(c *gin.Context) {
	var pageInfo config.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if err := config.Verify(pageInfo, config.PageInfoVerify); err != nil {
		config.FailWithMessage(err.Error(), c)
		return
	}
	err, listItems, total := getItemsData(pageInfo)
	if err != nil {
		config.FailWithMessage("Fail bro", c)
	} else {
		paginHelper := config.GetPaginHelper(pageInfo, total)
		config.OkWithDetailed(config.PageResult{
			List:      listItems,
			Total:     total,
			Page:      pageInfo.Page,
			Limit:     pageInfo.Limit,
			TotalPage: paginHelper.TotalPage,
			PrevPage:  paginHelper.PrevPage,
			NextPage:  paginHelper.NextPage,
		}, "Success", c)
	}

}

func getItemsData(pageInfo config.PageInfo) (err error, list interface{}, total int64) {
	limit := pageInfo.Limit
	offset := pageInfo.Limit * (pageInfo.Page - 1)
	db := models.DB.Model(&models.Item{})
	db.Count(&total)
	var items []models.Item
	err = db.Limit(limit).Offset(offset).Where("name LIKE ?", "%"+pageInfo.Search+"%").Order("name asc").Find(&items).Error
	return err, items, total
}

func FindItems(c *gin.Context) {
	var items []models.Item
	//var total int64
	q := c.Query("q")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	paginator := config.Paging(&config.Param{
		DB:      models.DB,
		Page:    page,
		Limit:   limit,
		OrderBy: "name desc",
		Search:  config.SearchFilter{Column: "name", Keyword: q},
		ShowSQL: false,
	}, &items)
	//models.DB.Model(items).Count(&total)
	//vd := models.DB.Scopes(Paginate(page, limit)).Where("name LIKE ?", "%"+q+"%").Order("name asc").First(&items)
	//c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": items, "total": vd.RowsAffected, "all": total})
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": paginator})
}
func FindItem(c *gin.Context) {
	var item models.Item
	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		config.FailWithMessage("Record not found!", c)
		return
	}
	config.OkWithData(item, c)
}
func UpdateItem(c *gin.Context) {
	var item models.Item
	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		config.FailWithMessage("Record not found!", c)
	}

	var input UpdateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		config.FailWithMessage(err.Error(), c)
		return
	}

	models.DB.Model(&item).Updates(input)
	config.OkWithData(item, c)
}
func DeleteItem(c *gin.Context) {
	var item models.Item
	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		config.FailWithMessage("Record not found!", c)
		return
	}
	models.DB.Delete(&item)
	config.OkWithMessage("Succes Deleted", c)
}

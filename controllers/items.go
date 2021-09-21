package controllers

import (
	"net/http"
	"pintarshop/models"
	"strconv"

	"pintarshop/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		//page, _ := strconv.Atoi(c.Query("page"))

		if page == 0 {
			page = 1
		}

		//pageSize, _ := strconv.Atoi(c.Query("limit"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
func CreateItem(c *gin.Context) {
	var createItem CreateItemInput
	if err := c.ShouldBindJSON(&createItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": err.Error(), "data": ""})
		return
	}

	item := models.Item{
		Name:          createItem.Name,
		Price:         createItem.Price,
		PriceOriginal: createItem.PriceOriginal,
		Description:   createItem.Description,
		Stock:         createItem.Stock}
	models.DB.Create(&item)
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": item})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": item})

}
func UpdateItem(c *gin.Context) {
	var item models.Item
	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	var input UpdateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&item).Updates(input)
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": item})
}
func DeleteItem(c *gin.Context) {
	var item models.Item
	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&item)

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success deleted", "data": ""})
}

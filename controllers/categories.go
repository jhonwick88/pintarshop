package controllers

import (
	"github.com/jhonwick88/pintarshop/config"
	"github.com/jhonwick88/pintarshop/models"

	"github.com/gin-gonic/gin"
)

type CreateCategoryInput struct {
	Name        string `json:"name" binding:"required"`
	Image       string `json:"image"`
	Description string `json:"description"`
	ParentId    *uint  `json:"parent_id"`
}
type UpdateCategoryInput struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	ParentId    *uint  `json:"parent_id"`
}

func CreateCategory(c *gin.Context) {
	var createCategory CreateCategoryInput
	if err := c.ShouldBindJSON(&createCategory); err != nil {
		config.FailWithMessage(err.Error(), c)
	}
	category := models.ItemCategory{
		Name:        createCategory.Name,
		Image:       createCategory.Image,
		Description: createCategory.Description,
		ParentId:    createCategory.ParentId,
	}
	models.DB.Create(&category)
	config.OkWithData(category, c)
}
func FindCategories(c *gin.Context) {
	var pageInfo config.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if err := config.Verify(pageInfo, config.PageInfoVerify); err != nil {
		config.FailWithMessage(err.Error(), c)
		return
	}
	err, listItems, total := getCategoriesData(pageInfo)
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
func getCategoriesData(pageInfo config.PageInfo) (err error, list interface{}, total int64) {
	limit := pageInfo.Limit
	offset := pageInfo.Limit * (pageInfo.Page - 1)
	db := models.DB.Model(&models.ItemCategory{})
	db.Count(&total)
	var items []models.ItemCategory
	err = db.Limit(limit).Offset(offset).Where("name LIKE ?", "%"+pageInfo.Search+"%").Order("ID asc").Find(&items).Error
	return err, items, total
}
func FindCategory(c *gin.Context) {
	var item models.ItemCategory
	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		config.FailWithMessage("Record not found!", c)
		return
	}
	config.OkWithData(item, c)
}
func UpdateCategory(c *gin.Context) {
	var item models.ItemCategory
	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		config.FailWithMessage("Record not found!", c)
	}
	var input UpdateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		config.FailWithMessage(err.Error(), c)
		return
	}
	models.DB.Model(&item).Updates(input)
	config.OkWithData(item, c)
}
func DeleteCategory(c *gin.Context) {
	var item models.ItemCategory
	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		config.FailWithMessage("Record not found!", c)
		return
	}
	models.DB.Delete(&item)
	config.OkWithMessage("Succes Deleted", c)
}

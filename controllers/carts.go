package controllers

import (
	"github.com/jhonwick88/pintarshop/models"

	"github.com/jhonwick88/pintarshop/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type AddCartInput struct {
	OptionID   uint   `json:"option_id"`
	OptionName string `json:"option_name"`
	ItemID     uint   `json:"item_id" binding:"required"`
	Qty        uint8  `json:"qty" binding:"required"`
}
type SelectIds struct {
	Ids []uint `json:"ids" binding:"required"`
}

func AddToCart(c *gin.Context) {
	var carts AddCartInput
	if err := c.ShouldBindJSON(&carts); err != nil {
		config.FailWithMessage(err.Error(), c)
	}
	var item models.CartItem
	userid, _ := ExtractTokenID(c)
	if err := models.DB.Preload(clause.Associations).Where(map[string]interface{}{"user_id": userid, "item_id": carts.ItemID}).First(&item).Error; err != nil {
		cartsItem := models.CartItem{
			OptionID:   carts.OptionID,
			OptionName: carts.OptionName,
			ItemID:     carts.ItemID,
			Qty:        carts.Qty,
			UserID:     userid,
		}
		models.DB.Create(&cartsItem)
	} else {
		models.DB.Model(&item).UpdateColumn("qty", carts.Qty)
	}
	pageInfo := config.PageInfo{
		Page:  1,
		Limit: 10,
	}
	getListCart(pageInfo, c)
}
func ListCarts(c *gin.Context) {
	var pageInfo config.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if err := config.Verify(pageInfo, config.PageInfoVerify); err != nil {
		config.FailWithMessage(err.Error(), c)
		return
	}
	getListCart(pageInfo, c)
}
func getListCart(pageInfo config.PageInfo, c *gin.Context) {
	userid, _ := ExtractTokenID(c)
	//var userid uint = 1
	err, listItems, total := getCartsData(pageInfo, userid)
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
func getCartsData(pageInfo config.PageInfo, userId uint) (err error, list interface{}, total int64) {
	limit := pageInfo.Limit
	offset := pageInfo.Limit * (pageInfo.Page - 1)
	db := models.DB.Model(&models.CartItem{})
	db.Count(&total)
	var items []models.CartItem
	err = db.Preload(clause.Associations).Limit(limit).Offset(offset).Where("user_id", userId).Order("ID asc").Find(&items).Error
	return err, items, total
}
func RemoveCartbyIds(c *gin.Context) {
	var ids SelectIds
	if err := c.ShouldBindJSON(&ids); err != nil {
		config.FailWithMessage(err.Error(), c)
	}
	//var cartItems models.CartItem
	if err := models.DB.Delete(&models.CartItem{}, ids.Ids).Error; err != nil {
		config.FailWithMessage(err.Error(), c)
	}
	config.OkWithMessage("Selected item has removed", c)
}

func RemoveCart(c *gin.Context) {
	var item models.CartItem
	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		config.FailWithMessage("Record not found!", c)
		return
	}
	models.DB.Delete(&item)
	config.OkWithMessage("Cart Deleted", c)
}

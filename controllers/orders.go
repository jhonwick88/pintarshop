package controllers

import (
	"strconv"

	"github.com/jhonwick88/pintarshop/config"
	"github.com/jhonwick88/pintarshop/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CreateOrderInput struct {
	NoResi       string               `json:"no_resi"`
	OrderQty     uint8                `json:"order_qty" binding:"required"`
	SubTotal     uint64               `json:"sub_total" binding:"required"`
	Discount     uint64               `json:"discount"`
	Tax          uint64               `json:"tax"`
	Total        uint64               `json:"total" binding:"required"`
	Status       string               `json:"status"`
	UserID       uint                 `json:"user_id" binding:"required"`
	CustomerID   uint                 `json:"customer_id" binding:"required"`
	OrderDetails []models.OrderDetail `json:"order_details"`
	CartId       []uint               `json:"cart_id" binding:"required"`
}
type UpdateOrderInput struct {
	ID     uint   `json:"id"`
	NoResi string `json:"no_resi" gorm:"type:varchar(150)"`
	Status string `json:"status"`
}
type UpdateQuantityInput struct {
	ID     uint  `json:"id" binding:"required"`
	Qty    uint8 `json:"qty" binding:"required"`
	ItemID uint  `json:"item_id" binding:"required"`
}

//	//OrderDetails []models.OrderDetail `json:"order_details"`

func CreateOrder(c *gin.Context) {
	var createOrder CreateOrderInput
	if err := c.ShouldBindJSON(&createOrder); err != nil {
		config.FailWithMessage(err.Error(), c)
	}
	order := models.Order{
		OrderNumber:  strconv.Itoa(config.GetRandomNumber()),
		NoResi:       createOrder.NoResi,
		OrderQty:     createOrder.OrderQty,
		SubTotal:     createOrder.SubTotal,
		Discount:     createOrder.Discount,
		Tax:          createOrder.Tax,
		Total:        createOrder.Total,
		Status:       createOrder.Status,
		UserID:       createOrder.UserID,
		CustomerID:   createOrder.CustomerID,
		OrderDetails: createOrder.OrderDetails,
	}

	err := models.DB.Preload(clause.Associations).Create(&order).Error
	if err != nil {
		updateStock(createOrder.OrderDetails)
		models.DB.Delete(&models.CartItem{}, createOrder.CartId)
	}
	config.OkWithData(order, c)
}
func updateStock(orderDetails []models.OrderDetail) {
	for _, orderDetail := range orderDetails {
		var product models.Item
		models.DB.Where("id = ?", orderDetail.ItemID).First(&product)
		models.DB.Model(&product).UpdateColumn("stock", gorm.Expr("stock - ?", orderDetail.Qty))
	}
}

func FindOrders(c *gin.Context) {
	var pageInfo config.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if err := config.Verify(pageInfo, config.PageInfoVerify); err != nil {
		config.FailWithMessage(err.Error(), c)
		return
	}
	err, listItems, total := getOrdersData(pageInfo)
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
func getOrdersData(pageInfo config.PageInfo) (err error, list interface{}, total int64) {
	limit := pageInfo.Limit
	offset := pageInfo.Limit * (pageInfo.Page - 1)
	db := models.DB.Model(&models.Order{})
	db.Count(&total)
	var items []models.Order
	//Preload(clause.Associations)
	err = db.Preload(clause.Associations).Limit(limit).Offset(offset).Where("order_number LIKE ?", "%"+pageInfo.Search+"%").Order("ID asc").Find(&items).Error
	return err, items, total
}
func FindOrder(c *gin.Context) {
	var item models.Order
	if err := models.DB.Preload(clause.Associations).Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		config.FailWithMessage("Record not found!", c)
		return
	}
	config.OkWithData(item, c)
}
func UpdateOrder(c *gin.Context) {
	var item models.Order
	if err := models.DB.Preload(clause.Associations).Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		config.FailWithMessage("Record not found!", c)
	}
	//fmt.Println(item)
	var input UpdateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		config.FailWithMessage(err.Error(), c)
		return
	}

	models.DB.Model(&item).Updates(models.Order{
		NoResi: input.NoResi,
		Status: input.Status,
	})
	config.OkWithData(item, c)
}
func UpdateQuantity(c *gin.Context) {
	var updateQuantityInput UpdateQuantityInput
	if err := c.ShouldBindJSON(&updateQuantityInput); err != nil {
		config.FailWithMessage(err.Error(), c)
		return
	}
	var item models.Item
	models.DB.First(&item, updateQuantityInput.ItemID)
	if updateQuantityInput.Qty > item.Stock {
		config.FailWithMessage("Out of stock", c)
		return
	}
	var orderdetail models.OrderDetail
	models.DB.First(&orderdetail, updateQuantityInput.ID)
	err := models.DB.Model(&orderdetail).Updates(map[string]interface{}{"qty": updateQuantityInput.Qty, "total_price": gorm.Expr("price * ?", updateQuantityInput.Qty)}).Error
	if err != nil {
		config.FailWithMessage("Cannot update quantity", c)
		return
	}
	var orderUpdate models.Order
	models.DB.Preload(clause.Associations).First(&orderUpdate, orderdetail.OrderID)
	var sub_total uint64 = 0
	for _, orderItem := range orderUpdate.OrderDetails {
		sub_total += orderItem.TotalPrice
	}
	models.DB.Model(&orderUpdate).Updates(models.Order{
		SubTotal: uint64(sub_total),
		Total:    uint64(sub_total) - uint64(orderUpdate.Discount),
	})
	config.OkWithData(orderUpdate, c)
}
func DeleteOrder(c *gin.Context) {
	var item models.Order
	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		config.FailWithMessage("Record not found!", c)
		return
	}
	models.DB.Select("OrderDetails").Delete(&item)
	config.OkWithMessage("Succes Deleted", c)
}

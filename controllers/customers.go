package controllers

import (
	"pintarshop/config"
	"pintarshop/models"

	"github.com/gin-gonic/gin"
)

type CreateCustomerInput struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
type UpdateCustomerInput struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func CreateCustomer(c *gin.Context) {
	var createCustomer CreateCustomerInput
	if err := c.ShouldBindJSON(&createCustomer); err != nil {
		config.FailWithMessage(err.Error(), c)
	}
	customer := models.Customer{
		Name:    createCustomer.Name,
		Address: createCustomer.Address,
		Phone:   createCustomer.Phone,
	}
	models.DB.Create(&customer)
	config.OkWithData(customer, c)
}
func FindCustomers(c *gin.Context) {
	var pageInfo config.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if err := config.Verify(pageInfo, config.PageInfoVerify); err != nil {
		config.FailWithMessage(err.Error(), c)
		return
	}
	err, listItems, total := getCustomersData(pageInfo)
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
func getCustomersData(pageInfo config.PageInfo) (err error, list interface{}, total int64) {
	limit := pageInfo.Limit
	offset := pageInfo.Limit * (pageInfo.Page - 1)
	db := models.DB.Model(&models.Customer{})
	db.Count(&total)
	var customers []models.Customer
	err = db.Limit(limit).Offset(offset).Where("name LIKE ?", "%"+pageInfo.Search+"%").Order("ID asc").Find(&customers).Error
	return err, customers, total
}
func FindCustomer(c *gin.Context) {
	var customer models.Customer
	if err := models.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		config.FailWithMessage("Record not found!", c)
		return
	}
	config.OkWithData(customer, c)
}
func UpdateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := models.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		config.FailWithMessage("Record not found!", c)
	}

	var input UpdateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		config.FailWithMessage(err.Error(), c)
		return
	}

	models.DB.Model(&customer).Updates(input)
	config.OkWithData(customer, c)
}
func DeleteCustomer(c *gin.Context) {
	var item models.Customer
	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		config.FailWithMessage("Record not found!", c)
		return
	}
	models.DB.Delete(&item)
	config.OkWithMessage("Succes Deleted", c)
}

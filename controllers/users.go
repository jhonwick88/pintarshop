package controllers

import (
	"pintarshop/config"
	"pintarshop/models"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	NickName string `json:"nickname"`
	Email    string `json:"email" binding:"required"`
	Role     uint   `json:"role"`
}
type UpdateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
	Email    string `json:"email"`
	Role     uint   `json:"role"`
}

func CreateUser(c *gin.Context) {
	var createUser CreateUserInput
	if err := c.ShouldBindJSON(&createUser); err != nil {
		config.FailWithMessage(err.Error(), c)
	}
	user := models.User{
		Username: createUser.Username,
		Password: createUser.Password,
		NickName: createUser.NickName,
		Email:    createUser.Email,
		Role:     createUser.Role,
	}
	models.DB.Create(&user)
	config.OkWithData(user, c)
}
func FindUsers(c *gin.Context) {
	var pageInfo config.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if err := config.Verify(pageInfo, config.PageInfoVerify); err != nil {
		config.FailWithMessage(err.Error(), c)
		return
	}
	err, listItems, total := getUsersData(pageInfo)
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
func getUsersData(pageInfo config.PageInfo) (err error, list interface{}, total int64) {
	limit := pageInfo.Limit
	offset := pageInfo.Limit * (pageInfo.Page - 1)
	db := models.DB.Model(&models.User{})
	db.Count(&total)
	var users []models.User
	err = db.Limit(limit).Offset(offset).Where("username LIKE ?", "%"+pageInfo.Search+"%").Order("ID asc").Find(&users).Error
	return err, users, total
}
func FindUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		config.FailWithMessage("Record not found!", c)
		return
	}
	config.OkWithData(user, c)
}
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		config.FailWithMessage("Record not found!", c)
	}

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		config.FailWithMessage(err.Error(), c)
		return
	}

	models.DB.Model(&user).Updates(input)
	config.OkWithData(user, c)
}
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		config.FailWithMessage("Record not found!", c)
		return
	}
	models.DB.Delete(&user)
	config.OkWithMessage("Success Deleted", c)
}

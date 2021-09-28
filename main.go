package main

import (
	"net/http"
	"pintarshop/controllers"
	"pintarshop/models"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Api is ready"})
	})
	//items
	r.POST("/items", controllers.CreateItem)
	r.GET("/items", controllers.FindCustomItems)
	r.GET("/items/:id", controllers.FindItem)
	r.PATCH("/items/:id", controllers.UpdateItem)
	r.DELETE("/items/:id", controllers.DeleteItem)
	//categories
	r.POST("/categories", controllers.CreateCategory)
	r.GET("/categories", controllers.FindCategories)
	r.GET("/categories/:id", controllers.FindCategory)
	r.PATCH("/categories/:id", controllers.UpdateCategory)
	r.DELETE("/categories/:id", controllers.DeleteCategory)
	//Carts
	r.POST("/carts", controllers.AddToCart)
	r.GET("/carts", controllers.ListCarts)
	r.POST("/carts/remove-all", controllers.RemoveCartbyIds)
	r.DELETE("/carts/:id", controllers.RemoveCart)
	//Orders
	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders", controllers.FindOrders)
	r.GET("/orders/:id", controllers.FindOrder)
	r.PATCH("/orders/:id", controllers.UpdateOrder)
	r.DELETE("/orders/:id", controllers.DeleteOrder)
	//Customers
	r.POST("/customers", controllers.CreateCustomer)
	r.GET("/customers", controllers.FindCustomers)
	r.GET("/customers/:id", controllers.FindCustomer)
	r.PATCH("/customers/:id", controllers.UpdateCustomer)
	r.DELETE("/customers/:id", controllers.DeleteCustomer)
	//Users
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.FindUsers)
	r.GET("/users/:id", controllers.FindUser)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

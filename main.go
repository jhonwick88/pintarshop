package main

import (
	"log"
	"os"

	"github.com/jhonwick88/pintarshop/controllers"
	"github.com/jhonwick88/pintarshop/middlewares"
	"github.com/jhonwick88/pintarshop/models"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	a := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	models.ConnectDatabase(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	b := a.Group("/api/v1/auth")
	b.POST("/login", controllers.Login)
	b.POST("/register", controllers.Register)

	r := a.Group("/api/v1")
	r.Use(middlewares.SetMiddlewareAuthentication())
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
	a.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

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
	r.POST("/items", controllers.CreateItem)
	r.GET("/items", controllers.FindCustomItems)
	r.GET("/items/:id", controllers.FindItem)
	r.PATCH("/items/:id", controllers.UpdateItem)
	r.DELETE("/items/:id", controllers.DeleteItem)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

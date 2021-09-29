package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jhonwick88/pintarshop/config"
	"github.com/jhonwick88/pintarshop/controllers"
)

func SetMiddlewareAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := controllers.TokenValid(c)
		if err != nil {
			config.FailWithMessage("Unauthorizedsss", c)
			c.Abort()
			return
		}
		c.Next()
	}
}

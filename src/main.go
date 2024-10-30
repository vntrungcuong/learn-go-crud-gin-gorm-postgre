package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Ping pong successful",
			})
		})
	}
	router.Run(":8080")
}

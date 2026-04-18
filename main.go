package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/", getEvents)
	server.Run(":8000")
}


	func getEvents(ctx *gin.Context){
				ctx.JSON(200, gin.H{"key": "Value"})
	}


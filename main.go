package main

import (
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events")
	server.Run(":8000")
}


	func getEvents(ctx *gin.Context){
		events := models.GetAllEvents()
		ctx.JSON(200, events)
	}


	func createEvent(ctx *gin.Context){
		var event models.Event //var of type event
		err :=ctx.ShouldBindJSON(&event) //bind req payload to type, can accept partial

		if err!= nil {
			ctx.JSON(400, gin.H{"message": "could not parse request data"})
			return
		}

		event.ID  = 1
		event.UserID = 1
		ctx.JSON(200, gin.H{"message": "Event created", "event": event})
	}


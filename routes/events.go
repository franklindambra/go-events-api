package routes

import (
	"fmt"
	"strconv"
	"time"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)


	func getEvents(ctx *gin.Context){
		events, err := models.GetAllEvents()
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Could not getch events"})
		}
		ctx.JSON(200, events)
	}


	func getEvent(ctx *gin.Context){
		eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Could not parse event ID"})
		}
		event, err := models.GetEventByID(eventId)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Could not find event"})
		}
		ctx.JSON(200, event)
	}


	func createEvent(ctx *gin.Context){

		var event models.Event //var of type event
		err :=ctx.ShouldBindJSON(&event) //bind req payload to type, can accept partial where binding not enforced
		if err != nil {
			ctx.JSON(400, gin.H{"message": "could not parse request data"})
			return
		}
		//fields
		
		userId := ctx.GetInt64("userId")
		event.UserID = userId

		event.DateTime = time.Now()

		//
		err = event.Save()
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Could not save event"})
		}
		ctx.JSON(200, gin.H{"message": "Event created", "event": event})
	}



	func updateEvent (ctx *gin.Context){
		eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Could not parse event ID"})
		}
		userId := ctx.GetInt64("userId")
		event, err := models.GetEventByID(eventId)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Update: could not locate ID"})
		}

		if event.UserID != userId {
			ctx.JSON(401, gin.H{"message": "Not authorized to update event"})
			return
		}

		var updatedEvent models.Event
		err =ctx.ShouldBindJSON(&updatedEvent)
		if err != nil{
			ctx.JSON(500, gin.H{"message": "Bad request"})
		}
		updatedEvent.ID = eventId
		err = updatedEvent.Update()
		if err != nil{
			ctx.JSON(500, gin.H{"message": "Could not update event."})
			fmt.Println(err)
			return
		}
		ctx.JSON(200, gin.H{"message": "Updated succesfully"})
	}



	func deleteEvent(ctx *gin.Context){
		eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Could not parse event ID"})
		}
		userId := ctx.GetInt64("userId")
		event, err := models.GetEventByID(eventId)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Delete: could not locate ID"})
			return
		}
		if event.UserID != userId{
			ctx.JSON(401, gin.H{"message": "Not authorized to delete event"})
			return
		}
		err = event.Delete()
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Could not delete"})
			return
		}
		ctx.JSON(200, gin.H{"message": "Event succesfully deleted"})
	}

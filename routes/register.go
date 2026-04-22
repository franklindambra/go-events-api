package routes

import (
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)



func registerForEvent(ctx * gin.Context){
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Could not parse event ID"})
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		ctx.JSON(500, gin.H{"message": "Failed to fetch event"})
		return
	}
	err = event.Register(userId)

	if err != nil {
		ctx.JSON(500, gin.H{"message": "Failed to fetch event"})
		return
	}

	ctx.JSON(201, gin.H{"message": "Successfully registered"})
	

}

func cancelRegistration(ctx * gin.Context){
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Could not parse event ID"})
	}
	var event models.Event
	event.ID = eventId
	err = event.CancelRegistration(userId)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Could not cancel registration"})
	}

	ctx.JSON(200, gin.H{"message": "Canceled!"})
}
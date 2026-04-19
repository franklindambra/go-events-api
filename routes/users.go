package routes

import (
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)


func signup(ctx *gin.Context){
	var user models.User

	err :=ctx.ShouldBindJSON(&user) //bind req payload to type, can accept partial where binding not enforced

 	if err != nil {
			ctx.JSON(400, gin.H{"message": "could not parse request data"})
			return
		}

	err = user.Save()

	if err != nil {
		ctx.JSON(400, gin.H{"message": "could not save user"})
			return
	}

	ctx.JSON(201, gin.H{"message": "User created succesfully"})




}

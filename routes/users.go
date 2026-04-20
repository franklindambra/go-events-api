package routes

import (
	"fmt"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
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



func login (ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	 	if err != nil {
			ctx.JSON(400, gin.H{"message": "could not parse request data"})
			fmt.Println(err)
			return
		}

		err = user.ValidateCredentials()

		if err != nil {
			ctx.JSON(401, gin.H{"message": err.Error()})
			return
		}

		token, err := utils.GenerateToken(user.Email, user.ID)

		if err != nil {
			ctx.JSON(500, gin.H{"message": "could not authenticate user"})
		}

		ctx.JSON(200, gin.H{"message": "Login succesfull", "token": token})


}

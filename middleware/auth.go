package middleware

import (
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)


func Authenticate(ctx *gin.Context){
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken((token))

	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "Not authorized."}) 
		return
	}

	//attach user data to context
	ctx.Set("userId", userId)

	ctx.Next()
}
package controller

import (
	"github.com/aparnasukesh/RESTapi-GO/model"
	"github.com/gin-gonic/gin"
)

var userData = make(map[string]model.User)

func Routes(r *gin.Engine) {
	r.POST("/signup", signup)
	r.POST("/login", login)
}

func signup(ctx *gin.Context) {
	var userInput model.User

	if err := ctx.BindJSON(&userInput); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	userData[userInput.UserName] = userInput

	ctx.JSON(200, gin.H{
		"message":  "Signup completed",
		"redirect": "http://localhost:2000/login",
	})
}

func login(ctx *gin.Context) {
	var userInput model.User

	if err := ctx.BindJSON(&userInput); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if user, found := userData[userInput.UserName]; found && userInput.Password == user.Password {
		ctx.JSON(200, gin.H{
			"message": "Successfully logged in",
			"data": map[string]interface{}{
				"username": user.UserName,
				"email":    user.Email,
			},
		})
	} else {
		ctx.JSON(404, gin.H{
			"message": "User not found",
		})
	}
}

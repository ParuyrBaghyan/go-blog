package routes

import (
	"go-blog/models"
	"go-blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.SignupUser
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create new user: " + err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "New user created successfully ✅"})
}

func signin(context *gin.Context) {
	var user models.SigninUser
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data." + err.Error()})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successfully ✅", "token": token})
}

func getUserProfile(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	user, err := models.GetUserById(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data"+ err.Error()})
		return
	}

	context.JSON(http.StatusOK, user)

}

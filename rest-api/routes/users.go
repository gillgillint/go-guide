package routes

import (
	"fmt"
	"net/http"

	"github.com/gillgillint/rest-api/models"
	"github.com/gillgillint/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong!"})
		return
	}

	err = user.Save()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not signup user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Create success"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong!"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Login failed"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successfully", "token": token})
}

func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		return
	}

	context.JSON(http.StatusOK, users)
}

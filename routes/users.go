package routes

import (
	"net/http"

	"github.com/frevent/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create user due to incorrect data."})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Coud not create user."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}

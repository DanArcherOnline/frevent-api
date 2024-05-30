package routes

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/frevent/models"
	"github.com/frevent/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event."})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event."})
		return
	}

	userID, err := getUserIDFromToken(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not update event due to incorrect data."})
		return
	}
	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized action."})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not update event due to incorrect data."})
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully."})
}

func getUserIDFromToken(context *gin.Context) (int64, error) {
	verifiedToken, ok := context.Get("token")
	if !ok {
		return 0, errors.New("could not find token in context")
	}

	userID, err := utils.GetUserIDFromToken(verifiedToken.(*jwt.Token))
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event."})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event."})
		return
	}

	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find events."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	verifiedToken, ok := context.Get("token")
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Coud not create event."})
	}

	userID, err := utils.GetUserIDFromToken(verifiedToken.(*jwt.Token))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create event due to invalid token."})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create event due to incorrect data."})
		return
	}

	event.UserID = userID
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Coud not create event."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully."})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event."})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event."})
		return
	}

	userID, err := getUserIDFromToken(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not update event due to incorrect data."})
		return
	}
	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized action."})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})
}

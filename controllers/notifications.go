package controllers

import (
	"Heytel/database"
	"Heytel/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNotification(ctx *gin.Context) {
	var notification models.Notification

	if err := ctx.ShouldBindJSON(&notification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	record := db.Create(&notification)
	if record.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": notification,
	})
}

func GetNotification(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var notification models.Notification

	if result := db.Where("id = ?", ctx.Param("id")).First(&notification); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": notification,
	})
}

func GetNotificationByReceiver(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var notifications []models.Notification

	if result := db.Where("receiver_id = ? OR receiver_id IS NULL", ctx.Param("id")).Find(&notifications); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": notifications,
	})
}

func GetAllNotifications(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var notifications []models.Notification

	if result := db.Find(&notifications); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": notifications,
	})
}

type notificationUpdate struct {
	Type       string    `json:"type" binding:"required"`    // Type of notification: "warning", "info", "alert"
	Message    string    `json:"message" binding:"required"` // Message of notification
}

func UpdateNotification(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var notification models.Notification

	if err := db.Where("id = ?", ctx.Param("id")).First(&notification).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input discountUpdate
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&notification).Updates(input)

	ctx.JSON(http.StatusOK, gin.H{"data": notification})
}


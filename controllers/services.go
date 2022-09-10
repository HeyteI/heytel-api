package controllers

import (
	"Heytel/database"
	"Heytel/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateService(ctx *gin.Context) {
	var service models.Service

	if err := ctx.ShouldBindJSON(&service); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	record := db.Create(&service)
	if record.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":  service,
	})
}

func GetService(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var service models.Service

	if result := db.Where("id = ?", ctx.Param("id")).First(&service); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": service,
	})
}

func GetAllServices(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var services []models.Service

	if result := db.Find(&services); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": services,
	})
}

type serviceUpdate struct {
	Name string `json:"service_name" binding:"required"`
	Price int `json:"service_suggested_price" binding:"required"`
	Description string `json:"description"`
}

func UpdateService(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var service models.Service

	if err := db.Where("id = ?", ctx.Param("id")).First(&service).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input serviceUpdate
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&service).Updates(input)

	ctx.JSON(http.StatusOK, gin.H{"data": service})
}


package controllers

import (
	"Heytel/database"
	"Heytel/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDiscount(ctx *gin.Context) {
	var discount models.Discount

	if err := ctx.ShouldBindJSON(&discount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	record := db.Create(&discount)
	if record.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": discount,
	})
}

func GetDiscount(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var discount models.Discount

	if result := db.Where("id = ?", ctx.Param("id")).First(&discount); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": discount,
	})
}

func GetAllDiscounts(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var discounts []models.Discount

	if result := db.Find(&discounts); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": discounts,
	})
}

type discountUpdate struct {
	Name string `json:"name" binding:"required"`
	Value int `json:"value" binding:"required"`
	Active bool `json:"active" binding:"required"`
}

func UpdateDiscount(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var discount models.Discount

	if err := db.Where("id = ?", ctx.Param("id")).First(&discount).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input discountUpdate
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&discount).Updates(input)

	ctx.JSON(http.StatusOK, gin.H{"data": discount})
}


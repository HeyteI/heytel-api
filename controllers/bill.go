package controllers

import (
	"Heytel/database"
	"Heytel/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBill(ctx *gin.Context) {
	var bill models.Bill

	if err := ctx.ShouldBindJSON(&bill); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	record := db.Create(&bill)
	if record.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":  bill,
	})
}

func GetBill(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var bill models.Bill

	if result := db.Where("id = ?", ctx.Param("id")).First(&bill); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": bill,
	})
}

func GetBillByInvoiceId(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var bill models.Bill

	if result := db.Where("invoice_id = ?", ctx.Param("id")).First(&bill); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": bill,
	})
}
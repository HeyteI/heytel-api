package controllers

import (
	"Heytel/database"
	"Heytel/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateShift(ctx *gin.Context) {
	var shift models.Shift

	if err := ctx.ShouldBindJSON(&shift); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	record := db.Create(&shift)
	if record.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": shift,
	})
}

func GetShift(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var shift models.Shift

	if result := db.Where("id = ?", ctx.Param("id")).First(&shift); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": shift,
	})
}

func GetShiftsByEmployee(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var shift []models.Shift

	if result := db.Where("employee = ?", ctx.Param("id")).Find(&shift); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": shift,
	})
}

type shiftUpdate struct {
	Start *time.Time `json:"shift_start"`
	ShouldWork *time.Time `json:"work_time"`

	End *time.Time `json:"shift_end"`
	Worked *time.Time `json:"worked"`

	StartDescription string `json:"start_description"`
	EndDescription string `json:"end_description"`
}

func UpdateShift(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var shift models.Shift

	if err := db.Where("id = ?", ctx.Param("id")).First(&shift).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input shiftUpdate
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Start != nil {
		shift.Start = input.Start
	}
	if input.ShouldWork != nil {
		shift.ShouldWork = input.ShouldWork
	}
	if input.End != nil {
		shift.End = input.End
	}
	if input.Worked != nil {
		shift.Worked = input.Worked
	}
	if input.StartDescription != "" {
		shift.StartDescription = input.StartDescription
	}
	if input.EndDescription != "" {
		shift.EndDescription = input.EndDescription
	}

	if err := db.Save(&shift).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": shift})
}


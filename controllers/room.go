package controllers

import (
	"Heytel/database"
	"Heytel/models"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/secured/room

// CreateRoom godoc
// @Summary create room
// @Schemes
// @Description creates room
// @Tags room
// @Accept json
// @Produce json
// @Success 200 {json} {room_data}
// @Router /api/secured/room/ [post]
func CreateRoom(ctx *gin.Context) {
	var room models.Room
	if err := ctx.ShouldBindJSON(&room); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	log.Print(room)

	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	record := db.Create(&room)
	if record.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": room,
	})
}

// AllRooms godoc
// @Summary get all rooms
// @Schemes
// @Description get all saved rooms
// @Tags room
// @Accept json
// @Produce json
// @Success 200 {json} {list of room_data}
// @Router /api/secured/room/all [get]
func AllRooms(ctx *gin.Context) {

	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var rooms []models.Room

	if result := db.Find(&rooms); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": rooms,
	})
}

// GetRoom godoc
// @Summary get room by id
// @Schemes
// @Description gets room by id
// @Tags room
// @Accept json
// @Produce json
// @Success 200 {json} {room_data}
// @Router /api/secured/room/ [get]
func GetRoom(ctx *gin.Context) {

	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var room models.Room

	if result := db.Where("id = ?", ctx.Param("id")).First(&room); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": room,
	})
}

// AllRoomsByFloor godoc
// @Summary get all rooms by it's floor
// @Schemes
// @Description get all rooms by it's floor
// @Tags room
// @Accept json
// @Produce json
// @Success 200 {json} {list of room_data}
// @Router /api/secured/room/floor/ [get]
func AllRoomsByFloor(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	floor := ctx.Param("floor")

	var rooms []models.Room

	if result := db.Where("floor = ?", floor).Find(&rooms); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": rooms,
	})
}

// GetAllFloors godoc
// @Summary get count of floors
// @Schemes
// @Description get count of floors
// @Tags room
// @Accept json
// @Produce json
// @Success 200 {json} {count of floors}
// @Router /api/secured/room/floors/ [get]
func GetAllFloors(ctx *gin.Context) {
	log.Print(os.Getenv("ALL_FLOORS"))
	ctx.JSON(http.StatusOK, gin.H{
		"data": os.Getenv("ALL_FLOORS"),
	})
}

type roomUpdate struct {
	Number       string `json:"number"`
	Class        string `json:"class"`
	People_range string `json:"people"`
	Description  string `json:"description"`
	Title        string `json:"title"`
	Price        int    `json:"price"`
	Floor        int    `json:"floor"`
	Available    bool   `json:"available"`
	InvoiceId    string `json:"invoice_id" binding:"uuid"`
}

// UpdateRoom godoc
// @Summary update room
// @Schemes
// @Description update room data by it's id
// @Tags room
// @Accept json
// @Produce json
// @Success 200 {json} {room_data}
// @Router /api/secured/room/ [patch]
func UpdateRoom(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var data map[string]interface{}

	jsonData, err := ctx.GetRawData()
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		panic(err)
	}

	var room models.Room
	if err := db.Where("id = ?", ctx.Param("id")).First(&room).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	tx := db.Begin()
	defer tx.Rollback()

	for key, value := range data {
		tx.Model(&room).Update(key, value)
	}

	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{"data": room})
}

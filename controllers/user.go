package controllers

import (
	"Heytel/database"
	"Heytel/helpers"
	"Heytel/middlewares"
	"Heytel/models"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// @BasePath /api/user

// RegisterUser godoc
// @Summary create user
// @Schemes
// @Description create user account
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {json} {user_data}
// @Router /api/user/register [post]
func RegisterUser(ctx *gin.Context) {

	var data map[string]interface{}

	jsonData, err := ctx.GetRawData()
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		panic(err)
	}

	log.Print(data)

	if data["group"] != "customer" && ctx.GetHeader("authorization") != os.Getenv("OWNER_AUTHORIZATION_KEY") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You got no permissions to do that."})
		ctx.Abort()
		return
	}

	var user models.User
	// if err := ctx.ShouldBindJSON(&user); err != nil {
	// 	log.Print(err)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	ctx.Abort()
	// 	return
	// }

	mapstructure.Decode(data, &user)

	if data["group"] != "customer" {
		if err := user.HashPassword(user.Password); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
	}

	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	record := db.Create(&user)
	if record.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}

// GetUser godoc
// @Summary get user by username
// @Schemes
// @Description get user by username
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {json} {user_data}
// @Router /api/secured/user/ [get]
func GetUser(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var user models.User

	username := ctx.Param("username")
	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong parameters passed.",
		})
		ctx.Abort()
		return
	}

	if result := db.Where("username = ?", username).First(&user); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// GetUserById godoc
// @Summary get user by id
// @Schemes
// @Description get user by id
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {json} {user_data}
// @Router /api/secured/user/id/ [get]
func GetUserById(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var user models.User

	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong parameters passed.",
		})
		ctx.Abort()
		return
	}

	if result := db.Where("id = ?", id).First(&user); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

type userUpdate struct {
	Username    string     `json:"username" gorm:"unique"`
	Group       string     `json:"group"`
	Birthday    *time.Time `json:"birthday"`
	Phone       string     `json:"phone"`
	Description string     `json:"description"`
	Photo_url   string     `json:"photo"`
}

// UpdateUser godoc
// @Summary update user
// @Schemes
// @Description update user data by it's id
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {json} {user_data}
// @Router /api/secured/user/ [patch]
func UpdateUser(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var user models.User

	log.Print(ctx.Params)
	if err := db.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input userUpdate
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&user).Updates(input)

	ctx.JSON(http.StatusOK, gin.H{"data": input})
}

type loginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginUser godoc
// @Summary login user
// @Schemes
// @Description login to users account
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {json} {user_data}
// @Router /api/user/login [post]
func LoginUser(ctx *gin.Context) {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	var user models.User
	var loginForm loginForm

	if err := ctx.ShouldBindJSON(&loginForm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if loginForm.Username == "" && loginForm.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong parameters passed.",
		})
		ctx.Abort()
		return
	}

	if result := db.Where("username = ?", loginForm.Username).First(&user); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		ctx.Abort()
		return
	}

	if (!helpers.Contains([]string{"admin", "employee"}, user.Group)) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You're unathorized to login"})
		ctx.Abort()
		return
	}
	
	if user.CheckPassword(loginForm.Password) != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Wrong password",
		})
		ctx.Abort()
		return
	}

	token, err:= middlewares.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"info": user,
		"token": token,
	})

}
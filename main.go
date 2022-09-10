// Local Api built for Heytel handles all the requests about discounts, shifts, rooms, invoices, users etc...
// Language: go
// Path: main.go

package main

import (
	"Heytel/controllers"
	"Heytel/database"
	"Heytel/helpers"
	"Heytel/middlewares"
	"Heytel/models"

	"log"
	"os"
	"strconv"

	docs "Heytel/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func initializeDatabase() {
	dbPort, _ := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	maxConns, _ := strconv.Atoi(os.Getenv("DATABASE_MAX_CONNS"))
	maxIdle, _ := strconv.Atoi(os.Getenv("DATABASE_MAX_IDLE_CONNS"))

	dbCfg := models.DatabaseConfig{
		Host:         os.Getenv("DATABASE_HOST"),
		Port:         dbPort,
		User:         os.Getenv("DATABASE_USER"),
		Password:     os.Getenv("DATABASE_PASS"),
		Database:     os.Getenv("DATABASE_NAME"),
		Ssl:          os.Getenv("DATABASE_SSL"),
		TimeZone:     os.Getenv("DATABASE_TIMEZONE"),
		MaxDbConns:   maxConns,
		MaxIdleConns: maxIdle,
	}

	database.CreateConnection(dbCfg)
	database.Migrate()
}

func initRouter() *gin.Engine {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// have no idea why its not working rn ^

	router.Use(gin.Logger())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("CORS_ORIGIN")}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)

		user := api.Group("/user")
		{
			registration, err := helpers.GetenvBool("REGISTRATION")

			if err != nil {
				panic(err)
			}
			if registration != false {
				user.POST("/register", controllers.RegisterUser)
			}
			user.POST("/login", controllers.LoginUser)
		}

		secured := api.Group("/secured")
		secured.Use(middlewares.Auth())
		{

			secured.PATCH("/filtered_update", controllers.FilteredUpdate)

			secured.GET("/user/:username", controllers.GetUser)  // "/api/user/:username"
			secured.GET("/user/id/:id", controllers.GetUserById) // "/api/user/:username"
			secured.PATCH("/user/:id", controllers.UpdateUser)

			invoice := secured.Group("/invoice")
			{
				invoice.POST("/", controllers.CreateInvoice)
				invoice.GET("/:id", controllers.GetInvoice)
				invoice.DELETE("/:id", controllers.DeleteInvoice)
				invoice.GET("/room/:id", controllers.GetInvoiceByRoom)
				invoice.GET("/all", controllers.AllInvoices)
				invoice.PATCH("/:id", controllers.UpdateInvoice)
			}
			room := secured.Group("/room")
			{
				room.POST("/", controllers.CreateRoom)
				room.GET("/:id", controllers.GetRoom)
				room.GET("/all", controllers.AllRooms)
				room.DELETE("/:id", controllers.UpdateRoom) // There is difference in Panel source
				room.PATCH("/:id", controllers.UpdateRoom)
				room.GET("/floor/:floor", controllers.AllRoomsByFloor)
				room.GET("/floors", controllers.GetAllFloors)
			}

			bill := secured.Group("/bill")
			{
				bill.POST("/", controllers.CreateBill)
				bill.GET("/:id", controllers.GetBill)
				bill.GET("/invoice/:id", controllers.GetBillByInvoiceId)
			}
			services := secured.Group("/services")
			{
				services.POST("/", controllers.CreateService)
				services.GET("/:id", controllers.GetService)
				services.GET("/all", controllers.GetAllServices)
				services.PATCH("/:id", controllers.UpdateService)

			}
			discounts := secured.Group("/discounts")
			{
				discounts.POST("/", controllers.CreateDiscount)
				discounts.GET("/:id", controllers.GetDiscount)
				discounts.GET("/all", controllers.GetAllDiscounts)
				discounts.PATCH("/:id", controllers.UpdateDiscount)
			}
			shifts := secured.Group("/shifts")
			{
				shifts.POST("/", controllers.CreateShift)
				shifts.GET("/:id", controllers.GetShift)
				shifts.GET("/employee/:id", controllers.GetShiftsByEmployee)
				shifts.PATCH("/:id", controllers.UpdateShift)
			}
			notifications := secured.Group("/notifications")
			{
				notifications.POST("/", controllers.CreateNotification)
				notifications.GET("/:id", controllers.GetNotification)
				notifications.GET("/all/:id", controllers.GetAllNotifications)
				notifications.GET("/receiver/:id", controllers.GetNotificationByReceiver)
				notifications.PATCH("/:id", controllers.UpdateNotification)
			}

			admin := secured.Group("/admin")
			admin.Use(middlewares.AuthAdministrator())
		}
	}
	return router
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	initializeDatabase()

	port := os.Getenv("WEBSERVER_PORT")
	if port == "" {
		port = "8080"
	}

	router := initRouter()
	router.Run(":8080")
}

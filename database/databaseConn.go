package database

import (
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"

	"Heytel/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func GetDSN(cfg models.DatabaseConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port, cfg.Ssl, cfg.TimeZone)
}

func CreateConnection(cfg models.DatabaseConfig) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  GetDSN(cfg),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Error occurred while connecting with the database")
		panic(err)
	}

	sqlDB, err := db.DB()

	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(cfg.MaxDbConns)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	dbConn = db
	return
}

func GetDatabaseConnection() (*gorm.DB, error) {
	sqlDB, err := dbConn.DB()
	if err != nil {
		return dbConn, err
	}
	if err := sqlDB.Ping(); err != nil {
		return dbConn, err
	}
	return dbConn, nil
}

func Migrate() {
	// Auto migrate database
	db, connErr := GetDatabaseConnection()
	if connErr != nil {
		panic(connErr)
	}
	// Add required models here
	err := db.AutoMigrate(&models.User{}, &models.Room{}, &models.Invoice{}, &models.Shift{}, &models.Bill{}, &models.Service{}, &models.Discount{}, &models.Notification{})
	if err != nil {
		panic(err)
	}
}

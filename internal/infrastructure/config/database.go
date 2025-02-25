package config

import (
	"e-commerce/internal/domain/models"
	"log"

	spannergorm "github.com/googleapis/go-gorm-spanner"
	"gorm.io/gorm"
)

func GetDBConnectionString() string {

	return "projects/adroit-resolver-451311-f1/instances/testing/databases/e-commerce"
}

func DbConnection() *gorm.DB {
	userDbConnection, err := gorm.Open(spannergorm.New(spannergorm.Config{
		DriverName: "spanner",
		DSN:        GetDBConnectionString(),
	}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect DB")
	}
	err = userDbConnection.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed successfully!")
	return userDbConnection
}

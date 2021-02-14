package db

import (
	"fibreApi/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// PgConn ...
	PgConn *gorm.DB
)

// ConnectDB ...
func ConnectDB() (*gorm.DB){
	dsn := "host=localhost user=postgres password=postgres dbname=go port=5432"
	var err error
	PgConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	PgConn.AutoMigrate(&models.Story{}, &models.User{})
	fmt.Println("Postgres is here")
	return PgConn
}
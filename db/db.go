package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB ...
func ConnectDB() (*gorm.DB){
	dsn := "host=localhost user=postgres password=postgres dbname=go port=5432"
	pg, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	fmt.Println("Postgres is here")
	return pg
}
package db

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

// ConnectDB ...
func ConnectDB() (*pg.DB){
	db := pg.Connect(&pg.Options{
		User: "postgres",
		Password: "postgres",
		Addr: ":5432",
		Database: "go",
	})

	defer db.Close()

	fmt.Println("Postgres is here")
	
	return db
}
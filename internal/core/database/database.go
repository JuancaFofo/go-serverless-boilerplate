package coredatabase

import (
	"log"

	"github.com/juank11memphis/go-serverless-boilerplate/internal/components/todo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	// TODO implement .env files support for local dev
	// TODO create Config utility to read config values from environment
	dsn := "host=localhost user=postgres password=postgres dbname=goserverless port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Setup Auto Migrations
	db.AutoMigrate(&todo_comp.TodoItem{})

	// TODO add the db host on this log
	log.Println("Successfully connected to the postgres database...")
	return db, nil
}

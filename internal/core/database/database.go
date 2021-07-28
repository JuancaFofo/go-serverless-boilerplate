package coredatabase

import (
	"fmt"
	"log"

	"github.com/juank11memphis/go-serverless-boilerplate/internal/components/todo"
	coreconfig "github.com/juank11memphis/go-serverless-boilerplate/internal/core/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "DB_HOST"
	userName = "DB_USERNAME"
	password = "DB_PASSWORD"
	database = "DB_DATABASE"
)

func Connect(c *coreconfig.Config) (*gorm.DB, error) {
	propertyMap := c.GetValuesMapString(host, userName, password, database)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		propertyMap[host], propertyMap[userName], propertyMap[password], propertyMap[database])

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Setup Auto Migrations
	db.AutoMigrate(&todo.TodoItem{})

	// TODO add the db host on this log
	log.Printf("Successfully connected to postgres DB at host %s\n", host)
	return db, nil
}

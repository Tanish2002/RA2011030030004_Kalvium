package configuration

import (
	"fmt"
	"kalvium/controllers"
	"kalvium/handlers"
	"kalvium/middleware"
	"kalvium/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Configuration struct {
	Handler    handlers.Handler
	Middleware middleware.Middleware
}

// Generate a new state with DB connection
func NewConfig() Configuration {
	// Open a connection to sqlite db
	db, err := gorm.Open(sqlite.Open("models/kalvium.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to db, Error: %v", err))
	}

	// Auto migrate tables
	db.AutoMigrate(new(models.RequestLog))

	return Configuration{
		Handler: handlers.Handler{
			Controller: controllers.Controller{
				Model: models.Model{
					DB: db,
				},
			},
		},
		Middleware: middleware.Middleware{
			Model: models.Model{
				DB: db,
			},
		},
	}

}

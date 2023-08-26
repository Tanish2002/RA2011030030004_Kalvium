package configuration

import (
	"kalvium/controllers"
	"kalvium/handlers"
	"kalvium/middleware"
	"kalvium/models"
)

type Configuration struct {
	Handler    handlers.Handler
	Middleware middleware.Middleware
}

// Generate a new state with DB connection
func NewConfig() Configuration {
	return Configuration{
		Handler: handlers.Handler{
			Controller: controllers.Controller{
				Model: models.Model{
					DB: "",
				},
			},
		},
	}

}

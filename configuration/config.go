package configuration

import (
	"kalvium/controllers"
	"kalvium/handlers"
	"kalvium/models"
)

type Configuration struct {
	Handler handlers.Handler
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

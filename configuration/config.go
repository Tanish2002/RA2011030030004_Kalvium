package configuration

import (
	"kalvium/controllers"
	"kalvium/handlers"
)

type Configuration struct {
	Handler handlers.Handler
}

// Generate a new state with DB connection
func NewConfig() Configuration {
	return Configuration{
		Handler: handlers.Handler{
			Controller: controllers.Controller{
				DB: "",
			},
		},
	}

}

// will add some sort of DB(prob sqllite) for persisten history

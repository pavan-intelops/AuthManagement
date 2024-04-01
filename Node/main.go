package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pavan-intelops/AuthManagement/Node/handlers"
)

func main() {
	e := echo.New()

	//todo: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CreatePets - Create a pet
	e.POST("/v1/pets", c.CreatePets)

	// ListPets - List all pets
	e.GET("/v1/pets", c.ListPets)

	// ShowPetById - Info for a specific pet
	e.GET("/v1/pets/:petId", c.ShowPetById)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/pavan-intelops/AuthManagement/Node/models"
	"net/http"
)

// CreatePets - Create a pet
func (c *Container) CreatePets(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// ListPets - List all pets
func (c *Container) ListPets(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// ShowPetById - Info for a specific pet
func (c *Container) ShowPetById(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

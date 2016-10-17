package handlers

import (
	"firemarksBackend/models"
	"net/http"

	"github.com/labstack/echo"
)

// UsersMountHandler sets the routes for the User resource
func UsersMountHandler(base *echo.Group) {
	base.POST("", createUser)
}

func createUser(c echo.Context) error {
	result := models.NewUser()

	// Read params from context
	if err := c.Bind(result); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Validate model
	if isValid, errors := result.Validate(); isValid != true {
		return c.JSON(http.StatusUnprocessableEntity, errors)
	}

	// Create document
	if err := models.CreateUser(result); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, result)
}

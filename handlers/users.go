package handlers

import (
	"firemarksBackend/models"
	"net/http"

	"github.com/labstack/echo"
)

// UsersMountHandler sets the routes for the User resource
func UsersMountHandler(base *echo.Group) {
	base.GET("", listUsers)
	base.POST("", createUser)
}

// GET /users
func listUsers(c echo.Context) error {
	results := &[]models.User{}

	// Find documents
	if err := models.QueryUsers(results); err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, results)
}

// POST /users
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

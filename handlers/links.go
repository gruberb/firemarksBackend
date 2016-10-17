package handlers

import (
	"firemarksBackend/models"
	"net/http"

	"github.com/labstack/echo"
)

// LinksMountHandler sets the routes for the Link resource
func LinksMountHandler(base *echo.Group) {
	base.GET("", listLinks)
	base.POST("", createLink)
	base.GET("/:id", getLink)
	base.PUT("/:id", updateLink)
	base.DELETE("/:id", deleteLink)
}

// GET /links
func listLinks(c echo.Context) error {
	results := &[]models.Link{}

	// Find documents
	if err := models.QueryLinks(results); err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, results)
}

// GET /links/:id
func getLink(c echo.Context) error {
	result := &models.Link{}
	if err := models.FindLink(c.Param("id"), result); err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, result)
}

// POST /links
func createLink(c echo.Context) error {
	result := models.NewLink()

	// Read params from context
	if err := c.Bind(result); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Validate model
	if isValid, errors := result.Validate(); isValid != true {
		return c.JSON(http.StatusUnprocessableEntity, errors)
	}

	// Create document
	if err := models.CreateLink(result); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, result)
}

// PUT /links/:id
func updateLink(c echo.Context) error {
	changes := new(models.Link)

	// Read params from context
	if err := c.Bind(changes); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Validate model
	if isValid, errors := changes.Validate(); isValid != true {
		return c.JSON(http.StatusUnprocessableEntity, errors)
	}

	// Update document
	if err := models.UpdateLink(c.Param("id"), changes); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

// DELETE /links/:id
func deleteLink(c echo.Context) error {
	// Delete document
	if err := models.DeleteLink(c.Param("id")); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

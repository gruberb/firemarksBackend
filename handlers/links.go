package handlers


import (
	"net/http"
	"github.com/labstack/echo"
	"firemarksBackend/models"
)


// LinksMountHandler sets the routes for /links
func LinksMountHandler(base *echo.Group) {
	base.GET("", list)
	base.POST("", create)
	base.GET("/:id", single)
	base.PUT("/:id", update)
	base.DELETE("/:id", delete)
}


// GET /links
func list(c echo.Context) error {
	results := &[]models.Link{}

	// Find documents
	if err := models.QueryLinks(results); err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, results)
}


// GET /links/:id
func single(c echo.Context) error {
	result := &models.Link{}
	if err := models.FindLink(c.Param("id"), result); err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, result)
}


// POST /links
func create(c echo.Context) error {
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
func update(c echo.Context) error {
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
func delete(c echo.Context) error {
	// Delete document
	if err := models.DeleteLink(c.Param("id")); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

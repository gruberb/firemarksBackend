package handlers

import (
	"net/http"
	// "fmt"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
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
	models.LinkCollection().Find(nil).All(results)

	return c.JSON(http.StatusOK, results)
}


// GET /links/:id
func single(c echo.Context) error {
	result := &models.Link{}

	// Find document
	query := bson.M{"_id": bson.ObjectIdHex(c.Param("id"))}
	if err := models.LinkCollection().Find(query).One(result); err != nil {
		c.NoContent(http.StatusNotFound)
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
	if err := models.LinkCollection().Insert(result); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, result)
}


// PUT /links/:id
func update(c echo.Context) error {
	result := new(models.Link)

	// Read params from context
	if err := c.Bind(result); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Validate model
	if isValid, errors := result.Validate(); isValid != true {
		return c.JSON(http.StatusUnprocessableEntity, errors)
	}

	// Update document
	id := bson.ObjectIdHex(c.Param("id"))
	update := bson.M{"$set": result}
	if err := models.LinkCollection().UpdateId(id, update); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}


// DELETE /links/:id
func delete(c echo.Context) error {
	// Delete document
	id := bson.ObjectIdHex(c.Param("id"))
	if err := models.LinkCollection().RemoveId(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

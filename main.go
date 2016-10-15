package main

import (
	"firemarksBackend/handlers"
	"firemarksBackend/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

type (
	// MountableResource makes posible to use Groups or Echo objects to mount handlers
	MountableResource interface {
		Group(prefix string, middleware ...echo.MiddlewareFunc) *echo.Group
	}

	// MountHandler setups the routes for a MountableResource
	MountHandler func(e *echo.Group)
)

// Mounts a MountHandler at a path for a given MountableResource
func mount(e MountableResource, path string, handler MountHandler) {
	handler(e.Group(path))
}

func main() {
	// Setup a global instance of the DB connection
	models.ConnectDB("localhost", "firemarks")

	// Make sure to disconnect before exiting the main function
	defer models.DisconnectDB()

	// Standard routes
	e := echo.New()
	mount(e, "/", handlers.RootMountHandler)

	// API Routes
	api := e.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	mount(v1, "/links", handlers.LinksMountHandler)
	mount(v1, "/users", handlers.UsersMountHandler)

	// Server stuff
	fmt.Println("\n== Running on http://localhost:3000 ==")
	e.Run(standard.New(":3000"))
}

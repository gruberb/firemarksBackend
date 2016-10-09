package handlers

import (
	_ "net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
)

func TestRoot(t *testing.T) {
	e := echo.New()

	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	expectedResult := root(c)
	actualResult := root(c)

	assert.Equal(t, expectedResult, actualResult, "they should be equal")
}

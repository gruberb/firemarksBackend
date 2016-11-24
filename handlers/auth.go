package handlers

import (
	"firemarksBackend/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func AuthMountHandler(base *echo.Group) {
	base.POST("/login", authUser)
	base.POST("/register", register)
}

// POST/ on login route
func authUser(c echo.Context) error {
	user := &models.User{}
	credentials := &models.User{}
	result := &models.PublicUser{}

	if err := c.Bind(credentials); err != nil {
		return err
	}
	// Find one User by provided E-Mail
	if err := models.FindUser(credentials.EMail, user); err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return c.NoContent(http.StatusForbidden)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	result.ID = user.ID
	result.Name = user.Name
	result.EMail = user.EMail
	t, err := token.SignedString([]byte(user.EMail))
	if err != nil {
		return err
	}

	result.Token = t

	return c.JSON(http.StatusOK, result)
}

// POST/ on register route
func register(c echo.Context) error {
	newUser := models.NewUser()
	result := &models.PublicUser{}

	// Bind the context to fill the user model
	if err := c.Bind(newUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Validate User
	if isValid, errors := newUser.Validate(); isValid != true {
		return c.JSON(http.StatusUnprocessableEntity, errors)
	}

	// Create the new User
	if err := models.CreateUser(newUser); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	result.ID = newUser.ID
	result.Name = newUser.Name
	result.EMail = newUser.EMail

	token := jwt.New(jwt.SigningMethodHS256)
	t, err := token.SignedString([]byte(newUser.EMail))
	if err != nil {
		return err
	}

	result.Token = t

	return c.JSON(http.StatusCreated, result)
}

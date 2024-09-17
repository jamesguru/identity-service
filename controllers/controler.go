package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jamesguru/dev-api/models"
	"github.com/labstack/echo/v4"
)

var identities []models.Identity

func AddIdentity(c echo.Context) error {
	var identity models.Identity
	err := c.Bind(&identity)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if identity.Name == "" {
		return c.JSON(http.StatusPreconditionFailed, "name cannot be empty")
	}
	if identity.Author == "" {
		return c.JSON(http.StatusPreconditionFailed, "author cannot be empty")
	}
	log.Printf("name: %s", identity.Name)
	log.Printf("Author: %s", identity.Author)
	identities = append(identities, identity)
	return c.JSON(http.StatusOK, identities)
}
func Getidentity(c echo.Context) error {
	id := c.Param("id")
	idx, _ := strconv.ParseInt(id, 10, 64)
	return c.JSON(http.StatusOK, identities[idx])
}

func Getidentities(c echo.Context) error {
	return c.JSON(http.StatusOK, "All identities are here.")
}

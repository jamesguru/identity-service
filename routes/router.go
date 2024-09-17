package routes

import (
	"github.com/jamesguru/dev-api/controllers"
	"github.com/labstack/echo/v4"
)

func Setup() {
	e := echo.New()
	e.POST("/identity", controllers.AddIdentity)
	e.GET("/identity/:id", controllers.Getidentity)
	e.GET("/identities", controllers.Getidentities)
	e.Start(":3001")
}

package handler

import (
	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, userHandler UserHandler) {
	e.GET("/user/:id", userHandler.GET())
}

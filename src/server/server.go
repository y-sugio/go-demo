package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func main() {
	e := echo.New()
	e.GET("/user/get", get)
	e.POST("/user/create", create)

	e.Logger.Fatal(e.Start(":8080"))
}

func get(c echo.Context) error {
	u := User{Name:"sugio", Email:"sugio@test.com"}
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func create(c echo.Context) error {
	u := User{}
	if err := c.Bind(&u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}
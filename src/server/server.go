package main

import (
	"github.com/labstack/echo"
	"handler"
	"infrastructure"
	"usecase"
)

func main() {
	userRepository := infrastructure.NewTaskRepository()
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	e := echo.New()
	handler.InitRouting(e, userHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"usecase"
)

type UserHandler interface {
	GET() echo.HandlerFunc
}

type userHandler struct{
	uc usecase.UserUseCase
}

func NewUserHandler(uc usecase.UserUseCase) UserHandler{
	return &userHandler{uc: uc}
}

type requestGetUser struct {
	Id int64 `json:"id"`
}

type responseGetUser struct {
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func (uh *userHandler) GET() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundUser, err := uh.uc.FindByID(int64(id))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseGetUser{
			Id:      foundUser.Id,
			Name:   foundUser.Name,
			Email: foundUser.Email,
		}

		return c.JSON(http.StatusOK, res)
	}
}
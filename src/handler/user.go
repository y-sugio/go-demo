package handler

import (
	"domain/model"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"usecase"
)

type UserHandler interface {
	GET() echo.HandlerFunc
	POST() echo.HandlerFunc
}

type userHandler struct {
	uc usecase.UserUseCase
}

func NewUserHandler(uc usecase.UserUseCase) UserHandler {
	return &userHandler{uc: uc}
}

type responseGetUser struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
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
			Id:    foundUser.Id,
			Name:  foundUser.Name,
			Email: foundUser.Email,
		}

		return c.JSON(http.StatusOK, res)
	}
}

type responseCreateUser struct {
	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (uh *userHandler) POST() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		u := new(model.User)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		createdUser, err := uh.uc.Create(u)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseCreateUser{
			Id:       createdUser.Id,
			Name:     createdUser.Name,
			Email:    createdUser.Email,
			Password: createdUser.Password,
		}

		return c.JSON(http.StatusOK, res)
	}
}

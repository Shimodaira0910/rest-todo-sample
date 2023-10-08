package controller

import (
	"net/http"
	"os"
	"rest-todo-sample/model"
	"rest-todo-sample/usecase"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface{
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
}

type UserController struct{
	uu usecase.IUserUsecase
}

func NewUserControlle(uu usecase.IUserUsecase) IUserController{
	return &UserController{uu}
}

func (uc *UserController) SignUp(c echo.Context) error{
	user := model.User{}
	if err := c.Bind(&user); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.SignUp(user)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

func (uc *UserController) LogIn(c echo.Context) error{
	user := model.User{}
	if err := c.Bind(&user); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteDefaultMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}
package dataRender

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct{
	Name string
	Address string
}

func GetUserController(c echo.Context) error{
	user := User{Name: "Nada", Address: "Jl. Jeruk Purut"}
	return c.JSON(http.StatusOK, user)
}

func RenderGetUser(){
	e:= echo.New()
	e.GET("/user", GetUserController)
	e.Start(":8080")
}
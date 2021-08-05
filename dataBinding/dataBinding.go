package dataBinding

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct{
	Name string `json: "name" form: "name"`
	Email string `json: "email" form: "email"`
}

func CreateUser(c echo.Context) error{
	//binding data 
	user := User{}
	fmt.Println(&user)
	c.Bind(&user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "user created successfully",
		"user" : user,
	})
}

func DataBinding(){
	e:= echo.New()
	e.POST("/users", CreateUser)
	e.Start(":8080")
}
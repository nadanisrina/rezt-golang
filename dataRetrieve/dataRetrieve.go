package dataRetrieve

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct{
	Id int 
	Name string
	Email string 
}

//query param id
func GetUserController(c echo.Context) error {
	id,_ := strconv.Atoi(c.Param("id"))
	user := User{Id: id, Name : "Nada", Email: "nisrinaseptiana@gmail.com"}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user" : user,
	})
}

func RetrieveGetUser(){
	e := echo.New()
	e.GET("/users/:id", GetUserController)
	e.Start(":8080")
}
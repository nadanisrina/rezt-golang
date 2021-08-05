package formValue

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//json nama keynya, form valuenya
// form source is form. Values are taken from query and request body
// json - source is request body. Uses Go json package for unmarshalling.
type User struct{
	Name string `json:"name" form: "name1"`
	Email string `json:"email" form: "email1"`
}


func CreateUserController(c echo.Context) error{
	//get data from value 
	name := c.FormValue("name1")
	email := c.FormValue("email1")

	var user User
	user.Name = name
	user.Email = email

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "user successfully created",
		"user" : user,
	})
}

func FormValue(){
	e := echo.New()
	e.POST("/users", CreateUserController)
	e.Start(":8080")
}
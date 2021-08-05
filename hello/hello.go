package hello

import (
	"net/http"

	"github.com/labstack/echo"
)


func Hello(){
	e:= echo.New()

	e.GET("/", HelloController)

	e.Start(":8080")
}

func HelloController(c echo.Context) error{
	return c.String(http.StatusOK, "Hello World")

}
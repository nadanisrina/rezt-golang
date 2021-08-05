package dataquery

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserSearchController(c echo.Context) error{
	//get data from query param
match := c.QueryParam("match")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"match" : match,
		"users" : []string{"adi", "nada", "asdf"}, //hardcode data
	})

}

func DataQuery(){
	e := echo.New()
	//routing with query parameter
	e.GET("/users", UserSearchController)
	e.Start(":8080")
}
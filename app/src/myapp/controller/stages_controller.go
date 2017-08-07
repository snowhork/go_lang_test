package controller

import (
	"net/http"
	"github.com/labstack/echo"
)

func StagesIndex(c echo.Context) error {
	level   := c.QueryParam("level")
	page    := c.QueryParam("page")
	user_id := c.QueryParam("user_id")
	return c.String(http.StatusOK, "level" + level + "page" + page + "user_id" + user_id)
}

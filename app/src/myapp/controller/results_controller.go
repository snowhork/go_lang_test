package controller

import (
	"github.com/labstack/echo"
	"net/http"


//	"myapp/model"
	"myapp/model"
	"myapp/db"
)

func ResultsCreate(c echo.Context) error {
	//stage_id := c.QueryParam("stage_id")
	user_id := 1
	result := model.Result{UserId: user_id}
	db.Db.Save(&result)
	return c.JSON(http.StatusCreated, map[string]int{"result_id": result.Id})
}


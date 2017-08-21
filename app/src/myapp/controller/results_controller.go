package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/validator.v2"

	"myapp/model"
	"myapp/db"
	"strconv"
)

func ResultsCreate(c echo.Context) error {
	stage_id, err := strconv.Atoi(c.FormValue("stage_id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
	}

	user_id := 1
	result := model.Result{UserId: user_id, StageId: stage_id}

	if errs := validator.Validate(result); errs != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": errs})
	}

	db.Db.Save(&result)

	return c.JSON(http.StatusOK, map[string]int{"result_id": result.Id})
}

func ResultsUpdate(c echo.Context) error {
	user_id := 1
	result_id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
	}

	result := model.Result{}
	db.Db.First(&result, result_id)

	if result.Id == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Result Not Found"})
	}

	if result.UserId != user_id {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Incorrect User"})
	}

	db.Db.Model(&result).Update("Status", true)
	return c.JSON(http.StatusOK, map[string]int{"result_id": result.Id})
}

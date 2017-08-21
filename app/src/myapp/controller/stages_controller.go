package controller

import (
	"net/http"
	"github.com/labstack/echo"
	"gopkg.in/validator.v2"

	"myapp/db"
	"strconv"
	"os"
	"io"
	"myapp/model"
)

type StagesInfo struct {
	Id int              `json:"id"`
	Name string			`json:"name"`
	UserId  int			`json:"user_id"`
	UserName string		`json:"user_name"`
	Status bool			`json:"status"`
	ChallengedCount int `json:"challenge_count"`
	ClearedCount int	`json:"cleared_count"`
}


func StagesIndex(c echo.Context) error {
	//level, err := strconv.Atoi(c.QueryParam("level"))
	//
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
	//}
	//
	//page, err := strconv.Atoi(c.QueryParam("level"))
	//
	//if err != nil {
	//	page = 1
	//}
	//
	//user_id, _ := strconv.Atoi(c.QueryParam("user_id"))

	//user_id := 1


	var stages_info []StagesInfo
	db.Db.Raw(`
		SELECT stages.id, stages.name, users.id AS user_id, users.name AS user_name, results.challenged_count, results.cleared_count
		FROM stages
		JOIN users
		ON stages.user_id = users.id
		LEFT JOIN(
			SELECT results.stage_id, count(results.id) AS challenged_count, cleared_results.cleared_count
			FROM results
			JOIN stages
			ON stages.id = results.stage_id
			LEFT JOIN (
				SELECT count(results.id) AS cleared_count, stage_id
				FROM results
				WHERE results.status = true
				GROUP BY results.stage_id
			) AS cleared_results
			ON stages.id = cleared_results.stage_id
			group by results.stage_id
		) AS results
		ON stages.id = results.stage_id`).Scan(&stages_info)

	return c.JSON(http.StatusOK, stages_info)
}

func StagesShow(c echo.Context) error {
	stage_id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
	}

	stage := model.Stage{}
	db.Db.First(&stage, stage_id)

	if stage.UserId == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Stage Not Found"})
	}

	return c.File("file/csv/" + strconv.Itoa(stage.UserId) + "/" + strconv.Itoa(stage.Id)+ ".csv")
}


func StagesCreate(c echo.Context) error {
	tx := db.Db.Begin() // begin transaction

	name := c.FormValue("name")

	if name == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Name must not be empty"})
	}

	user_id := 2
	stage := model.Stage{UserId: user_id, Name: name}

	if errs := validator.Validate(stage); errs != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": errs})
	}

	tx.Save(&stage)

	//-----------
	// Read file
	//-----------

	// Source

	file, err := c.FormFile("csv")
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	src, err := file.Open()
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}
	defer src.Close()

	// Destination
	file_name := "file/csv/" + strconv.Itoa(user_id)
	os.Mkdir(file_name, 0777)
	dst, err := os.Create(file_name + "/" + strconv.Itoa(stage.Id) + ".csv")
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	tx.Commit()
	return c.JSON(http.StatusOK, map[string]int{"stage_id": stage.Id})
}


func StagesDelete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	stage := model.Stage{}
	db.Db.First(&stage, id)

	if stage.Id == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Stage not Found"})
	}

	user_id := 1
	if stage.UserId != user_id {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Permission denied"})
	}

	db.Db.Delete(&stage, id)
	return c.JSON(http.StatusOK, map[string]int{})
}

package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lucass-segura/go-echo-restapi/db"
	"github.com/lucass-segura/go-echo-restapi/models"
)

func PostTaskHandler(c echo.Context) error {
	var task models.Task
	json.NewDecoder(c.Request().Body).Decode(&task)
	createTask := db.DB.Create(&task)
	err := createTask.Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Error creating task in database"})
	}

	return json.NewEncoder(c.Response()).Encode(&task)
}

func GetTasksHandler(c echo.Context) error {
	var tasks []models.Task
	db.DB.Find(&tasks)

	return json.NewEncoder(c.Response()).Encode(&tasks)
}

func GetTaskHandler(c echo.Context) error {
	var task models.Task
	id := c.Param("id")
	taskID, err := strconv.Atoi(id)

	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid task ID")
	}

	db.DB.First(&task, taskID)

	if task.ID == 0 {
		return c.String(http.StatusNotFound, "Not found task ID")
	}
	return json.NewEncoder(c.Response()).Encode(&task)
}

func DeleteTaskHandler(c echo.Context) error {
	var task models.Task
	id := c.Param("id")
	taskID, err := strconv.Atoi(id)

	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid task ID")
	}
	if task.ID == 0 {
		fmt.Println(task.ID)
		return c.String(http.StatusNotFound, "Not found task ID")
	}

	fmt.Println(task.ID)
	fmt.Println(taskID)
	db.DB.Unscoped().Delete(&task, taskID)

	return c.String(http.StatusAccepted, "Task delete")
}

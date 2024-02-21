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

func PostUserHandler(c echo.Context) error {
	var user models.User
	fmt.Println(c.Request().Body)

	json.NewDecoder(c.Request().Body).Decode(&user)
	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error creating user in database"})
	}

	return c.JSON(http.StatusCreated, user)
}

func GetUsersHandler(c echo.Context) error {
	var users []models.User
	db.DB.Find(&users)

	return json.NewEncoder(c.Response()).Encode(&users)
}

func GetUserHandler(c echo.Context) error {
	var user models.User
	id := c.Param("id")
	userID, err := strconv.Atoi(id)

	if err != nil {
		c.String(http.StatusBadRequest, "Invalid user ID")
		return err
	}

	db.DB.First(&user, userID)

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	if user.ID == 0 {
		return c.String(http.StatusNotFound, "Not found user")
	}

	return json.NewEncoder(c.Response()).Encode(&user)
}

func DeleteUserHandler(c echo.Context) error {
	var user models.User
	id := c.Param("id")
	userID, err := strconv.Atoi(id)

	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user ID")
	}
	if user.ID == 0 {
		return c.String(http.StatusNotFound, "Not found user")
	}

	db.DB.Unscoped().Delete(&user, userID)

	return c.String(http.StatusAccepted, "User deleted")
}

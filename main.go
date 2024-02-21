package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lucass-segura/go-echo-restapi/db"
	"github.com/lucass-segura/go-echo-restapi/models"
	"github.com/lucass-segura/go-echo-restapi/routes"
)

func main() {
	db.DBconnection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	e := echo.New() //Instancia del marco web
	e.GET("/", routes.IndexHandle)

	//USERS
	e.GET("/users", routes.GetUsersHandler)
	e.GET("/user/:id", routes.GetUserHandler)
	e.POST("/user", routes.PostUserHandler)
	e.DELETE("/user/:id", routes.DeleteUserHandler)

	//TASKS
	e.GET("/tasks", routes.GetTasksHandler)
	e.GET("/task/:id", routes.GetTaskHandler)
	e.POST("/task", routes.PostTaskHandler)
	e.DELETE("/task/:id", routes.DeleteTaskHandler)

	e.Logger.Fatal(e.Start(":8001")) //comenzará a escuchar las solicitudes entrantes en ese puerto
}

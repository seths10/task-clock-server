package main

import (
	"github.com/gin-gonic/gin"
	"github.com/task-clock-server/internal/config"
	"github.com/task-clock-server/internal/controllers"
	"github.com/task-clock-server/internal/middleware"
)

func main() {
	config.InitMongo()

	r := gin.Default()
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.Use(middleware.AuthMiddleware())

	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks", controllers.GetTasks)
	r.DELETE("/tasks/:id", controllers.DeleteTask)

	r.Run()
}

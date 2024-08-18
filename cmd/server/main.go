package main

import (
	"log"
	"net/http"

	"github.com/task-clock-server/internal/config"
	"github.com/task-clock-server/internal/handler"
	"github.com/task-clock-server/internal/repository/mongodb"
	"github.com/task-clock-server/internal/service"
	mongoClient "github.com/task-clock-server/pkg/mongodb"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	client, err := mongoClient.NewClient(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	taskRepo := mongodb.NewTaskRepository(client.Database(cfg.MongoDB))
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	r := mux.NewRouter()
	r.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{cfg.FrontendURL},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(r)
	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler))

}

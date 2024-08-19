package controllers

import (
    "context"
    "net/http"
    "strconv"
    "github.com/task-clock-server/internal/config"
    "github.com/task-clock-server/internal/models"
    "go.mongodb.org/mongo-driver/bson"
    "github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userId, _ := c.Get("userId")
    task.UserID = userId.(string)

    _, err := config.TaskCollection.InsertOne(context.TODO(), task)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, task)
}


func GetTasks(c *gin.Context) {
    userId, _ := c.Get("userId")

    var tasks []models.Task
    cursor, err := config.TaskCollection.Find(context.TODO(), gin.H{"userId": userId})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if err := cursor.All(context.TODO(), &tasks); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, tasks)
}

func DeleteTask(c *gin.Context) {
    userId, _ := c.Get("userId")
    id := c.Param("id")

    taskID, err := strconv.Atoi(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }

    filter := bson.M{"id": taskID, "userId": userId}

    result, err := config.TaskCollection.DeleteOne(context.TODO(), filter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if result.DeletedCount == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

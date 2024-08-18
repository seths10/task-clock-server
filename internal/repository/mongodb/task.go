package mongodb

import (
    "context"

    "github.com/task-clock-server/internal/model"
    "github.com/task-clock-server/internal/repository"

    "go.mongodb.org/mongo-driver/bson"
		"go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type taskRepository struct {
    collection *mongo.Collection
}

func NewTaskRepository(db *mongo.Database) repository.TaskRepository {
    return &taskRepository{
        collection: db.Collection("tasks"),
    }
}

func (r *taskRepository) FindByUserID(ctx context.Context, userID string) ([]model.Task, error) {
    var tasks []model.Task
    cursor, err := r.collection.Find(ctx, bson.M{"userId": userID})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    if err = cursor.All(ctx, &tasks); err != nil {
        return nil, err
    }
    return tasks, nil
}

func (r *taskRepository) Create(ctx context.Context, task model.Task) (model.Task, error) {
    result, err := r.collection.InsertOne(ctx, task)
    if err != nil {
        return model.Task{}, err
    }
    task.ID = result.InsertedID.(primitive.ObjectID)
    return task, nil
}

// Implement other methods
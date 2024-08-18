package repository

import (
    "context"
    "github.com/task-clock-server/internal/model"
)

type TaskRepository interface {
    FindByUserID(ctx context.Context, userID string) ([]model.Task, error)
    Create(ctx context.Context, task model.Task) (model.Task, error)
}
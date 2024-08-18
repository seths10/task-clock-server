package service

import (
    "context"

    "github.com/task-clock-server/internal/model"
    "github.com/task-clock-server/internal/repository"
)

type TaskService struct {
    repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
    return &TaskService{repo: repo}
}

func (s *TaskService) GetTasks(ctx context.Context, userID string) ([]model.Task, error) {
    return s.repo.FindByUserID(ctx, userID)
}

func (s *TaskService) CreateTask(ctx context.Context, task model.Task) (model.Task, error) {
    return s.repo.Create(ctx, task)
}
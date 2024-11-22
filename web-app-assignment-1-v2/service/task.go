package service

import (
	"fmt"
	repo "a21hc3NpZ25tZW50/repository"
)

type TaskService interface {
	Store(task *repo.Task) error
	Update(task *repo.Task) error
	Delete(id int) error
	GetByID(id int) (*repo.Task, error)
	GetList() ([]repo.Task, error)
	GetTaskCategory(id int) ([]repo.TaskCategory, error)
}

type taskService struct {
	taskRepository repo.TaskRepository
}

func NewTaskService(taskRepository repo.TaskRepository) TaskService {
	return &taskService{taskRepository}
}

func (s *taskService) Store(task *repo.Task) error {
	if task == nil {
		return fmt.Errorf("task tidak boleh nil")
	}

	err := s.taskRepository.Store(task)
	if err != nil {
		return fmt.Errorf("gagal menyimpan task: %v", err)
	}

	return nil
}

// ... (metode lainnya serupa, ganti model.Task menjadi repo.Task)

func convertTaskToTaskCategory(task repo.Task) repo.TaskCategory {
	return repo.TaskCategory{
		ID:          task.CategoryID,
		Name:        task.CategoryName,
		Description: task.CategoryDescription,
	}
}

func (s *taskService) GetTaskCategory(id int) ([]repo.TaskCategory, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID kategori tidak valid")
	}

	tasks, err := s.taskRepository.GetTaskCategory(id)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil kategori task: %v", err)
	}

	taskCategories := make([]repo.TaskCategory, len(tasks))
	for i, task := range tasks {
		taskCategories[i] = convertTaskToTaskCategory(task)
	}

	return taskCategories, nil
}
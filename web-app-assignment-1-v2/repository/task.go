package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"fmt"
)
type Task struct {
	ID                  int
	Title               string
	Description         string
	CategoryID          int
	CategoryName        string
	CategoryDescription string
}
type TaskCategory struct {
	ID          int
	Name        string
	Description string
}

type TaskRepository interface {
	Store(task *model.Task) error
	Update(task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(categoryID int) ([]model.Task, error)
}

type taskRepository struct {
	filebased *filebased.Data
}

func NewTaskRepo(filebasedDb *filebased.Data) TaskRepository {
	return &taskRepository{
		filebased: filebasedDb,
	}
}

func (t *taskRepository) Store(task *model.Task) error {
	if task == nil {
		return fmt.Errorf("task cannot be nil")
	}

	err := t.filebased.StoreTask(*task)
	if err != nil {
		return fmt.Errorf("failed to store task: %v", err)
	}

	return nil
}

func (t *taskRepository) Update(task *model.Task) error {
	if task == nil {
		return fmt.Errorf("task cannot be nil")
	}

	// Pastikan task dengan ID tersebut ada
	_, err := t.GetByID(task.ID)
	if err != nil {
		return fmt.Errorf("task not found: %v", err)
	}

	// Panggil UpdateTask dengan ID dan task
	err = t.filebased.UpdateTask(task.ID, *task)
	if err != nil {
		return fmt.Errorf("failed to update task: %v", err)
	}

	return nil
}

func (t *taskRepository) Delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid task ID")
	}

	// Pastikan task dengan ID tersebut ada
	_, err := t.GetByID(id)
	if err != nil {
		return fmt.Errorf("task not found: %v", err)
	}

	err = t.filebased.DeleteTask(id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %v", err)
	}

	return nil
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid task ID")
	}

	task, err := t.filebased.GetTaskByID(id)
	if err != nil {
		return nil, fmt.Errorf("task not found: %v", err)
	}

	return task, nil
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	tasks, err := t.filebased.GetTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to get task list: %v", err)
	}

	return tasks, nil
}

func (t *taskRepository) GetTaskCategory(categoryID int) ([]model.Task, error) {
	if categoryID <= 0 {
		return nil, fmt.Errorf("invalid category ID")
	}

	tasks, err := t.filebased.GetTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %v", err)
	}

	var filteredTasks []model.Task
	for _, task := range tasks {
		if task.CategoryID == categoryID {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return filteredTasks, nil
}
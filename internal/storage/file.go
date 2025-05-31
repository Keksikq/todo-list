package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"todo-list-for-levus/internal/model"
)

var (
	ErrTaskNotFound = fmt.Errorf("task not found")
)

type FileStorage struct {
	filePath string
	mu       sync.RWMutex
	tasks    []*model.Task
}

func NewFileStorage(filePath string) (*FileStorage, error) {
	storage := &FileStorage{
		filePath: filePath,
		tasks:    make([]*model.Task, 0),
	}

	if err := storage.load(); err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *FileStorage) load() error {
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		dir := filepath.Dir(s.filePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
		if err := os.WriteFile(s.filePath, []byte("[]"), 0644); err != nil {
			return err
		}
	}

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if len(data) == 0 {
		s.tasks = make([]*model.Task, 0)
		return nil
	}

	return json.Unmarshal(data, &s.tasks)
}

func (s *FileStorage) save() error {
	s.mu.RLock()
	tasksCopy := make([]*model.Task, len(s.tasks))
	copy(tasksCopy, s.tasks)
	s.mu.RUnlock()

	data, err := json.MarshalIndent(tasksCopy, "", "  ")
	if err != nil {
		return err
	}

	dir := filepath.Dir(s.filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(s.filePath, data, 0644)
}

func (s *FileStorage) AddTask(task *model.Task) error {
	s.mu.Lock()
	if len(s.tasks) == 0 {
		task.ID = 1
	} else {
		task.ID = s.tasks[len(s.tasks)-1].ID + 1
	}
	s.tasks = append(s.tasks, task)
	s.mu.Unlock()

	return s.save()
}

func (s *FileStorage) GetTasks() []*model.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]*model.Task, len(s.tasks))
	copy(tasks, s.tasks)
	return tasks
}

func (s *FileStorage) UpdateTask(id int, task *model.Task) error {
	s.mu.Lock()
	found := false
	for i, t := range s.tasks {
		if t.ID == id {
			task.ID = id
			s.tasks[i] = task
			found = true
			break
		}
	}
	s.mu.Unlock()

	if !found {
		return ErrTaskNotFound
	}

	return s.save()
}

func (s *FileStorage) DeleteTask(id int) error {
	s.mu.Lock()
	found := false
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			found = true
			break
		}
	}
	s.mu.Unlock()

	if !found {
		return ErrTaskNotFound
	}

	return s.save()
}

func (s *FileStorage) GetTaskByID(id int) (*model.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, task := range s.tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return nil, ErrTaskNotFound
}

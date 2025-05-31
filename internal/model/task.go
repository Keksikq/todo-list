package model

import (
	"fmt"
	"time"
)

var (
	ErrEmptyTitle      = fmt.Errorf("title cannot be empty")
	ErrInvalidPriority = fmt.Errorf("priority must be between 1 and 5")
	ErrInvalidDate     = fmt.Errorf("invalid date format, use YYYY-MM-DD")
	ErrInvalidStatus   = fmt.Errorf("status must be either 'pending' or 'done'")
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	DueDate     string `json:"due_date"`
	Status      string `json:"status"`
}

func NewTask(title string, description string, priority int, dueDate string) *Task {
	return &Task{
		Title:       title,
		Description: description,
		Priority:    priority,
		DueDate:     dueDate,
		Status:      "pending",
	}
}

func (t *Task) Validate() error {
	if t.Title == "" {
		return ErrEmptyTitle
	}
	if t.Priority < 1 || t.Priority > 5 {
		return ErrInvalidPriority
	}
	if t.DueDate != "" {
		if _, err := time.Parse("2006-01-02", t.DueDate); err != nil {
			return ErrInvalidDate
		}
	}
	if t.Status != "pending" && t.Status != "done" {
		return ErrInvalidStatus
	}
	return nil
}

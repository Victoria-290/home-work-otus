package task

import (
	"time"
)

// Status — статус задачи
type Status string

const (
	StatusTodo       Status = "TODO"
	StatusInProgress Status = "IN_PROGRESS"
	StatusDone       Status = "DONE"
)

// Task — бизнес-сущность задачи
type Task struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	OwnerID     int64     `json:"owner_id"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewTask — конструктор новой задачи
func NewTask(title string, description string, ownerID int64) *Task {
	return &Task{
		Title:       title,
		Description: description,
		OwnerID:     ownerID,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// ChangeStatus — изменение статуса задачи
func (t *Task) ChangeStatus(newStatus Status) {
	t.Status = newStatus
	t.UpdatedAt = time.Now()
}

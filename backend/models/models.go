package app

import "time"

type Todo struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	Favorite    bool       `json:"favorite"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	Priority    int        `json:"priority"`
}

type FocusSession struct {
	ID              int        `json:"id"`
	TodoID          int        `json:"todo_id"`
	StartTime       time.Time  `json:"start_time"`
	EndTime         *time.Time `json:"end_time,omitempty"`
	DurationMinutes int        `json:"duration_minutes"`
	Completed       bool       `json:"completed"`
	Todo            *Todo      `json:"todo,omitempty"`
}

type CreateTodoRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	Priority    int        `json:"priority"`
}

type UpdateTodoRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	Favorite    bool       `json:"favorite"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	Priority    int        `json:"priority"`
}

type StartFocusRequest struct {
	TodoID          int `json:"todo_id"`
	DurationMinutes int `json:"duration_minutes"`
}

package models// Task representa la estructura de una tarea

import "time"

type Task struct {
    ID          string    `json:"id,omitempty"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    StartAt     time.Time `json:"startAt"`
    DueDate     time.Time `json:"dueDate"`
    UserID      string    `json:"userId"` // ID del usuario al que pertenece la tarea
}
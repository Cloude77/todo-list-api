// Определяем модели

package models

import "time"

// Task представляет собой одну задачу в To-Do List

type Task struct {
	ID        int       `json:"id"`         // Уникальный ID задачи
	Title     string    `json:"title"`      // Название задачи
	Done      bool      `json:"done"`       // Статус выполнения
	CreatedAt time.Time `json:"created_at"` // Время создания
}

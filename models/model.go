package models

import "time"

type TODO struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

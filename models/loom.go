package models

import "time"

type Loom struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    string    `json:"username"`
	Link      string    `json:"link"`
	Content   string    `json:"content"`
	Type      string    `json:"type"` // Event OR Normal
	Likes     int       `json:"likes"`
	Event     Event     `json:"event"`
	CreatedAt time.Time `json:"createdAt"`
}

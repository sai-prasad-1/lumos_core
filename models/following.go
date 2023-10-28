package models

import "time"

type Following struct {
	ID          string `gorm:"primaryKey"`
	UserID      string
	FollowingID string
	CreatedAt   time.Time
}

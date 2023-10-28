package models

import "time"

type Follower struct {
	ID         string `gorm:"primaryKey"`
	UserID     string
	FollowerID string
	CreatedAt  time.Time
}

package models

import "time"

type Event struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	LoomID      string    `json:"loomId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Date        time.Time `json:"date"`
	MediaLink   string    `json:"mediaLink"`
}

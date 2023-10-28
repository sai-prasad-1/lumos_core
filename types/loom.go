package types

import (
	"time"

	"github.com/akhil-is-watching/post_service/models"
)

type LoomCreate struct {
	UserID  string `json:"userID"`
	Link    string `json:"link"`
	Content string `json:"content"`
}

type EventCreate struct {
	UserID      string    `json:"userID"`
	Link        string    `json:"link"`
	Content     string    `json:"content"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	MaxMembers  int       `json:"maxNumbers"`
	Type        string    `json:"type"`
	Date        time.Time `json:"date"`
}

type FeedLoomFetch struct {
	Username string `json:"username"`
}

func (pc LoomCreate) Convert() *models.Loom {
	return &models.Loom{
		UserID:    pc.UserID,
		Link:      pc.Link,
		Content:   pc.Content,
		Type:      "Loom",
		Likes:     0,
		CreatedAt: time.Now(),
	}
}

// func (pc LoomCreate) ConvertToEvent() (*models.Loom, *models.Event) {
// 	loom := &models.Loom{
// 		UserID:    pc.UserID,
// 		Link:      pc.Link,
// 		Type:      "Event",
// 		Likes:     0,
// 		CreatedAt: time.Now(),
// 	}

// 	event := &models.Event{
// 		ID: helper.GenerateID(),

// 	}
// }

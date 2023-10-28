package repository

import (
	"github.com/akhil-is-watching/post_service/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db}
}

func (repo *EventRepository) Create(p *models.Event) error {
	if err := repo.db.Create(p).Error; err != nil {
		return err
	}
	return nil
}

func (repo *EventRepository) Get(EventID string) (models.Event, error) {
	var event models.Event
	if err := repo.db.Where("id = ?", EventID).First(&event).Error; err != nil {
		return event, err
	}

	return event, nil
}

func (repo *EventRepository) All() ([]models.Event, error) {
	var events []models.Event
	if err := repo.db.Find(&events).Error; err != nil {
		return events, err
	}

	return events, nil
}

func (repo *EventRepository) AllEventsByUser(Username string) ([]models.Event, error) {
	var user *models.User
	if err := repo.db.Preload("Teams").First(&user, Username).Error; err != nil {
		return nil, err
	}

	var events []models.Event
	for _, team := range user.Teams {
		var event models.Event
		if err := repo.db.First(&event, team.EventID).Error; err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

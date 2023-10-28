package models

type Team struct {
	ID      string `gorm:"primaryKey"`
	Name    string
	EventID string
	Event   Event `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Status  string
	RSVP    bool
	Users   []User `gorm:"many2many:team_users;"`
}

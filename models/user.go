package models

import "time"

type User struct {
	ID                string      `gorm:"primaryKey" json:"username"`
	Name              string      `json:"name"`
	College           string      `json:"college"`
	Description       string      `json:"description"`
	Email             string      `gorm:"unique" json:"email"`
	PasswordHash      string      `json:"passwordHash"`
	WalletAddress     string      `json:"walletAddress"`
	WalletPrivkeyHash string      `json:"walletKeyEnc"`
	Followers         []Follower  `gorm:"foreignKey:UserID" json:"followers"`
	Following         []Following `gorm:"foreignKey:UserID" json:"following"`
	Teams             []Team      `gorm:"many2many:team_users;"`
	Looms             []Loom      `json:"looms"`
	CreatedAt         time.Time   `json:"createdAt"`
}

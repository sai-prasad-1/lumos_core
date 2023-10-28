package repository

import (
	"errors"
	"strings"

	"github.com/akhil-is-watching/post_service/models"
	"gorm.io/gorm"
)

type TeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{db}
}

func (repo *TeamRepository) Create(u *models.Team) error {
	if err := repo.db.Create(u).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique") {
			return errors.New("email already exists")
		}
		return err
	}
	return nil
}

func (repo *TeamRepository) Get(Teamname string) (*models.Team, error) {
	var team *models.Team
	if err := repo.db.Preload("Users").Where("name = ?", Teamname).First(&team).Error; err != nil {
		return team, err
	}

	return team, nil
}

func (repo *TeamRepository) AddMember(UserID string, TeamID string) (*models.Team, error) {
	var user *models.User
	var team *models.Team

	if err := repo.db.First(&user, UserID).Error; err != nil {
		return team, err
	}

	if err := repo.db.Preload("Users").First(&team, TeamID).Error; err != nil {
		return team, err
	}

	if err := repo.db.Model(&team).Association("Users").Append(&user); err != nil {
		return team, err
	}

	return team, nil
}

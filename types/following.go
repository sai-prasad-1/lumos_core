package types

import (
	"fmt"

	"github.com/akhil-is-watching/post_service/models"
)

type FollowingCreate struct {
	UserID      string `json:"userId"`
	FollowingID string `json:"followerId"`
}

func (fc FollowingCreate) ConvertToFollowing() *models.Following {
	return &models.Following{
		ID:          fmt.Sprintf("%s_%s", fc.UserID, fc.FollowingID),
		UserID:      fc.UserID,
		FollowingID: fc.FollowingID,
	}
}

func (fc FollowingCreate) ConvertToFollower() *models.Follower {
	return &models.Follower{
		ID:         fmt.Sprintf("%s_%s", fc.FollowingID, fc.UserID),
		UserID:     fc.FollowingID,
		FollowerID: fc.UserID,
	}
}

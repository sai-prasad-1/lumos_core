package controllers

// import (
// 	"time"

// 	"github.com/akhil-is-watching/post_service/helper"
// 	"github.com/akhil-is-watching/post_service/models"
// 	"github.com/akhil-is-watching/post_service/repository"
// 	"github.com/akhil-is-watching/post_service/storage"
// 	"github.com/akhil-is-watching/post_service/types"
// 	"github.com/gofiber/fiber/v2"
// )

// func CreateEvent(c *fiber.Ctx) error {
// 	var params *types.EventCreate
// 	if err := c.BodyParser(&params); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	loom := &models.Loom{
// 		ID:        helper.GenerateID(),
// 		UserID:    params.UserID,
// 		Link:      params.Link,
// 		Content:   params.Content,
// 		Type:      "Event",
// 		Likes:     0,
// 		CreatedAt: time.Now(),
// 	}

// 	loomRepo := repository.NewLoomRepository(storage.)
// }

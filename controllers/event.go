package controllers

import (
	"time"

	"github.com/akhil-is-watching/post_service/helper"
	"github.com/akhil-is-watching/post_service/models"
	"github.com/akhil-is-watching/post_service/repository"
	"github.com/akhil-is-watching/post_service/storage"
	"github.com/akhil-is-watching/post_service/types"
	"github.com/gofiber/fiber/v2"
)

func CreateEvent(c *fiber.Ctx) error {
	var params *types.EventCreate
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	loom := &models.Loom{
		ID:        helper.GenerateID(),
		UserID:    params.UserID,
		Link:      params.Link,
		Content:   params.Content,
		Type:      "Event",
		Likes:     0,
		CreatedAt: time.Now(),
	}

	loomRepo := repository.NewLoomRepository(storage.GetDB())
	if err := loomRepo.Create(loom); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	event := &models.Event{
		ID:          helper.GenerateID(),
		LoomID:      loom.ID,
		Name:        params.Name,
		Description: params.Description,
		Type:        "In-Person",
		Date:        time.Date(2023, time.April, 23, 12, 0, 0, 0, time.UTC),
	}
	eventRepo := repository.NewEventRepository(storage.GetDB())
	if err := eventRepo.Create(event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"loom":  loom,
		"event": event,
	})
}

func Events(c *fiber.Ctx) error {
	eventsRepo := repository.NewEventRepository(storage.GetDB())
	events, err := eventsRepo.All()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"events": events,
	})
}

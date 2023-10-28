package controllers

import (
	"github.com/akhil-is-watching/post_service/repository"
	"github.com/akhil-is-watching/post_service/storage"
	"github.com/akhil-is-watching/post_service/types"
	"github.com/gofiber/fiber/v2"
)

func CreateLoom(c *fiber.Ctx) error {
	var params types.LoomCreate
	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	loom := params.Convert()
	loomRepo := repository.NewLoomRepository(storage.GetDB())
	if err := loomRepo.Create(loom); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"loom": loom,
	})
}

func GetLooms(c *fiber.Ctx) error {
	username := c.Params("username")

	loomRepo := repository.NewLoomRepository(storage.GetDB())
	feed, err := loomRepo.GetFeedForUser(username)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"feed": feed,
	})
}

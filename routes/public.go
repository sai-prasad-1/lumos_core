package routes

import (
	"github.com/akhil-is-watching/post_service/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "Still Alive!!!",
		})
	})
	user := app.Group("/user")
	user.Get("/:username", controllers.UserProfile)
	user.Post("/", controllers.Createuser)

	user.Post("/follow", controllers.FollowUser)
}

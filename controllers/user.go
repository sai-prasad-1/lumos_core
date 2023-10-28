package controllers

import (
	"github.com/akhil-is-watching/post_service/repository"
	"github.com/akhil-is-watching/post_service/storage"
	"github.com/akhil-is-watching/post_service/types"
	"github.com/gofiber/fiber/v2"
)

func Createuser(c *fiber.Ctx) error {
	var params *types.UserCreate
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := params.Convert()
	userRepo := repository.NewUserRepository(storage.GetDB())
	if err := userRepo.Create(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user": user,
	})
}

func FollowUser(c *fiber.Ctx) error {
	var params types.FollowingCreate
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	following := params.ConvertToFollowing()
	follower := params.ConvertToFollower()

	followingRepo := repository.NewFollowingRepository(storage.GetDB())
	followerRepo := repository.NewFollowerRepository(storage.GetDB())

	if err := followingRepo.Create(following); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	if err := followerRepo.Create(follower); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"follower":  follower,
		"following": following,
	})

}

// func GetFollowersFollowing(c *fiber.Ctx) error {
// 	username := c.Params("username")
// 	if username == "" {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Username parameter is required",
// 		})
// 	}
// 	followerRepo := repository.NewFollowerRepository(storage.GetDB())
// 	followingRepo := repository.NewFollowingRepository(storage.GetDB())
// 	followers, err := followerRepo.AllFollowerByUser(username)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"msg": err.Error(),
// 		})
// 	}

// 	following, err := followingRepo.AllFollowingByUser(username)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"msg": err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"followerCount":  len(followers),
// 		"followingCount": len(following),
// 		"followers":      followers,
// 		"following":      following,
// 	})
// }

func UserProfile(c *fiber.Ctx) error {
	username := c.Params("username")

	userRepo := repository.NewUserRepository(storage.GetDB())
	followerRepo := repository.NewFollowerRepository(storage.GetDB())
	followingRepo := repository.NewFollowingRepository(storage.GetDB())
	// eventRepo := repository.NewEventRepository(storage.GetDB())
	loomRepo := repository.NewLoomRepository(storage.GetDB())

	user, err := userRepo.Get(username)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	followers, err := followerRepo.AllFollowerByUser(username)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	following, err := followingRepo.AllFollowingByUser(username)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	looms, err := loomRepo.AllLoomsByUser(username)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":           user,
		"looms":          looms,
		"followers":      followers,
		"following":      following,
		"followerCount":  len(followers),
		"followingCount": len(following),
	})
}

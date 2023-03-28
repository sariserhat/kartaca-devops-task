package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/serhhatsari/task/repository"
)

func GetCountry(repo *repository.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		country, err := repo.GetRandomCountry()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
		return c.JSON(country)
	}
}

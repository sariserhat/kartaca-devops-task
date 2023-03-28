package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/serhhatsari/task/handler"
	"github.com/serhhatsari/task/repository"
	"log"
)

type Application struct {
	app  *fiber.App
	repo *repository.Repository
}

func main() {
	repo := repository.New()

	newFiber := fiber.New()

	app := &Application{
		app:  newFiber,
		repo: repo,
	}

	app.app.Get("/", handler.Hello())
	app.app.Get("/staj", handler.GetCountry(app.repo))
	app.app.Use(handler.NotFound())

	log.Fatal(app.app.Listen(":5555"))
}

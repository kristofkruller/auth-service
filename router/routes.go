package router

import (
	"fmt"
	"log"

	"firebase.google.com/go/v4/db"
	"github.com/gofiber/fiber/v2"
	"github.com/kristofkruller/calendar-service/handlers"
	"github.com/kristofkruller/calendar-service/models"
)

func SetupRoutes(app *fiber.App, db *db.Client) {
	auth := app.Group("/")

	auth.Post("/login", func(c *fiber.Ctx) error {
		event := new(models.Event)
		if err := c.BodyParser(event); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		if err := handlers.SaveEvent(db, event); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(event)
	})

	auth.Delete("/logout", func(c *fiber.Ctx) error {
		if err := handlers.DeleteAllEvents(db); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return nil
	})
}

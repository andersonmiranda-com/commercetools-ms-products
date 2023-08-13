package router

import (
	"commercetools-ms-product/controller"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Middleware
	// api := app.Group("/api", middleware.AuthReq())

	api := app.Group("/api/products")
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).
			JSON(map[string]interface{}{
				"health": "ok",
				"status": http.StatusOK,
			})
	})

	api.Get("/", controller.Find)

}
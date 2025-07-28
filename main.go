package main

import (
	"log"
	"ocr-server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/otiai10/gosseract/v2"
)

func main() {
	app := fiber.New(
		fiber.Config{
			EnablePrintRoutes: false,
		},
	)

	routes.OCR_CLIENT = gosseract.NewClient()
	routes.OCR_CLIENT.Languages = []string{"eng"}

	// Adding CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	// Adding logger middleware
	app.Use(logger.New())

	// Define your routes here
	app.Get("/", routes.Perform_ocr)
	app.Get("/byte", routes.Ocr_from_byte)
	log.Fatal(app.Listen(":7700"))
}

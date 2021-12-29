package main

import (
	"aftermath.link/repo/logs"
	"byvko.dev/repo/openmemorydb/database/driver"
	"byvko.dev/repo/openmemorydb/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	driver.OpenDatabase()
	defer driver.CloseDatabase() // Must be called at the end of the program to not lose data

	// Setup a server
	app := fiber.New()

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(logger.New())

	// Setup routes
	v1 := app.Group("/api/v1")

	// Create
	v1.Post("/create", handlers.CreateOneHandler)
	v1.Post("/createMany", handlers.CreateManyHandler)

	// Read
	v1.Get("/read", handlers.ReadOneHandler)
	v1.Get("/readMany", handlers.ReadManyHandler)

	// Update
	v1.Put("/update", handlers.UpdateOneHandler)
	v1.Put("/updateMany", handlers.UpdateManyHandler)

	// Patch -- TODO

	// Delete
	v1.Delete("/delete", handlers.DeleteOneHandler)
	v1.Delete("/deleteMany", handlers.DeleteManyHandler)

	logs.Fatal("Failed to start a server: %v", app.Listen(":3000"))
}

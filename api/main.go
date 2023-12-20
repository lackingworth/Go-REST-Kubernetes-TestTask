package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/lackingworth/Go-REST-API/routes"
	"github.com/lackingworth/Go-REST-API/database"
)

func main() {
	db, err := database.NewConnection(database.Dsn)

	if err != nil {
		log.Fatal("Error while connecting to database")
	}

	r := routes.Repository{
		DB: db,
	}

	app := fiber.New()

	r.SetupRoutes(app)
	app.Listen(":3000")
}
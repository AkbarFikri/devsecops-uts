package main

import (
	"github.com/AkbarFikri/devsecops-uts/entity"
	"github.com/AkbarFikri/devsecops-uts/pkg/database"
	"github.com/AkbarFikri/devsecops-uts/src/handler"
	"github.com/AkbarFikri/devsecops-uts/src/repository"
	"github.com/AkbarFikri/devsecops-uts/src/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())
	app.Get("", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	db, err := database.NewInstance()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&entity.User{})

	repo := repository.NewUserRepository(db)
	authService := service.NewAuthService(repo)
	handler.NewAuthHandler(app.Group("/api"), authService)

	app.Listen(":3000")
}

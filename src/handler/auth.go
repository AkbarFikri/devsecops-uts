package handler

import (
	"github.com/AkbarFikri/devsecops-uts/src/service"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(app fiber.Router, authService service.AuthService) {
	authHandler := &authHandler{authService: authService}

	app.Post("/login", authHandler.Login)
}

func (authHandler *authHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	res, err := authHandler.authService.Login(req.Username, req.Password)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendString(res)
}

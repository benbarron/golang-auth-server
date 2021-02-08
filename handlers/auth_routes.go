package handlers

import (
	"github.com/benbarron/UserMicroService/services"
	"github.com/gofiber/fiber/v2"
)

type AuthRoutes struct {
	AuthService *services.AuthService
	Logger *services.LoggingService
	LocalsService *services.LocalsStorage
}

type LoginRequest struct {
	Username 	string 		`json:"username"`
	Password 	string 		`json:"password"`
}

func NewAuthRoutes() *AuthRoutes {
	return &AuthRoutes{
		Logger:      services.NewLoggingService("AuthRoutes"),
		AuthService: services.NewAuthService(),
		LocalsService: services.NewLocalsStorage(),
	}
}

func (r *AuthRoutes) Login(ctx *fiber.Ctx) error {
	request := new(LoginRequest)
	ctx.BodyParser(request)
	user, accessToken, refreshToken, err := r.AuthService.Login(request.Username, request.Password)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"user": user,
		"access-token": accessToken,
		"refresh-token": refreshToken,
	})
}

func (r *AuthRoutes) ValidateUser(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"user": r.LocalsService.GetUser(ctx),
	})
}

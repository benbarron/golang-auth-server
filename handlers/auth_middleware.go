package handlers

import (
	"github.com/benbarron/UserMicroService/database"
	"github.com/benbarron/UserMicroService/services"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	db := database.GetDatabase()
	jwtService := services.NewJwtService()
	localService := services.NewLocalsStorage()
	accessToken := ctx.Get("access-token")
	refreshToken := ctx.Get("refresh-token")

	if claims, err := jwtService.ValidateToken(accessToken); err == nil {
		localService.SetUser(ctx, claims.User)
		return ctx.Next()
	}

	if claims, err := jwtService.ValidateToken(refreshToken); err == nil {
		var user database.User
		res := db.Where("uid = ?", claims.User.Uid).First(&user)
		if res.Error != nil || user.TokenStep != claims.User.TokenStep {
			return ctx.Status(401).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		newAccessToken, _ := jwtService.GenerateAccessToken(claims.User)
		ctx.Response().Header.Add("access-token", newAccessToken)
		localService.SetUser(ctx, claims.User)
		return ctx.Next()

	}

	return ctx.Status(401).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}

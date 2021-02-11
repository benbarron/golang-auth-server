package handlers

import (
	"fmt"
	"github.com/benbarron/golang-auth-server/services"
	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware(ctx *fiber.Ctx) error {
	logger := services.NewLoggingService("LoggerMiddleware")
	method := ctx.Method()
	path := ctx.Path()
	logger.Log(fmt.Sprintf("%s - %s", method, path))
	return ctx.Next()
}

package handlers

import (
	"github.com/benbarron/golang-auth-server/database"
	"github.com/benbarron/golang-auth-server/services"
	"github.com/gofiber/fiber/v2"
)

/**
 *
 */
type UserRoutes struct {
	UserService *services.UserService
	Logger *services.LoggingService
}

/**
 *
 */
func NewUserRoutes() UserRoutes {
	return UserRoutes{
		Logger:      services.NewLoggingService("UserService"),
		UserService: services.NewUserService(),
	}
}

/**
 *
 */
func (r *UserRoutes) GetAllUsers(ctx *fiber.Ctx) error {
	users, err := r.UserService.FindAll()

	if err != nil {
		r.Logger.Log("Error fetching users")
		return ctx.Status(501).JSON(fiber.Map{
			"error": "Error fetching users",
		})
	}

	r.Logger.Log("Test")
	return ctx.Status(200).JSON(fiber.Map{
		"users": users,
	})
}

/**
 *
 */
func (r *UserRoutes) GetUserById(ctx *fiber.Ctx) error {
	user, err := r.UserService.FindById(ctx.Params("id"))

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "User not found.",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"providers": user,
	})
}

/**
 *
 */
func (r *UserRoutes) CreateUser(ctx *fiber.Ctx) error {
	user := new(database.User)
	err := ctx.BodyParser(user)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	err = r.UserService.CreateUser(user)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"providers": user,
	})
}

/**
 *
 */
func (r *UserRoutes) DeleteById(ctx *fiber.Ctx) error {
	err := r.UserService.DeleteById(ctx.Params("id"))

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "User deleted",
	})
}

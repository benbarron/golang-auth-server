package main

import (
	"fmt"
	"github.com/benbarron/UserMicroService/handlers"
	"github.com/benbarron/UserMicroService/database"
	"github.com/benbarron/UserMicroService/services"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func RegisterHandlers(app *fiber.App) {
	/**
	 *	User Routes
	 */
	userRoutes := handlers.NewUserRoutes()
	app.Get("/api/users", handlers.AuthMiddleware, userRoutes.GetAllUsers)
	app.Get("/api/users/:id", handlers.AuthMiddleware, userRoutes.GetUserById)
	app.Post("/api/users", handlers.AuthMiddleware, userRoutes.CreateUser)
	app.Delete("/api/users/:id", handlers.AuthMiddleware, userRoutes.DeleteById)

	/**
	 *	Auth Routes
	 */
	authRoutes := handlers.NewAuthRoutes()
	app.Post("/api/auth/login", authRoutes.Login)
	app.Get("/api/auth/validate", handlers.AuthMiddleware, authRoutes.ValidateUser)
}

func RegisterAppMiddleware(app *fiber.App) {
	app.Use(handlers.LoggerMiddleware)
}

func CreateApp() *fiber.App {
	appConfig := fiber.Config{}
	return fiber.New(appConfig)
}

func ServeApp(app *fiber.App) {
	port := os.Getenv("SERVER_PORT")
	host := os.Getenv("SERVER_HOST")
	logger := services.NewLoggingService("Main")
	logger.Log(fmt.Sprintf("API live at http://%s:%s\n", host, port))
	log.Fatal(app.Listen(":" + port))
}

func main() {
	godotenv.Load(".env")
	database.MigrateDatabase()
	app := CreateApp()
	RegisterAppMiddleware(app)
	RegisterHandlers(app)
	ServeApp(app)
}

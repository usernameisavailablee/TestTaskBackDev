package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/usernameisavailablee/TestTaskBackDev/handlers"
)

func setupRoutes(app *fiber.App) {

    app.Post("/user", handlers.CreateUser)

    app.Post("/token", handlers.CreateToken)
}

package main

import (
    "github.com/gofiber/fiber/v2"
    authHandlers "github.com/usernameisavailablee/TestTaskBackDev/auth/handlers"
    generalHandlers "github.com/usernameisavailablee/TestTaskBackDev/handlers"

)

func setupRoutes(app *fiber.App) {

    app.Post("/user", generalHandlers.CreateUser)

    app.Post("/auth/generate-pair", authHandlers.GenerateTokenPair)
    app.Post("/auth/refresh", authHandlers.RefreshToken)
}

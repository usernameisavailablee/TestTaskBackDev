package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/usernameisavailablee/TestTaskBackDev/database"
    "github.com/usernameisavailablee/TestTaskBackDev/models"
)

// CreateToken создает новый токен
func CreateToken(c *fiber.Ctx) error {
    token := new(models.Token)

    if err := c.BodyParser(token); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid request body: " + err.Error(),
        })
    }

    if result := database.DB.Db.Create(&token); result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Failed to create token: " + result.Error.Error(),
        })
    }

    return c.Status(fiber.StatusCreated).JSON(token)
}

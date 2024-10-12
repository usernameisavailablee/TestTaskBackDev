package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/usernameisavailablee/TestTaskBackDev/database"
    "github.com/usernameisavailablee/TestTaskBackDev/models"
    "errors"
    "gorm.io/gorm"
)

func CreateUser(c *fiber.Ctx) error {
    user := new(models.User)

    if err := c.BodyParser(user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid request body: " + err.Error(),
        })
    }

    if result := database.DB.Db.Create(&user); result.Error != nil {
        if errors.Is(result.Error, gorm.ErrDuplicatedKey) ||
           result.Error.Error() == "ERROR: duplicate key value violates unique constraint \"uni_users_email\" (SQLSTATE 23505)" {
            return c.Status(fiber.StatusConflict).JSON(fiber.Map{
                "message": "Email already exists",
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Failed to create user: " + result.Error.Error(),
        })
    }

    return c.Status(fiber.StatusCreated).JSON(user)
}

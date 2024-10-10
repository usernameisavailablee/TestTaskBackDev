package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/usernameisavailablee/TestTaskBackDev/database"
    "github.com/usernameisavailablee/TestTaskBackDev/models"
    "github.com/usernameisavailablee/TestTaskBackDev/auth"
    "github.com/google/uuid"
)

type TokenRequest struct {
    UserID string `json:"user_id"`
}

func GenerateTokenPair(c *fiber.Ctx) error {
    if c.Method() != fiber.MethodPost {
        return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{
            "message": "Method not allowed",
        })
    }

    var request TokenRequest

    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid request format",
        })
    }

    if request.UserID == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "User ID is required",
        })
    }

    userID, err := uuid.Parse(request.UserID)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid User ID format",
        })
    }

    var user models.User
    if err := database.DB.Db.First(&user, userID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "User not found",
        })
    }

    ip := c.IP()

    accessToken, err := auth.GenerateAccessToken(userID.String(), ip)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Failed to generate access token",
        })
    }

    refreshToken, hashedRefreshToken, err := auth.GenerateRefreshToken()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Failed to generate refresh token",
        })
    }

    token := models.Token{
        UserID:    userID,
        Refresh:   hashedRefreshToken,
        IPAddress: ip,
    }
    if err := database.DB.Db.Create(&token).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Failed to store refresh token",
        })
    }

    return c.JSON(fiber.Map{
        "access_token":  accessToken,
        "refresh_token": refreshToken,
    })
}

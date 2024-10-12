package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/usernameisavailablee/TestTaskBackDev/database"
    "github.com/usernameisavailablee/TestTaskBackDev/models"
    "github.com/usernameisavailablee/TestTaskBackDev/auth"
    "log"
)

type RefreshTokenRequest struct {
    RefreshToken string `json:"refresh_token"`
    UserID       string `json:"user_id"`
}

func RefreshToken(c *fiber.Ctx) error {
    var request RefreshTokenRequest

    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid request format",
        })
    }

    if request.RefreshToken == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Refresh token is required",
        })
    }

    if request.UserID == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "User ID is required",
        })
    }

    ip := auth.GetClientIP(c)
    log.Printf("Client IP: %s", ip)

    var token models.Token
    if err := database.DB.Db.Where("user_id = ?", request.UserID).First(&token).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "Token or user not found",
        })
    }
    var user models.User
    if err := database.DB.Db.Where("id = ?", request.UserID).First(&user).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "User not found",
        })
    }


    // Проверка изменения IP-адреса и отправка email предупреждения
    if token.IPAddress != "" && token.IPAddress != ip {
        userEmail := user.Email
        auth.SendEmailWarning(userEmail, token.IPAddress, ip)
    }

    if err := auth.ValidateRefreshToken(request.RefreshToken, token.Refresh); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "Invalid refresh token",
        })
    }

    newAccessToken, err := auth.GenerateAccessToken(request.UserID, ip)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Failed to generate new access token",
        })
    }

    newRefreshToken, newHashedRefreshToken, err := auth.GenerateRefreshToken()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Failed to generate new refresh token",
        })
    }

    token.Refresh = newHashedRefreshToken
    token.IPAddress = ip
    if err := database.DB.Db.Save(&token).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Failed to update refresh token in the database",
        })
    }

    return c.JSON(fiber.Map{
        "access_token":  newAccessToken,
        "refresh_token": newRefreshToken,
    })
}

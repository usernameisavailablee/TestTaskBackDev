package auth

import (
    "crypto/sha512"
    "crypto/rand"
    "encoding/base64"
    "time"
    "os"
    "log"

    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
    "github.com/gofiber/fiber/v2"
)

func GenerateAccessToken(userID, ip string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "ip":      ip,
        "exp":     time.Now().Add(time.Minute * 15).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

    secretKey := os.Getenv("JWT_SECRET")
    if secretKey == "" {
        log.Fatal("JWT_SECRET is not set in the .env file")
    }

    secret := sha512.Sum512([]byte(secretKey))

    return token.SignedString(secret[:])
}

func GenerateRefreshToken() (string, string, error) {
    token := make([]byte, 32)
    _, err := rand.Read(token)
    if err != nil {
        return "", "", err
    }

    refreshToken := base64.StdEncoding.EncodeToString(token)
    hashedToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
    if err != nil {
        return "", "", err
    }

    return refreshToken, string(hashedToken), nil
}

func ValidateRefreshToken(refreshToken string, hashedToken string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedToken), []byte(refreshToken))
}


func SendEmailWarning(email, oldIP, newIP string) {
    log.Printf("Warning: IP address has changed from %s to %s for user with email %s", oldIP, newIP, email)
}

func GetClientIP(c *fiber.Ctx) string {

    log.Printf("Headers:", c.GetReqHeaders())

    forwarded := c.Get("X-Forwarded-For")
    if forwarded != "" {
        return forwarded
    }
    return c.IP()
}

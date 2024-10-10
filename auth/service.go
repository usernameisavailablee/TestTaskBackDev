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

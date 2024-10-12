package models

import (
    "gorm.io/gorm"
    "github.com/google/uuid"
)

type User struct {
    gorm.Model
    ID     uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
    Email    string `gorm:"unique"`
}

type Token struct {
    gorm.Model
    ID         uint   `gorm:"primaryKey"`
    UserID     uuid.UUID `gorm:"index"`
    Refresh    string
    IPAddress  string
}

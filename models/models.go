package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    ID       string `gorm:"primaryKey"`
    Email    string
}

type Token struct {
    gorm.Model
    ID         uint   `gorm:"primaryKey"`
    UserID     string `gorm:"index"`
    Refresh    string
    IPAddress  string 
}

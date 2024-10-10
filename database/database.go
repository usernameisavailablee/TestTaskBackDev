
package database

import (
    "fmt"
    "log"
    "os"

    "github.com/usernameisavailablee/TestTaskBackDev/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

type Dbinstance struct {
    Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
    dsn := fmt.Sprintf(
        "host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Moscow",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })

    if err != nil {
        log.Fatal("Failed to connect to database. \n", err)
        os.Exit(2)
    }

    log.Println("connected")

    if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
    log.Fatal("Failed to create uuid-ossp extension: ", err)
    }

    db.Logger = logger.Default.LogMode(logger.Info)

    log.Println("running migrations")
    if err := db.AutoMigrate(&models.User{}, &models.Token{}); err != nil {
        log.Fatal("Migration failed: ", err)
    }

    DB = Dbinstance{
        Db: db,
    }
}

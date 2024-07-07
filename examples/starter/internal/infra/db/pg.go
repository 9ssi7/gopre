package db

import (
	"fmt"
	"log"
	"os"

	"github.com/9ssi7/gopre-starter/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error

	cfg := config.ReadValue()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.Port, cfg.Database.SslMode)

	DB, err = gorm.Open(
		postgres.New(
			postgres.Config{
				DSN: dsn,
			},
		),
	)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		os.Exit(1)
	}

	log.Println("Connected to database!")

	return DB
}

func GetDB() *gorm.DB {
	return DB
}

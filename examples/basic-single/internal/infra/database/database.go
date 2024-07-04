package database

import (
	"fmt"
	"log"
	"os"

	"github.com/9ssi7/gopre/basic-single/internal/domain/aggregates"
	"github.com/9ssi7/gopre/basic-single/internal/domain/entities"
	"github.com/9ssi7/gopre/basic-single/internal/infra/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	cfg := config.ReadValue()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.Port, cfg.Database.SslMode)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		os.Exit(1)
	}

	// Migrate the schema
	err = DB.AutoMigrate(
		&aggregates.Order{},
		&entities.Product{},
		&entities.OrderItem{},
	)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
		os.Exit(1)
	}

	log.Println("Connected to database!")
}

func GetDB() *gorm.DB {
	return DB
}

package seeds

import (
	"log"

	"github.com/9ssi7/gopre/basic-single/internal/domain/entities"
	"github.com/9ssi7/gopre/basic-single/internal/infra/config"
	"github.com/9ssi7/gopre/basic-single/internal/infra/database"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func SeedData() {
	cfg := config.ReadValue()

	if !cfg.RunSeed {
		log.Println("Seed data disabled in configuration. Skipping seeding...")
		return
	}

	db := database.GetDB()

	// Seed Products
	products := []entities.Product{
		{ID: uuid.New(), Name: "Product A", Description: "Description for Product A", Price: decimal.NewFromFloat(19.99), Stock: 100},
		{ID: uuid.New(), Name: "Product B", Description: "Description for Product B", Price: decimal.NewFromFloat(29.99), Stock: 50},
		// ... add more products as needed
	}

	if err := db.Create(&products).Error; err != nil {
		log.Fatalf("Failed to seed products: %v", err)
	}

	log.Println("Successfully seeded products!")
}

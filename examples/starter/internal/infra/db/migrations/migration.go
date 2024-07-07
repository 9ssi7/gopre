package migrations

import (
	"github.com/9ssi7/gopre-starter/internal/domain/entities"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	db.AutoMigrate(&entities.Contact{}, &entities.User{})
}

package seeds

import (
	"time"

	"github.com/9ssi7/gopre-starter/config"
	"github.com/9ssi7/gopre-starter/internal/domain/entities"
	"github.com/9ssi7/gopre-starter/pkg/ptr"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

func runUserSeeds(db *gorm.DB) {
	var user entities.User
	if err := db.Model(&entities.User{}).Where("email = ?", "test@test.com").First(&user).Error; err != nil {
		db.Create(&entities.User{
			Email: "test@test.com",
			Name:  "Test",
			Roles: pq.StringArray{
				config.Roles.Admin,
				config.Roles.AdminSuper,
			},
			IsActive:   true,
			VerifiedAt: ptr.Time(time.Now()),
		})
	}
}

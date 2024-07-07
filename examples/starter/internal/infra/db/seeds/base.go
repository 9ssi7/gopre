package seeds

import "gorm.io/gorm"

func Run(db *gorm.DB) {
	runUserSeeds(db)
}

package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	Id        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (g *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if g.Id == uuid.Nil {
		g.Id = uuid.New()
	}
	return nil
}

package entities

import (
	"time"

	"github.com/9ssi7/gopre/internal/domain/valobj"
	"github.com/9ssi7/gopre/pkg/ptr"
	"github.com/google/uuid"
)

type User struct {
	Base
	Name       string           `json:"name" gorm:"type:varchar(255);not null"`
	Email      string           `json:"email" gorm:"type:varchar(255);not null;unique"`
	IsActive   bool             `json:"is_active" gorm:"type:boolean;not null;default:true"`
	RoleIds    valobj.UUIDArray `json:"role_ids" gorm:"type:text[]"`
	TempToken  *string          `json:"temp_token" gorm:"type:varchar(255);default:null;index:idx_verifier"`
	VerifiedAt *time.Time       `json:"verified_at" gorm:"type:timestamp;default:null;index:idx_verifier"`
}

func (u *User) AddRole(roleId uuid.UUID) {
	if u.CheckRole(roleId) {
		return
	}
	u.RoleIds = append(u.RoleIds, roleId)
}

func (u *User) CheckRole(roleId uuid.UUID) bool {
	for _, r := range u.RoleIds {
		if r == roleId {
			return true
		}
	}
	return false
}

func (u *User) RemoveRole(roleId uuid.UUID) {
	for i, r := range u.RoleIds {
		if r == roleId {
			u.RoleIds = append(u.RoleIds[:i], u.RoleIds[i+1:]...)
			break
		}
	}
}

func (u *User) Verify() {
	u.VerifiedAt = ptr.Time(time.Now())
	u.TempToken = nil
}

func (u *User) Enable() {
	u.IsActive = true
}

func (u *User) Disable() {
	u.IsActive = false
}

func NewUser(name string, email string) *User {
	return &User{
		Name:      name,
		Email:     email,
		IsActive:  true,
		RoleIds:   valobj.UUIDArray{},
		TempToken: ptr.String(uuid.New().String()),
	}
}

package entities

import "github.com/lib/pq"

type Role struct {
	Base
	Name        string         `json:"name" gorm:"type:varchar(255);not null"`
	Description string         `json:"description" gorm:"type:text;default:null"`
	IsActive    bool           `json:"is_active" gorm:"type:boolean;not null;default:true"`
	Claims      pq.StringArray `json:"claims" gorm:"type:text[]"`
}

func (r *Role) AddClaim(claim string) {
	if r.CheckClaim(claim) {
		return
	}
	r.Claims = append(r.Claims, claim)
}

func (r *Role) CheckClaim(claim string) bool {
	for _, c := range r.Claims {
		if c == claim {
			return true
		}
	}
	return false
}

func (r *Role) RemoveClaim(claim string) {
	for i, c := range r.Claims {
		if c == claim {
			r.Claims = append(r.Claims[:i], r.Claims[i+1:]...)
			break
		}
	}
}

func (r *Role) Enable() {
	r.IsActive = true
}

func (r *Role) Disable() {
	r.IsActive = false
}

func NewRole(name string, description string) *Role {
	return &Role{
		Name:        name,
		Description: description,
		Claims:      pq.StringArray{},
		IsActive:    true,
	}
}

package valobj

import "github.com/google/uuid"

type Audit struct {
	UpdatedBy *uuid.UUID `json:"updated_by" gorm:"default:null" example:"f0a0a1e9-69bd-4bef-b8c6-4e8c0d3a1212"`
	DeletedBy *uuid.UUID `json:"deleted_by" gorm:"default:null" example:"f0a0a1e9-69bd-4bef-b8c6-4e8c0d3a1212"`
	MakedBy   *uuid.UUID `json:"maked_by" gorm:"default:null" example:"f0a0a1e9-69bd-4bef-b8c6-4e8c0d3a1212"`
}

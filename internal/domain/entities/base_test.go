package entities

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestBase_BeforeCreate(t *testing.T) {
	testCases := []struct {
		name      string
		initialID uuid.UUID
		expected  uuid.UUID
	}{
		{
			name:      "Nil ID",
			initialID: uuid.Nil,
			expected:  uuid.UUID{}, // Expect a new UUID to be generated
		},
		{
			name:      "Existing ID",
			initialID: uuid.New(),
			expected:  uuid.UUID{}, // Expect no change to the ID
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			base := &Base{Id: tc.initialID}

			// Create a mock GORM DB instance
			db, _ := gorm.Open(nil, nil)
			tx := db.Begin()
			defer tx.Rollback()

			err := base.BeforeCreate(tx)
			assert.NoError(t, err)

			if tc.initialID == uuid.Nil {
				// When the ID was initially nil, a new one should be generated
				assert.NotEqual(t, uuid.Nil, base.Id)
			} else {
				// When the ID was not nil, it should remain the same
				assert.Equal(t, tc.initialID, base.Id)
			}
		})
	}
}

package entities

import (
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestRole_AddClaim(t *testing.T) {
	testCases := []struct {
		name     string
		role     *Role
		claim    string
		expected pq.StringArray
	}{
		{
			name:     "Add new claim",
			role:     &Role{Claims: pq.StringArray{"claim1"}},
			claim:    "claim2",
			expected: pq.StringArray{"claim1", "claim2"},
		},
		{
			name:     "Add existing claim (no change)",
			role:     &Role{Claims: pq.StringArray{"claim1"}},
			claim:    "claim1",
			expected: pq.StringArray{"claim1"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.role.AddClaim(tc.claim)
			assert.Equal(t, tc.expected, tc.role.Claims)
		})
	}
}

func TestRole_RemoveClaim(t *testing.T) {
	testCases := []struct {
		name     string
		role     *Role
		claim    string
		expected pq.StringArray
	}{
		{
			name:     "Remove existing claim",
			role:     &Role{Claims: pq.StringArray{"claim1", "claim2"}},
			claim:    "claim1",
			expected: pq.StringArray{"claim2"},
		},
		{
			name:     "Remove non-existent claim (no change)",
			role:     &Role{Claims: pq.StringArray{"claim1"}},
			claim:    "claim3",
			expected: pq.StringArray{"claim1"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.role.RemoveClaim(tc.claim)
			assert.Equal(t, tc.expected, tc.role.Claims)
		})
	}
}

func TestRole_Enable(t *testing.T) {
	role := &Role{IsActive: false}
	role.Enable()
	assert.True(t, role.IsActive)
}

func TestRole_Disable(t *testing.T) {
	role := &Role{IsActive: true}
	role.Disable()
	assert.False(t, role.IsActive)
}

func TestNewRole(t *testing.T) {
	name := "test_role"
	description := "This is a test role"
	role := NewRole(name, description)

	assert.Equal(t, name, role.Name)
	assert.Equal(t, description, role.Description)
	assert.Empty(t, role.Claims)
	assert.True(t, role.IsActive)
}

func TestCheckClaim(t *testing.T) {
	role := &Role{Claims: pq.StringArray{"claim1", "claim2"}}

	assert.True(t, role.CheckClaim("claim1"))
	assert.True(t, role.CheckClaim("claim2"))
	assert.False(t, role.CheckClaim("claim3"))
}

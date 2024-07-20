package entities

import (
	"testing"

	"github.com/9ssi7/gopre/internal/domain/valobj"
	"github.com/9ssi7/gopre/pkg/ptr"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUser_AddRole(t *testing.T) {
	staticId := uuid.New()
	staticIdSecond := uuid.New()
	testCases := []struct {
		name         string
		initialRoles valobj.UUIDArray
		roleId       uuid.UUID
		expected     valobj.UUIDArray
	}{
		{"Add to empty", valobj.UUIDArray{}, staticId, valobj.UUIDArray{staticId}},
		{"Add new role", valobj.UUIDArray{staticId}, staticIdSecond, valobj.UUIDArray{staticId, staticIdSecond}},
		{"Add existing role (no change)", valobj.UUIDArray{staticId}, staticId, valobj.UUIDArray{staticId}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user := &User{RoleIds: tc.initialRoles}
			user.AddRole(tc.roleId)
			assert.Equal(t, tc.expected, user.RoleIds)
		})
	}
}

func TestUser_RemoveRole(t *testing.T) {
	roleId1 := uuid.New()
	roleId2 := uuid.New()
	testCases := []struct {
		name         string
		initialRoles valobj.UUIDArray
		roleId       uuid.UUID
		expected     valobj.UUIDArray
	}{
		{"Remove from empty", valobj.UUIDArray{}, roleId1, valobj.UUIDArray{}},
		{"Remove existing role", valobj.UUIDArray{roleId1, roleId2}, roleId1, valobj.UUIDArray{roleId2}},
		{"Remove non-existent role (no change)", valobj.UUIDArray{roleId1}, uuid.New(), valobj.UUIDArray{roleId1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user := &User{RoleIds: tc.initialRoles}
			user.RemoveRole(tc.roleId)
			assert.Equal(t, tc.expected, user.RoleIds)
		})
	}
}

func TestUser_Verify(t *testing.T) {
	user := &User{TempToken: ptr.String("temp_token")}
	user.Verify()

	assert.NotNil(t, user.VerifiedAt)
	assert.Nil(t, user.TempToken)
}

func TestUser_EnableDisable(t *testing.T) {
	user := &User{IsActive: false}

	user.Enable()
	assert.True(t, user.IsActive)

	user.Disable()
	assert.False(t, user.IsActive)
}

func TestNewUser(t *testing.T) {
	name := "John Doe"
	email := "johndoe@example.com"
	user := NewUser(name, email)

	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)
	assert.Empty(t, user.RoleIds)
	assert.NotNil(t, user.TempToken)
	assert.True(t, user.IsActive)

	// Since Base has a BeforeCreate hook, we can't test the ID here directly
	// but we can ensure it's not nil after the hook would run:
	user.BeforeCreate(nil)
	assert.NotEqual(t, uuid.Nil, user.Id)
}

func TestUserCheckRole(t *testing.T) {
	roleId1 := uuid.New()
	roleId2 := uuid.New()
	user := &User{RoleIds: valobj.UUIDArray{roleId1, roleId2}}

	assert.True(t, user.CheckRole(roleId1))
	assert.True(t, user.CheckRole(roleId2))
	assert.False(t, user.CheckRole(uuid.New()))
}

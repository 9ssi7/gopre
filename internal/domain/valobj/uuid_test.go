package valobj

import (
	"database/sql/driver"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert" // Or use your preferred assertion library
)

func TestUUIDArray_Value(t *testing.T) {
	firstId := uuid.New()
	secondId := uuid.New()
	testCases := []struct {
		name     string
		input    UUIDArray
		expected driver.Value
		wantErr  bool
	}{
		{
			name:     "Empty array",
			input:    UUIDArray{},
			expected: nil,
			wantErr:  false,
		},
		{
			name:     "Valid UUID array",
			input:    UUIDArray{firstId, secondId},
			expected: `["` + firstId.String() + `","` + secondId.String() + `"]`, // Example JSON string
			wantErr:  false,
		},
		// You can add more test cases here if you want to simulate JSON marshalling errors.
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := tc.input.Value()
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, val)
			}
		})
	}
}

func TestUUIDArray_Scan(t *testing.T) {
	ids := []uuid.UUID{uuid.New(), uuid.New()}
	testCases := []struct {
		name     string
		input    interface{}
		expected UUIDArray
		wantErr  bool
	}{
		{
			name:     "Scan from []byte",
			input:    []byte(`["` + ids[0].String() + `","` + ids[1].String() + `"]`),
			expected: UUIDArray{ids[0], ids[1]},
			wantErr:  false,
		},
		{
			name:     "Scan from string",
			input:    `["` + ids[0].String() + `","` + ids[1].String() + `"]`,
			expected: UUIDArray{ids[0], ids[1]},
			wantErr:  false,
		},
		{
			name:    "Unsupported type",
			input:   123,
			wantErr: true,
		},
		// You can add more test cases here to simulate invalid JSON data.
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var p UUIDArray
			err := p.Scan(tc.input)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, p)
			}
		})
	}
}

func TestUUIDArray_ToStringArray(t *testing.T) {
	uuids := UUIDArray{uuid.New(), uuid.New()}
	expected := []string{uuids[0].String(), uuids[1].String()}

	result := uuids.ToStringArray()

	assert.Equal(t, expected, result)
}

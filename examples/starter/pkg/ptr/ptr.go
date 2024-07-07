package ptr

import (
	"time"

	"github.com/google/uuid"
)

// Int returns a pointer to the given int.
// Example: ptr.Int(42)
// Returns: *int
func UUID(id uuid.UUID) *uuid.UUID {
	return &id
}

// String returns a pointer to the given string.
// Example: ptr.String("hello")
// Returns: *string
func String(s string) *string {
	return &s
}

// Time returns a pointer to the given time.
// Example: ptr.Time(time.Now())
// Returns: *time.Time
func Time(t time.Time) *time.Time {
	return &t
}

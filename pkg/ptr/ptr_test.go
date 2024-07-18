package ptr

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestUUID(t *testing.T) {
	// Test with a valid UUID
	id := uuid.New()
	ptr := UUID(id)
	if *ptr != id {
		t.Errorf("UUID pointer does not match original: got %v, want %v", *ptr, id)
	}
}

func TestString(t *testing.T) {
	// Test with a non-empty string
	str := "hello"
	ptr := String(str)
	if *ptr != str {
		t.Errorf("String pointer does not match original: got %v, want %v", *ptr, str)
	}

	// Test with an empty string
	emptyStr := ""
	emptyPtr := String(emptyStr)
	if *emptyPtr != emptyStr {
		t.Errorf("String pointer for empty string does not match original")
	}
}

func TestTime(t *testing.T) {
	// Test with current time
	now := time.Now()
	ptr := Time(now)
	if !ptr.Equal(now) { // Use time.Equal for comparison
		t.Errorf("Time pointer does not match original: got %v, want %v", ptr, now)
	}
}

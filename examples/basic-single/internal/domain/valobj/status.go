package valobj

import (
	"fmt"
	"strings"
)

// Status represents the status of an order.
type Status string

const (
	StatusPending    Status = "pending"
	StatusProcessing Status = "processing"
	StatusShipped    Status = "shipped"
	StatusDelivered  Status = "delivered"
	StatusCancelled  Status = "cancelled"
)

// ValidStatuses is a slice of all valid order statuses.
var ValidStatuses = []Status{StatusPending, StatusProcessing, StatusShipped, StatusDelivered, StatusCancelled}

// NewStatus creates a new Status value object.
func NewStatus(value string) (Status, error) {
	normalizedValue := Status(strings.ToLower(strings.TrimSpace(value)))

	for _, validStatus := range ValidStatuses {
		if validStatus == normalizedValue {
			return normalizedValue, nil
		}
	}

	return "", fmt.Errorf("invalid order status: %s", value)
}

// String returns the string representation of the Status.
func (s Status) String() string {
	return string(s)
}

// IsValid checks if the Status is valid.
func (s Status) IsValid() bool {
	for _, validStatus := range ValidStatuses {
		if validStatus == s {
			return true
		}
	}
	return false
}

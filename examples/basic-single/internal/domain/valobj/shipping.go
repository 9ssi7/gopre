package valobj

import "time"

type ShippingInfo struct {
    RecipientName    string   `json:"recipientName"`
    Address          Address  `json:"address"`
    ShippingMethod   string   `json:"shippingMethod"`
    TrackingNumber   string   `json:"trackingNumber,omitempty"` // Optional
    EstimatedDeliveryDate time.Time `json:"estimatedDeliveryDate,omitempty"` // Optional
}

// NewShippingInfo creates a new ShippingInfo value object.
func NewShippingInfo(recipientName string, address Address, shippingMethod string) (ShippingInfo, error) {
    // Add validation logic here (e.g., check if recipientName is not empty, address is valid, etc.)

    return ShippingInfo{
        RecipientName:    recipientName,
        Address:          address,
        ShippingMethod:   shippingMethod,
    }, nil
}

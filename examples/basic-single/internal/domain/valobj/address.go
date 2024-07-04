package valobj

import "errors"

// Address represents a physical address.
type Address struct {
    Street     string `json:"street"`
    City       string `json:"city"`
    State      string `json:"state"`
    PostalCode string `json:"postalCode"`
    Country    string `json:"country"`
}

// NewAddress creates a new Address value object.
func NewAddress(street, city, state, postalCode, country string) (Address, error) {
    // Add validation logic here (e.g., check if fields are not empty, etc.)
    if street == "" || city == "" || state == "" || postalCode == "" || country == "" {
        return Address{}, errors.New("all address fields are required")
    }

    return Address{
        Street:     street,
        City:       city,
        State:      state,
        PostalCode: postalCode,
        Country:    country,
    }, nil
}

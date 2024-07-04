package valobj

type BillingInfo struct {
	BillingAddress Address `json:"billingAddress"`
	PaymentMethod  string  `json:"paymentMethod"`
	// Add more fields as needed (e.g., card number, expiration date, etc.)
}

// NewBillingInfo creates a new BillingInfo value object.
func NewBillingInfo(billingAddress Address, paymentMethod string) (BillingInfo, error) {
	// Add validation logic here (e.g., check if billingAddress is valid, paymentMethod is supported, etc.)

	return BillingInfo{
		BillingAddress: billingAddress,
		PaymentMethod:  paymentMethod,
	}, nil
}

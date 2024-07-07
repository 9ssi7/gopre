package validation

type ErrorResponse struct {

	// Field is the field name.
	Field string `json:"field"`

	// Message is the error message.
	Message string `json:"message"`

	// Namespace is the namespace of the error.
	Namespace string `json:"namespace,omitempty"`

	// Value is the value of the field.
	Value interface{} `json:"value"`
}

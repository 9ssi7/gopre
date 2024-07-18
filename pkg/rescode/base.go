package rescode

import "net/http"

var (
	ValidationFailed = New(1000, http.StatusUnprocessableEntity, "validation_failed")
)

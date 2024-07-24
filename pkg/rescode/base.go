package rescode

import "net/http"

var (
	ValidationFailed = New(1000, http.StatusUnprocessableEntity, "validation_failed")
	Failed           = New(1001, http.StatusInternalServerError, "failed")
	NotFound         = New(1002, http.StatusNotFound, "not_found")
)

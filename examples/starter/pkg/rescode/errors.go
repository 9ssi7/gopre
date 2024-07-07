package rescode

import "net/http"

var (
	Failed   = New(codeFailed, msgFailed, nil)
	NotFound = New(codeNotFound, msgNotFound, nil, Extra{
		HttpStatus:    http.StatusNotFound,
		Translateable: true,
	})
	UserDisabled = New(codeUserDisabled, msgUserDisabled, nil, Extra{
		HttpStatus:    http.StatusForbidden,
		Translateable: true,
	})
	UserVerifyRequired = New(codeUserVerifyRequired, msgUserVerifyRequired, nil, Extra{
		HttpStatus:    http.StatusForbidden,
		Translateable: true,
	})
	VerificationExpired = New(codeVerificationExpired, msgVerificationExpired, map[string]interface{}{
		"isExpired": true,
	}, Extra{
		HttpStatus:    http.StatusForbidden,
		Translateable: true,
	})
	VerificationExceeded = New(codeVerificationExceeded, MsfVerificationExceeded, map[string]interface{}{
		"isExceeded": true,
	}, Extra{
		HttpStatus:    http.StatusForbidden,
		Translateable: true,
	})
	VerificationInvalid = New(codeVerificationInvalid, msgVerificationInvalid, map[string]interface{}{
		"isInvalid": true,
	}, Extra{
		HttpStatus:    http.StatusForbidden,
		Translateable: true,
	})
	InvalidRefreshOrAccessTokens = New(codeInvalidRefreshOrAccessTokens, msgInvalidRefreshOrAccessTokens, nil, Extra{
		HttpStatus:    http.StatusForbidden,
		Translateable: true,
	})
	InvalidOrExpiredToken = New(codeInvalidOrExpiredToken, msgInvalidOrExpiredToken, map[string]interface{}{
		"isExpired": true,
	}, Extra{
		HttpStatus:    http.StatusForbidden,
		Translateable: true,
	})
	InvalidAccess = New(codeInvalidAccess, msgInvalidAccess, nil, Extra{
		HttpStatus:    http.StatusForbidden,
		Translateable: true,
	})
	InvalidRefreshToken = New(codeInvalidRefreshToken, msgInvalidRefreshToken, nil, Extra{
		HttpStatus:    http.StatusForbidden,
		Translateable: true,
	})
	RequiredVerifyToken = New(codeRequiredVerifyToken, msgRequiredVerifyToken, nil, Extra{
		HttpStatus:    http.StatusForbidden,
		Translateable: true,
	})
	ExcludedVerifyToken = New(codeExcludedVerifyToken, msgExcludedVerifyToken, nil, Extra{
		HttpStatus:    http.StatusForbidden,
		Translateable: true,
	})
	Unauthorized = New(codeUnauthorized, msgUnauthorized, nil, Extra{
		HttpStatus:    http.StatusUnauthorized,
		Translateable: true,
	})
	PermissionDenied = New(codePermissionDenied, msgPermissionDenied, nil, Extra{
		HttpStatus:    http.StatusForbidden,
		Translateable: true,
	})
	RecaptchaFailed = New(codeRecaptchaFailed, msgRecaptchaFailed, nil, Extra{
		HttpStatus:    http.StatusForbidden,
		Translateable: true,
	})
	RecaptchaRequired = New(codeRecaptchaRequired, msgRecaptchaRequired, nil, Extra{
		HttpStatus:    http.StatusForbidden,
		Translateable: true,
	})
)

func ValidationFailed(data interface{}) *RC {
	return New(codeValidationFailed, msgValidationFailed, data, Extra{
		HttpStatus:    http.StatusUnprocessableEntity,
		Translateable: true,
	})
}

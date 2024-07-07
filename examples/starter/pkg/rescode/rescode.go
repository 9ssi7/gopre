package rescode

// RC is a struct that contains the code, message, http status, and translateable.
// Code is the code of the error.
type RC struct {

	// Code is the code of the error.
	Code uint64

	// Message is the message of the error.
	Message string

	// HttpStatus is the http status of the error.
	HttpStatus int

	// Translateable is the flag to determine whether the message is translatable.
	Translateable bool

	// Data is the data of the error.
	Data any
}

// Extra is a struct that contains the http status and translateable.
type Extra struct {
	// HttpStatus is the http status of the error.
	HttpStatus int

	// Translateable is the flag to determine whether the message is translatable.
	Translateable bool
}

// DefaultConfig is a default configuration for the RC.
// HttpStatus is 400 and Translateable is true.
// if you want to change the default configuration, you can change the value here.
// Example:
//
//	rescode.DefaultConfig = rescode.Extra{
//		HttpStatus:    500,
//		Translateable: false,
//	}
var DefaultConfig = Extra{
	HttpStatus:    400,
	Translateable: true,
}

// New is a function to create a new RC.
func New(code uint64, message string, data any, extra ...Extra) *RC {
	e := DefaultConfig
	if len(extra) > 0 {
		e = extra[0]
	}
	return &RC{
		Code:          code,
		Message:       message,
		Data:          data,
		HttpStatus:    e.HttpStatus,
		Translateable: e.Translateable,
	}
}

// JSON is a function to return the RC as a JSON.
func (r *RC) JSON(msgs ...string) map[string]interface{} {
	msg := r.Message
	if len(msgs) > 0 {
		msg = msgs[0]
	}
	json := map[string]interface{}{
		"code": r.Code,
	}
	if r.Message != "" {
		json["message"] = msg
	}
	if r.Data != nil {
		json["data"] = r.Data
	}
	return json
}

// Error is a function to return the message of the RC.
func (r *RC) Error() string {
	return r.Message
}

package rescode

// RC is a struct that contains the code, message, http status, and translateable.
// Code is the code of the error.
type RC struct {

	// Code is the code of the error.
	Code uint64

	// Message is the message of the error.
	Message string

	// StatusCode is the http/rpc status code of the error.
	StatusCode int

	// Data is the data of the error.
	Data any

	// Error is the error of the error. (if the error is not nil)
	errors []error
}

type RcCreator func(err ...error) *RC

// New is a function to create a new RC.
func New(code uint64, status int, message string, data ...any) RcCreator {
	var d any
	if len(data) > 0 {
		d = data[0]
	}
	return func(err ...error) *RC {
		return &RC{
			Code:       code,
			Message:    message,
			Data:       d,
			StatusCode: status,
			errors:     err,
		}
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

func (r *RC) SetData(data any) *RC {
	r.Data = data
	return r
}

// Error is a function to return the message of the RC.
func (r *RC) Error() string {
	if len(r.errors) > 0 {
		return r.errors[0].Error()
	}
	return r.Message
}

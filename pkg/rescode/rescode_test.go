package rescode

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		code    uint64
		status  int
		message string
		data    any
		err     error
		want    *RC
	}{
		{
			name:    "Basic",
			code:    100,
			status:  200,
			message: "Success",
			want:    &RC{Code: 100, Message: "Success", StatusCode: 200},
		},
		{
			name:    "WithData",
			code:    101,
			status:  201,
			message: "Created",
			data:    map[string]string{"foo": "bar"},
			want:    &RC{Code: 101, Message: "Created", StatusCode: 201, Data: map[string]string{"foo": "bar"}},
		},
		{
			name:    "WithError",
			code:    500,
			status:  500,
			message: "Internal Server Error",
			err:     errors.New("database error"),
			want:    &RC{Code: 500, Message: "Internal Server Error", StatusCode: 500, err: errors.New("database error")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			creator := New(tt.code, tt.status, tt.message, tt.data)
			rc := creator(tt.err)

			if rc.Code != tt.want.Code || rc.Message != tt.want.Message || rc.StatusCode != tt.want.StatusCode {
				t.Errorf("New() = %v, want %v", rc, tt.want)
			}

			if tt.data != nil {
				if rc.Data == nil {
					t.Errorf("New() data = %v, want %v", rc.Data, tt.want.Data)
				}
			}

			if tt.err != nil {
				if rc.OriginalError().Error() != tt.err.Error() {
					t.Errorf("New() error = %v, want %v", rc.Error(), tt.err.Error())
				}
			}
		})
	}
}

func TestRC_SetData(t *testing.T) {
	rc := &RC{}
	data := "some data"
	rc.SetData(data)

	if rc.Data != data {
		t.Errorf("SetData() = %v, want %v", rc.Data, data)
	}
}

func TestRC_JSON(t *testing.T) {
	tests := []struct {
		name string
		rc   *RC
		msgs []string
		want map[string]interface{}
	}{
		{
			name: "Basic",
			rc:   &RC{Code: 100, Message: "Success"},
			want: map[string]interface{}{"code": uint64(100), "message": "Success"},
		},
		{
			name: "Empty Message",
			rc:   &RC{Code: 100},
			want: map[string]interface{}{"code": uint64(100)},
		},
		{
			name: "WithData",
			rc:   &RC{Code: 100, Message: "Success", Data: "some data"},
			want: map[string]interface{}{"code": uint64(100), "message": "Success", "data": "some data"},
		},
		{
			name: "WithCustomMessage",
			rc:   &RC{Code: 100, Message: "Success"},
			msgs: []string{"Custom Message"},
			want: map[string]interface{}{"code": uint64(100), "message": "Custom Message"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.rc.JSON(tt.msgs...)
			if tt.want["code"] != got["code"] || tt.want["message"] != got["message"] || tt.want["data"] != got["data"] {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRC_Error(t *testing.T) {
	tests := []struct {
		name    string
		rc      *RC
		wantErr string
	}{
		{
			name:    "With Message",
			rc:      &RC{Message: "Error message", err: errors.New("Error message")},
			wantErr: "Error message",
		},
		{
			name:    "With Errors",
			rc:      &RC{Message: "Error message", err: errors.New("underlying error")},
			wantErr: "underlying error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.rc.OriginalError().Error(); err != tt.wantErr {
				t.Errorf("Error() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

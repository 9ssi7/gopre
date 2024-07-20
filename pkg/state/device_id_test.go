package state

import (
	"context"
	"testing"
)

func TestSetDeviceId(t *testing.T) {
	ctx := context.Background()
	deviceId := "test-device-id"

	ctx = SetDeviceId(ctx, deviceId)

	got := ctx.Value(deviceIdKey)
	if got == nil {
		t.Errorf("SetDeviceId() did not set deviceId in context")
	}

	if got != deviceId {
		t.Errorf("SetDeviceId() = %v, want %v", got, deviceId)
	}
}

func TestGetDeviceId(t *testing.T) {
	tests := []struct {
		name string
		ctx  context.Context
		want string
	}{
		{
			name: "DeviceIdExists",
			ctx:  context.WithValue(context.Background(), deviceIdKey, "device123"),
			want: "device123",
		},
		{
			name: "DeviceIdDoesNotExist",
			ctx:  context.Background(),
			want: "",
		},
		{
			name: "WrongTypeInContext",
			ctx:  context.WithValue(context.Background(), deviceIdKey, 123),
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetDeviceId(tt.ctx)
			if got != tt.want {
				t.Errorf("GetDeviceId() = %v, want %v", got, tt.want)
			}
		})
	}
}

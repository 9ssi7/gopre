package state

import "context"

type deviceIDKey string

var deviceIdKey deviceIDKey = "deviceId"

// SetDeviceId sets the device id in the context
func SetDeviceId(ctx context.Context, deviceId string) context.Context {
	return context.WithValue(ctx, deviceIdKey, deviceId)
}

// GetDeviceId gets the device id from the context
func GetDeviceId(ctx context.Context) string {
	if deviceId, ok := ctx.Value(deviceIdKey).(string); ok {
		return deviceId
	}
	return ""
}

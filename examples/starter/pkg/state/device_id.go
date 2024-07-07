package state

import "context"

var deviceIdKey = "deviceId"

// SetDeviceId sets the device id in the context
func SetDeviceId(ctx context.Context, deviceId string) {
	ctx = context.WithValue(ctx, deviceIdKey, deviceId)
}

// GetDeviceId gets the device id from the context
func GetDeviceId(ctx context.Context) string {
	if deviceId, ok := ctx.Value(deviceIdKey).(string); ok {
		return deviceId
	}
	return ""
}

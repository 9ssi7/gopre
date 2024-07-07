package middlewares

import (
	"context"
	"time"

	"github.com/9ssi7/gopre-starter/config"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func DeviceId(c *fiber.Ctx) error {
	deviceId := c.Cookies("device_id")
	_, err := uuid.Parse(deviceId)
	if deviceId == "" || err != nil {
		deviceId = uuid.New().String()
		c.Cookie(config.ApplyCookie(&fiber.Cookie{
			Name:    "device_id",
			Value:   deviceId,
			Expires: time.Now().Add(24 * time.Hour * 365),
			Domain:  config.ReadValue().HttpHeaders.Domain,
		}))
	}
	c.Locals("deviceId", deviceId)
	c.SetUserContext(context.WithValue(c.UserContext(), "deviceId", deviceId))
	return c.Next()
}

func DeviceIdParse(c *fiber.Ctx) string {
	return c.Locals("deviceId").(string)
}

package middlewares

import "github.com/gofiber/fiber/v2"

func IpAddr(c *fiber.Ctx) error {
	ip := c.Get("CF-Connecting-IP") // this for Cloudflare origin server proxy, keeper of the IP
	if ip == "" {
		ip = c.IP()
	}
	c.Locals("ip", ip)
	return c.Next()
}

func IpMustParse(c *fiber.Ctx) string {
	return c.Locals("ip").(string)
}

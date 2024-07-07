package middlewares

import (
	"github.com/9ssi7/gopre-starter/config"
	"github.com/9ssi7/gopre-starter/pkg/rescode"
	"github.com/9ssi7/turnstile"
	"github.com/gofiber/fiber/v2"
)

func NewTurnstile() fiber.Handler {
	srv := turnstile.New(turnstile.Config{
		Secret: config.ReadValue().Turnstile.Secret,
	})
	return func(ctx *fiber.Ctx) error {
		if config.ReadValue().Turnstile.Skip {
			return ctx.Next()
		}
		ip := IpMustParse(ctx)
		token := ctx.Get("X-Turnstile-Token")
		ok, err := srv.Verify(ctx.UserContext(), token, ip)
		if err != nil {
			return rescode.RecaptchaFailed
		}
		if !ok {
			return rescode.RecaptchaRequired
		}
		return ctx.Next()
	}
}

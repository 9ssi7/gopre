package middlewares

import (
	"time"

	"github.com/9ssi7/gopre-starter/config"
	"github.com/9ssi7/gopre-starter/internal/app"
	"github.com/9ssi7/gopre-starter/internal/app/queries"
	"github.com/9ssi7/gopre-starter/pkg/rescode"
	"github.com/9ssi7/gopre-starter/pkg/token"
	"github.com/gofiber/fiber/v2"
)

func RefreshRequired(c *fiber.Ctx) error {
	u := c.Locals("user_refresh")
	if u == nil || u.(*token.UserClaim).IsExpired() || !u.(*token.UserClaim).IsRefresh {
		return rescode.Unauthorized
	}
	return c.Next()
}

func RefreshExcluded(c *fiber.Ctx) error {
	u := c.Locals("user_refresh")
	if u == nil || u.(*token.UserClaim).IsExpired() || !u.(*token.UserClaim).IsRefresh {
		return c.Next()
	}
	return rescode.InvalidRefreshToken
}

func NewRefreshInitialize(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		t := refreshGetToken(c)
		ip := IpMustParse(c)
		res, err := app.Queries.AuthVerifyRefresh(c.UserContext(), queries.AuthVerifyRefresh{
			AccessToken:  AccessGetToken(c),
			RefreshToken: t,
			IpAddr:       ip,
		})
		if err != nil {
			return err
		}
		c.Locals("user_refresh", res.User)
		c.Locals("refresh_token", t)
		return c.Next()
	}
}
func RefreshMustParse(c *fiber.Ctx) *token.UserClaim {
	return c.Locals("user_refresh").(*token.UserClaim)
}
func RefreshParse(c *fiber.Ctx) *token.UserClaim {
	u := c.Locals("user_refresh")
	if u == nil {
		return nil
	}
	return u.(*token.UserClaim)
}

func RefreshParseToken(c *fiber.Ctx) string {
	return c.Locals("refresh_token").(string)
}

func RefreshTokenSetCookie(ctx *fiber.Ctx, t string) {
	ctx.Cookie(config.ApplyCookie(&fiber.Cookie{
		Name:    "refresh_token",
		Value:   t,
		Domain:  config.ReadValue().HttpHeaders.Domain,
		Expires: time.Now().Add(token.RefreshTokenDuration),
	}))
}

func RefreshTokenRemoveCookie(ctx *fiber.Ctx) {
	ctx.Cookie(config.ApplyCookie(&fiber.Cookie{
		Name:    "refresh_token",
		Value:   "",
		Domain:  config.ReadValue().HttpHeaders.Domain,
		Expires: time.Now().Add(-1 * time.Hour),
	}))
}

func refreshGetToken(ctx *fiber.Ctx) string {
	t := ctx.Cookies("refresh_token")
	if t == "" {
		t = refreshGetBearerToken(ctx)
	}
	if t == "" {
		t = ctx.Get("X-Refresh-Token")
	}
	return t
}

func refreshGetBearerToken(c *fiber.Ctx) string {
	b := c.Get("X-Refresh-Token")
	if b == "" {
		return ""
	}
	return b[7:]
}

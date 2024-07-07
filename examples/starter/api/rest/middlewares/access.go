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

func NewAccessRequired(isUnverified bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		u := c.Locals("user")
		if u == nil || (!isUnverified && u.(*token.UserClaim).IsExpired()) || !u.(*token.UserClaim).IsAccess {
			return rescode.Unauthorized
		}
		return c.Next()
	}
}

func AccessExcluded(c *fiber.Ctx) error {
	u := c.Locals("user")
	if u == nil || u.(*token.UserClaim).IsExpired() || !u.(*token.UserClaim).IsAccess {
		return c.Next()
	}
	return rescode.InvalidAccess
}

func NewAccessInitialize(app app.App, isUnverified bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		t := AccessGetToken(c)
		if t == "" {
			// if access required, use accessrequired middleware
			return c.Next()
		}
		ip := IpMustParse(c)
		res, err := app.Queries.AuthVerifyAccess(c.UserContext(), queries.AuthVerifyAccess{
			AccessToken:  t,
			IpAddr:       ip,
			IsUnverified: isUnverified,
		})
		if err != nil {
			return err
		}
		c.Locals("user", res.User)
		c.Locals("access_token", t)
		return c.Next()
	}
}

func AccessMustParse(c *fiber.Ctx) *token.UserClaim {
	return c.Locals("user").(*token.UserClaim)
}
func AccessParse(c *fiber.Ctx) *token.UserClaim {
	u := c.Locals("user")
	if u == nil {
		return nil
	}
	return u.(*token.UserClaim)
}

func AccessParseToken(c *fiber.Ctx) string {
	return c.Locals("access_token").(string)
}

func AccessTokenSetCookie(ctx *fiber.Ctx, t string) {
	ctx.Cookie(config.ApplyCookie(&fiber.Cookie{
		Name:    "access_token",
		Value:   t,
		Domain:  config.ReadValue().HttpHeaders.Domain,
		Expires: time.Now().Add(token.AccessTokenDuration),
	}))
}

func AccessTokenRemoveCookie(ctx *fiber.Ctx) {
	ctx.Cookie(config.ApplyCookie(&fiber.Cookie{
		Name:    "access_token",
		Value:   "",
		Domain:  config.ReadValue().HttpHeaders.Domain,
		Expires: time.Now().Add(-1 * time.Hour),
	}))
}

func AccessGetToken(ctx *fiber.Ctx) string {
	t := ctx.Cookies("access_token")
	if t == "" {
		t = accessGetBearerToken(ctx)
	}
	if t == "" {
		t = ctx.Get("Authorization")
	}
	return t
}

func accessGetBearerToken(c *fiber.Ctx) string {
	b := c.Get("Authorization")
	if b == "" {
		return ""
	}
	return b[7:]
}

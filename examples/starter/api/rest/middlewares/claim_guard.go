package middlewares

import (
	"github.com/9ssi7/gopre-starter/pkg/claguard"
	"github.com/9ssi7/gopre-starter/pkg/rescode"
	"github.com/gofiber/fiber/v2"
)

type ClaimGuardConfig struct {
	Claims []string
}

func NewClaimGuard(claims []string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		u := AccessMustParse(ctx)
		if claguard.Check(u.Roles, claims) {
			return ctx.Next()
		}
		return rescode.PermissionDenied
	}
}

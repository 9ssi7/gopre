package restsrv

import (
	"fmt"
	"strings"
	"time"

	"github.com/9ssi7/gopre-starter/api/rest/middlewares"
	"github.com/9ssi7/gopre-starter/config"
	"github.com/9ssi7/gopre-starter/internal/app"
	"github.com/9ssi7/gopre-starter/internal/domain/valobj"
	"github.com/9ssi7/gopre-starter/pkg/rescode"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/mileusna/useragent"
)

type Srv interface {
	ErrorHandler() fiber.ErrorHandler
	Turnstile() fiber.Handler
	RateLimit(limit int) fiber.Handler
	Cors() fiber.Handler
	IpAddr() fiber.Handler
	Timeout(fn fiber.Handler) fiber.Handler
	I18n() fiber.Handler

	DeviceId() fiber.Handler

	AccessInit(isUnverified ...bool) fiber.Handler
	AccessExcluded() fiber.Handler
	AccessRequired(isUnverified ...bool) fiber.Handler

	RefreshInit() fiber.Handler
	RefreshExcluded() fiber.Handler
	RefreshRequired() fiber.Handler

	VerifyTokenRequired() fiber.Handler
	VerifyTokenExcluded() fiber.Handler
	ClaimGuard(extra ...string) fiber.Handler

	MakeDevice(ctx *fiber.Ctx) *valobj.Device
}

type srv struct {
	app app.App
}

func New(app app.App) Srv {
	return srv{
		app: app,
	}
}

func (s srv) ErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusBadRequest
		fmt.Println(err.Error())
		if e, ok := err.(*rescode.RC); ok {
			msg := e.Message
			if e.Translateable {
				// msg = translator.Translate(e.Message, i18n.ParseLocale(c))
			}
			return c.Status(e.HttpStatus).JSON(e.JSON(msg))
		}
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}
		return c.Status(code).JSON(map[string]interface{}{})
	}
}

func (s srv) IpAddr() fiber.Handler {
	return middlewares.IpAddr
}

func (s srv) I18n() fiber.Handler {
	return middlewares.I18n
}

func (s srv) Turnstile() fiber.Handler {
	return middlewares.NewTurnstile()
}

func (h srv) RateLimit(limit int) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        limit,
		Expiration: 3 * time.Minute,
	})
}

func (h srv) Cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowMethods:     config.ReadValue().HttpHeaders.AllowedMethods,
		AllowHeaders:     config.ReadValue().HttpHeaders.AllowedHeaders,
		AllowCredentials: config.ReadValue().HttpHeaders.AllowCredentials,
		ExposeHeaders:    config.ReadValue().HttpHeaders.ExposeHeaders,
		AllowOriginsFunc: func(origin string) bool {
			origins := strings.Split(config.ReadValue().HttpHeaders.AllowedOrigins, ",")
			for _, o := range origins {
				if strings.Contains(origin, o) {
					return true
				}
			}
			return false
		},
	})
}

func (h srv) Timeout(fn fiber.Handler) fiber.Handler {
	return timeout.NewWithContext(fn, 50*time.Second)
}

func (h srv) DeviceId() fiber.Handler {
	return middlewares.DeviceId
}

func (h srv) AccessInit(isUnverified ...bool) fiber.Handler {
	verified := false
	if len(isUnverified) > 0 {
		verified = isUnverified[0]
	}
	return middlewares.NewAccessInitialize(h.app, verified)
}

func (h srv) AccessExcluded() fiber.Handler {
	return middlewares.AccessExcluded
}

func (h srv) AccessRequired(isUnverified ...bool) fiber.Handler {
	verified := false
	if len(isUnverified) > 0 {
		verified = isUnverified[0]
	}
	return middlewares.NewAccessRequired(verified)
}

func (h srv) RefreshInit() fiber.Handler {
	return middlewares.NewRefreshInitialize(h.app)
}

func (h srv) RefreshExcluded() fiber.Handler {
	return middlewares.RefreshExcluded
}

func (h srv) RefreshRequired() fiber.Handler {
	return middlewares.RefreshRequired
}

func (h srv) VerifyTokenRequired() fiber.Handler {
	return middlewares.VerifyRequired
}

func (h srv) VerifyTokenExcluded() fiber.Handler {
	return middlewares.VerifyExcluded
}

func (h srv) ClaimGuard(extra ...string) fiber.Handler {
	claims := []string{config.Roles.AdminSuper}
	if len(extra) > 0 {
		claims = append(claims, extra...)
	}
	return middlewares.NewClaimGuard(claims)
}

func (h srv) MakeDevice(ctx *fiber.Ctx) *valobj.Device {
	ua := useragent.Parse(ctx.Get("User-Agent"))
	t := "desktop"
	if ua.Mobile {
		t = "mobile"
	} else if ua.Tablet {
		t = "tablet"
	} else if ua.Bot {
		t = "bot"
	}
	ip := ctx.Get("CF-Connecting-IP")
	if ip == "" {
		ip = ctx.IP()
	}
	return &valobj.Device{
		Name: ua.Name,
		Type: t,
		OS:   ua.OS,
		IP:   ip,
	}
}

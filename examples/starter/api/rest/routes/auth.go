package routes

import (
	"github.com/9ssi7/gopre-starter/api/rest/middlewares"
	restsrv "github.com/9ssi7/gopre-starter/api/rest/srv"
	"github.com/9ssi7/gopre-starter/internal/app"
	"github.com/9ssi7/gopre-starter/internal/app/commands"
	"github.com/9ssi7/gopre-starter/internal/app/queries"
	"github.com/gofiber/fiber/v2"
)

func Auth(router fiber.Router, srv restsrv.Srv, app app.App) {
	router.Post("/auth/start", srv.VerifyTokenExcluded(), srv.Timeout(authStart(srv, app)))
	router.Post("/auth/login", srv.AccessInit(), srv.AccessExcluded(), srv.VerifyTokenRequired(), srv.Timeout(authLogin(srv, app)))
	router.Post("/auth/logout", srv.AccessInit(true), srv.AccessRequired(true), srv.Timeout(authLogout(app)))
	router.Put("/auth/refresh", srv.RefreshInit(), srv.RefreshRequired(), srv.Timeout(authRefresh(app)))
	router.Get("/auth/check", srv.AccessInit(), srv.AccessExcluded(), srv.Timeout(authCheck(app)))
	router.Get("/auth", srv.AccessInit(), srv.AccessRequired(), srv.Timeout(authCurrent(app)))
}

func authLogin(srv restsrv.Srv, app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.AuthLogin
		if err := c.BodyParser(&cmd); err != nil {
			return err
		}
		cmd.Device = srv.MakeDevice(c)
		cmd.VerifyToken = middlewares.VerifyTokenParse(c)
		res, err := app.Commands.AuthLogin(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		middlewares.VerifyTokenRemove(c)
		middlewares.AccessTokenSetCookie(c, res.AccessToken)
		middlewares.RefreshTokenSetCookie(c, res.RefreshToken)
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func authLogout(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cmd := commands.AuthLogout{
			UserId: middlewares.AccessMustParse(c).Id,
		}
		res, err := app.Commands.AuthLogout(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		middlewares.AccessTokenRemoveCookie(c)
		middlewares.RefreshTokenRemoveCookie(c)
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func authRefresh(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cmd := commands.AuthRefresh{
			UserId:       middlewares.RefreshMustParse(c).Id,
			AccessToken:  middlewares.AccessGetToken(c),
			RefreshToken: middlewares.RefreshParseToken(c),
			IpAddress:    middlewares.IpMustParse(c),
		}
		res, err := app.Commands.AuthRefresh(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		middlewares.AccessTokenSetCookie(c, res.AccessToken)
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func authStart(srv restsrv.Srv, app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.AuthStart
		if err := c.BodyParser(&cmd); err != nil {
			return err
		}
		cmd.Device = srv.MakeDevice(c)
		res, err := app.Commands.AuthStart(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		middlewares.VerifyTokenSet(c, res.VerifyToken)
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func authCheck(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		query := queries.AuthCheck{
			VerifyToken: middlewares.VerifyTokenParse(c),
		}
		res, err := app.Queries.AuthCheck(c.UserContext(), query)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func authCurrent(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(middlewares.AccessMustParse(c).User)
	}
}

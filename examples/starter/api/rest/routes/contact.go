package routes

import (
	restsrv "github.com/9ssi7/gopre-starter/api/rest/srv"
	"github.com/9ssi7/gopre-starter/config"
	"github.com/9ssi7/gopre-starter/internal/app"
	"github.com/9ssi7/gopre-starter/internal/app/commands"
	"github.com/9ssi7/gopre-starter/internal/app/queries"
	"github.com/9ssi7/gopre-starter/pkg/list"
	"github.com/gofiber/fiber/v2"
)

func Contact(router fiber.Router, srv restsrv.Srv, app app.App) {
	router.Post("/contact", srv.Turnstile(), srv.Timeout(contactCreate(app)))
	router.Get("/contact", srv.AccessInit(), srv.AccessRequired(), srv.ClaimGuard(config.Roles.Contact.Super, config.Roles.Contact.List), srv.Timeout(contactList(app)))
	router.Patch("/contact/:id/seen", srv.AccessInit(), srv.AccessRequired(), srv.ClaimGuard(config.Roles.Contact.Super, config.Roles.Contact.MarkSeen), srv.Timeout(contactMarkSeen(app)))
}

func contactCreate(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.ContactCreate
		if err := c.BodyParser(&cmd); err != nil {
			return err
		}
		res, err := app.Commands.ContactCreate(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusCreated).JSON(res)
	}
}

func contactList(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var pagi list.PagiRequest
		if err := c.QueryParser(&pagi); err != nil {
			return err
		}
		res, err := app.Queries.ContactList(c.UserContext(), queries.ContactList{PagiRequest: pagi})
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func contactMarkSeen(app app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cmd commands.ContactMarkSeen
		if err := c.ParamsParser(&cmd); err != nil {
			return err
		}
		res, err := app.Commands.ContactMarkSeen(c.UserContext(), cmd)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

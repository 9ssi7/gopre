package middlewares

import (
	"context"
	"strings"

	"github.com/9ssi7/gopre-starter/config"
	"github.com/gofiber/fiber/v2"
)

var localeKey = "lang"

func ParseLocale(ctx *fiber.Ctx) string {
	return ctx.Locals(localeKey).(string)
}

func I18n(c *fiber.Ctx) error {
	acceptedLanguages := config.ReadValue().I18n.Locales
	l := c.Query("lang")
	list := strings.Split(l, ";")
	alternative := ""
	locales := findLocales(list)
	for _, v := range acceptedLanguages {
		if locales[v] {
			l = v
			break
		}
	}
	if len(list) > 1 {
		alternative = list[1]
	}
	if alternative != "" && l != "" && locales[alternative] {
		l = alternative
	}
	c.Locals(localeKey, l)
	c.SetUserContext(context.WithValue(c.UserContext(), localeKey, l))
	return c.Next()
}

func findLocales(list []string) map[string]bool {
	locales := make(map[string]bool)
	acceptedLanguages := config.ReadValue().I18n.Locales
	for _, li := range list {
		lineItems := strings.Split(li, ",")
		for _, word := range lineItems {
			for _, v := range acceptedLanguages {
				if strings.Contains(word, v) {
					locales[v] = true
				}
			}
			if len(word) == 2 && word[1] == '-' {
				locales[strings.ToLower(word)] = true
			}
			if len(word) == 5 && word[2] == '-' {
				double := strings.Split(word, "-")
				locales[double[0]] = true
			}
		}
	}
	return locales
}

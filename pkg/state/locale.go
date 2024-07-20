package state

import "context"

type localeKeyType string

var localeKey localeKeyType = "locale"

// SetLocale sets the locale in the context
func SetLocale(ctx context.Context, locale string) context.Context {
	return context.WithValue(ctx, localeKey, locale)
}

// GetLocale gets the locale from the context
func GetLocale(ctx context.Context) string {
	if locale, ok := ctx.Value(localeKey).(string); ok {
		return locale
	}
	return "tr"
}

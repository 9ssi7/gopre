package state

import "context"

var localeKey = "locale"

// SetLocale sets the locale in the context
func SetLocale(ctx context.Context, locale string) {
	ctx = context.WithValue(ctx, localeKey, locale)
}

// GetLocale gets the locale from the context
func GetLocale(ctx context.Context) string {
	if locale, ok := ctx.Value(localeKey).(string); ok {
		return locale
	}
	return "tr"
}

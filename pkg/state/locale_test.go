package state

import (
	"context"
	"testing"
)

func TestSetLocale(t *testing.T) {
	ctx := context.Background()
	locale := "en-US"

	newCtx := SetLocale(ctx, locale)

	got := newCtx.Value(localeKey)
	if got == nil {
		t.Error("SetLocale() did not set locale in context")
	}

	if got != locale {
		t.Errorf("SetLocale() = %v, want %v", got, locale)
	}
}

func TestGetLocale(t *testing.T) {
	tests := []struct {
		name string
		ctx  context.Context
		want string
	}{
		{
			name: "LocaleExists",
			ctx:  SetLocale(context.Background(), "fr-FR"),
			want: "fr-FR",
		},
		{
			name: "LocaleDoesNotExist",
			ctx:  context.Background(),
			want: "tr", // Default locale
		},
		{
			name: "WrongTypeInContext",
			ctx:  context.WithValue(context.Background(), localeKey, 123),
			want: "tr", // Default locale
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetLocale(tt.ctx)
			if got != tt.want {
				t.Errorf("GetLocale() = %v, want %v", got, tt.want)
			}
		})
	}
}

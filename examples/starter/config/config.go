package config

import (
	"github.com/gofiber/fiber/v2"
	"go.deanishe.net/env"
)

type Smtp struct {
	Host     string `env:"SMTP_HOST"`
	Port     int    `env:"SMTP_PORT"`
	From     string `env:"SMTP_FROM"`
	Sender   string `env:"SMTP_SENDER"`
	Reply    string `env:"SMTP_REPLY"`
	Password string `env:"SMTP_PASSWORD"`
}

type Database struct {
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
	SslMode  string `env:"DB_SSL_MODE"`

	Migrate bool `env:"DB_MIGRATE" envDefault:"false"`
	Seed    bool `env:"DB_SEED" envDefault:"false"`
}

type I18n struct {
	Fallback string   `env:"I18N_FALLBACK_LANGUAGE" envDefault:"en"`
	Dir      string   `env:"I18N_DIR" envDefault:"./src/locales"`
	Locales  []string `env:"I18N_LOCALES" envDefault:"en,tr"`
}
type Http struct {
	Host  string `env:"HTTP_HOST" envDefault:"localhost"`
	Port  int    `env:"HTTP_PORT" envDefault:"3000"`
	Debug bool   `env:"HTTP_DEBUG" envDefault:"true"`
}

type KeyValueDb struct {
	Host string `env:"KV_DB_HOST"`
	Port string `env:"KV_DB_PORT"`
	Pw   string `env:"KV_DB_PASSWORD"`
	Db   int    `env:"KV_DB_NAME"`
}

type HttpHeaders struct {
	AllowedOrigins   string `env:"CORS_ALLOWED_ORIGINS" envDefault:"*"`
	AllowedMethods   string `env:"CORS_ALLOWED_METHODS" envDefault:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders   string `env:"CORS_ALLOWED_HEADERS" envDefault:"*"`
	ExposeHeaders    string `env:"CORS_EXPOSE_HEADERS" envDefault:"*"`
	AllowCredentials bool   `env:"CORS_ALLOW_CREDENTIALS" envDefault:"true"`
	Domain           string `env:"HTTP_HEADER_DOMAIN" envDefault:"*"`
}

type TokenSrv struct {
	Expiration int    `env:"TOKEN_EXPIRATION" envDefault:"3600"`
	Project    string `env:"TOKEN_PROJECT" envDefault:"empty"`
}

type RSA struct {
	PrivateKeyFile string `env:"RSA_PRIVATE_KEY"`
	PublicKeyFile  string `env:"RSA_PUBLIC_KEY"`
}

type Turnstile struct {
	Secret string `env:"CF_TURNSTILE_SECRET_KEY"`
	Skip   bool   `env:"CF_TURNSTILE_SKIP" envDefault:"false"`
}

var IsDevelopment bool = false

type App struct {
	IsDevelopment bool `env:"IS_DEVELOPMENT" envDefault:"false"`
	Adapters      struct {
		Smtp Smtp
		// Sms  Sms
		// Push Push
	}
	Turnstile   Turnstile
	Database    Database
	I18n        I18n
	Http        Http
	KeyValueDb  KeyValueDb
	HttpHeaders HttpHeaders
	TokenSrv    TokenSrv
	RSA         RSA
}

func ApplyCookie(cookie *fiber.Cookie) *fiber.Cookie {
	cookie.SameSite = "Strict"
	if IsDevelopment {
		cookie.SameSite = "Lax"
	} else {
		cookie.HTTPOnly = true
		cookie.Secure = true
	}
	return cookie
}

var configs *App

func ReadValue() *App {
	if configs != nil {
		return configs
	}
	if configs == nil {
		configs = &App{}
	}
	err := env.Bind(configs)
	if err != nil {
		panic(err)
	}
	IsDevelopment = configs.IsDevelopment
	return configs
}

package app

import (
	"github.com/9ssi7/gopre-starter/internal/app/commands"
	"github.com/9ssi7/gopre-starter/internal/app/queries"
	"github.com/9ssi7/gopre-starter/internal/app/services"
)

type App struct {
	Commands commands.Handlers
	Queries  queries.Handlers
	Services services.Handlers
}

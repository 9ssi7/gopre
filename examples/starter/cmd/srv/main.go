package main

import (
	"github.com/9ssi7/gopre-starter/api/rest"
	"github.com/9ssi7/gopre-starter/config"
	"github.com/9ssi7/gopre-starter/internal/app"
	"github.com/9ssi7/gopre-starter/internal/app/commands"
	"github.com/9ssi7/gopre-starter/internal/app/queries"
	"github.com/9ssi7/gopre-starter/internal/app/services"
	"github.com/9ssi7/gopre-starter/internal/domain/abstracts"
	"github.com/9ssi7/gopre-starter/internal/infra/db"
	"github.com/9ssi7/gopre-starter/internal/infra/db/migrations"
	"github.com/9ssi7/gopre-starter/internal/infra/db/seeds"
	"github.com/9ssi7/gopre-starter/internal/infra/keyval"
	"github.com/9ssi7/gopre-starter/internal/infra/repos"
	"github.com/9ssi7/gopre-starter/pkg/token"
	"github.com/9ssi7/gopre-starter/pkg/validation"
)

func main() {
	cnf := config.ReadValue()
	token.Init()
	db := db.ConnectDB()
	kvdb := keyval.ConnectDB()

	if cnf.Database.Migrate {
		migrations.Run(db)
	}
	if cnf.Database.Seed {
		seeds.Run(db)
	}

	r := abstracts.Repositories{
		ContactRepo: repos.NewContactRepo(db),
		UserRepo:    repos.NewUserRepo(db),
		SessionRepo: repos.NewSessionRepo(kvdb),
		VerifyRepo:  repos.NewVerifyRepo(kvdb),
	}

	v := validation.New()

	rest.New(app.App{
		Commands: commands.NewHandler(r, v),
		Queries:  queries.NewHandler(r),
		Services: services.NewHandler(),
	}).Listen()
}

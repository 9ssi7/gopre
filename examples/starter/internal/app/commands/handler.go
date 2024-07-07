package commands

import (
	"github.com/9ssi7/gopre-starter/internal/domain/abstracts"
	"github.com/9ssi7/gopre-starter/pkg/validation"
)

type Handlers struct {
	ContactCreate   ContactCreateHandler
	ContactMarkSeen ContactMarkSeenHandler
	AuthLogin       AuthLoginHandler
	AuthStart       AuthStartHandler
	AuthRefresh     AuthRefreshHandler
	AuthLogout      AuthLogoutHandler
}

func NewHandler(r abstracts.Repositories, v validation.Service) Handlers {
	return Handlers{
		ContactCreate:   NewContactCreateHandler(v, r.ContactRepo),
		ContactMarkSeen: NewContactMarkSeenHandler(v, r.ContactRepo),
		AuthLogin:       NewAuthLoginHandler(v, r.UserRepo, r.VerifyRepo, r.SessionRepo),
		AuthStart:       NewAuthStartHandler(v, r.VerifyRepo, r.UserRepo),
		AuthLogout:      NewAuthLogoutHandler(r.SessionRepo),
		AuthRefresh:     NewAuthRefreshHandler(r.SessionRepo, r.UserRepo),
	}
}

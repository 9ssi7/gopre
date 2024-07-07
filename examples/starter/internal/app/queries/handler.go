package queries

import "github.com/9ssi7/gopre-starter/internal/domain/abstracts"

type Handlers struct {
	ContactList       ContactListHandler
	AuthCheck         AuthCheckHandler
	AuthVerifyAccess  AuthVerifyAccessHandler
	AuthVerifyRefresh AuthVerifyRefreshHandler
}

func NewHandler(r abstracts.Repositories) Handlers {
	return Handlers{
		ContactList:       NewContactListHandler(r.ContactRepo),
		AuthCheck:         NewAuthCheckHandler(r.VerifyRepo),
		AuthVerifyAccess:  NewAuthVerifyAccessHandler(r.SessionRepo),
		AuthVerifyRefresh: NewAuthVerifyRefreshHandler(r.SessionRepo),
	}
}

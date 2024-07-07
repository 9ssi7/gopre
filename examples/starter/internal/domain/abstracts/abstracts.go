package abstracts

import (
	"context"

	"github.com/9ssi7/gopre-starter/internal/domain/aggregates"
	"github.com/9ssi7/gopre-starter/internal/domain/entities"
	"github.com/9ssi7/gopre-starter/pkg/list"
	"github.com/google/uuid"
)

type ContactRepo interface {
	Save(ctx context.Context, contact *entities.Contact) error
	FindById(ctx context.Context, id uuid.UUID) (*entities.Contact, error)
	FindAll(ctx context.Context, req *list.PagiRequest) (*list.PagiResponse[*entities.Contact], error)
}

type UserRepo interface {
	Save(ctx context.Context, user *entities.User) error
	FindByToken(ctx context.Context, token string) (*entities.User, error)
	FindById(ctx context.Context, id uuid.UUID) (*entities.User, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	FindByPhone(ctx context.Context, phone string) (*entities.User, error)
	Filter(ctx context.Context, req *list.PagiRequest, search string, isActive string) (*list.PagiResponse[*entities.User], error)
	FilterByRoles(ctx context.Context, req *list.PagiRequest, roles []string) (*list.PagiResponse[*entities.User], error)
	ListByIds(ctx context.Context, userIds []uuid.UUID) ([]*entities.User, error)
}

type SessionRepo interface {
	Save(ctx context.Context, userId uuid.UUID, session *aggregates.Session) error
	FindByIds(ctx context.Context, userId uuid.UUID, deviceId string) (*aggregates.Session, error)
	FindAllByUserId(ctx context.Context, userId uuid.UUID) ([]*aggregates.Session, error)
	Destroy(ctx context.Context, userId uuid.UUID, deviceId string) error
}

type VerifyRepo interface {
	Save(ctx context.Context, token string, verify *aggregates.Verify) error
	IsExists(ctx context.Context, token string, deviceId string) (bool, error)
	Find(ctx context.Context, token string, deviceId string) (*aggregates.Verify, error)
	Delete(ctx context.Context, token string, deviceId string) error
}

type Repositories struct {
	ContactRepo ContactRepo
	VerifyRepo  VerifyRepo
	SessionRepo SessionRepo
	UserRepo    UserRepo
}

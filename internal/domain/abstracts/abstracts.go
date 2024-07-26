package abstracts

import (
	"context"

	"github.com/9ssi7/gopre/internal/domain/aggregates"
	"github.com/9ssi7/gopre/internal/domain/entities"
	"github.com/9ssi7/gopre/pkg/list"
	"github.com/9ssi7/txn"
	"github.com/google/uuid"
)

type TxnAdapterRepo interface {
	GetTxnAdapter() txn.Adapter
}

type UserRepo interface {
	TxnAdapterRepo

	Save(ctx context.Context, user *entities.User) error
	IsExistsByEmail(ctx context.Context, email string) (bool, error)
	FindByToken(ctx context.Context, token string) (*entities.User, error)
	FindById(ctx context.Context, id uuid.UUID) (*entities.User, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	FindByPhone(ctx context.Context, phone string) (*entities.User, error)
	Filter(ctx context.Context, req *list.PagiRequest, search string, isActive string) (*list.PagiResponse[*entities.User], error)
}

type RoleRepo interface {
	TxnAdapterRepo

	Save(ctx context.Context, role *entities.Role) error
	FindById(ctx context.Context, id uuid.UUID) (*entities.Role, error)
	Filter(ctx context.Context, req *list.PagiRequest, search string, isActive string) (*list.PagiResponse[*entities.Role], error)
}

type SessionRepo interface {
	Save(ctx context.Context, userId uuid.UUID, session *aggregates.Session) error
	FindByIds(ctx context.Context, userId uuid.UUID, deviceId string) (*aggregates.Session, bool, error)
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
	VerifyRepo  VerifyRepo
	SessionRepo SessionRepo
	UserRepo    UserRepo
}

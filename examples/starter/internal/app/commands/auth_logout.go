package commands

import (
	"context"

	"github.com/9ssi7/gopre-starter/internal/domain/abstracts"
	"github.com/9ssi7/gopre-starter/pkg/cqrs"
	"github.com/9ssi7/gopre-starter/pkg/state"
	"github.com/google/uuid"
)

type AuthLogout struct {
	UserId uuid.UUID `json:"-"`
}

type AuthLogoutHandler cqrs.HandlerFunc[AuthLogout, *cqrs.Empty]

func NewAuthLogoutHandler(sessionRepo abstracts.SessionRepo) AuthLogoutHandler {
	return func(ctx context.Context, cmd AuthLogout) (*cqrs.Empty, error) {
		err := sessionRepo.Destroy(ctx, cmd.UserId, state.GetDeviceId(ctx))
		if err != nil {
			return nil, err
		}
		return &cqrs.Empty{}, nil
	}
}

package queries

import (
	"context"

	"github.com/9ssi7/gopre-starter/internal/domain/abstracts"
	"github.com/9ssi7/gopre-starter/pkg/cqrs"
	"github.com/9ssi7/gopre-starter/pkg/rescode"
	"github.com/9ssi7/gopre-starter/pkg/state"
)

type AuthCheck struct {
	VerifyToken string `json:"-"`
}

type AuthCheckHandler cqrs.HandlerFunc[AuthCheck, *cqrs.Empty]

func NewAuthCheckHandler(verifyRepo abstracts.VerifyRepo) AuthCheckHandler {
	return func(ctx context.Context, query AuthCheck) (*cqrs.Empty, error) {
		exists, err := verifyRepo.IsExists(ctx, query.VerifyToken, state.GetDeviceId(ctx))
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, rescode.NotFound
		}
		return &cqrs.Empty{}, nil
	}
}

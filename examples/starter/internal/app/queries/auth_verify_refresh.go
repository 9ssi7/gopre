package queries

import (
	"context"

	"github.com/9ssi7/gopre-starter/internal/domain/abstracts"
	"github.com/9ssi7/gopre-starter/pkg/cqrs"
	"github.com/9ssi7/gopre-starter/pkg/rescode"
	"github.com/9ssi7/gopre-starter/pkg/state"
	"github.com/9ssi7/gopre-starter/pkg/token"
)

type AuthVerifyRefresh struct {
	AccessToken  string
	RefreshToken string
	IpAddr       string
}

type AuthVerifyRefreshRes struct {
	User *token.UserClaim
}

type AuthVerifyRefreshHandler cqrs.HandlerFunc[AuthVerifyRefresh, *AuthVerifyRefreshRes]

func NewAuthVerifyRefreshHandler(sessionRepo abstracts.SessionRepo) AuthVerifyRefreshHandler {
	return func(ctx context.Context, query AuthVerifyRefresh) (*AuthVerifyRefreshRes, error) {
		claims, err := token.Client().Parse(query.RefreshToken)
		if err != nil {
			return nil, rescode.Failed
		}
		isValid, err := token.Client().Verify(query.RefreshToken)
		if err != nil {
			return nil, rescode.Failed
		}
		if !isValid {
			return nil, rescode.InvalidOrExpiredToken
		}
		session, err := sessionRepo.FindByIds(ctx, claims.Id, state.GetDeviceId(ctx))
		if err != nil {
			return nil, err
		}
		if !session.IsRefreshValid(query.AccessToken, query.RefreshToken, query.IpAddr) {
			return nil, rescode.InvalidRefreshToken
		}
		return &AuthVerifyRefreshRes{
			User: claims,
		}, nil
	}
}

package commands

import (
	"context"

	"github.com/9ssi7/gopre-starter/internal/domain/abstracts"
	"github.com/9ssi7/gopre-starter/pkg/cqrs"
	"github.com/9ssi7/gopre-starter/pkg/rescode"
	"github.com/9ssi7/gopre-starter/pkg/state"
	"github.com/9ssi7/gopre-starter/pkg/token"
	"github.com/google/uuid"
)

type AuthRefresh struct {
	AccessToken  string
	RefreshToken string
	IpAddress    string
	UserId       uuid.UUID
}

type AuthRefreshRes struct {
	AccessToken string
}

type AuthRefreshHandler cqrs.HandlerFunc[AuthRefresh, *AuthRefreshRes]

func NewAuthRefreshHandler(sessionRepo abstracts.SessionRepo, userRepo abstracts.UserRepo) AuthRefreshHandler {
	return func(ctx context.Context, cmd AuthRefresh) (*AuthRefreshRes, error) {
		session, err := sessionRepo.FindByIds(ctx, cmd.UserId, state.GetDeviceId(ctx))
		if err != nil {
			return nil, err
		}
		if !session.IsRefreshValid(cmd.AccessToken, cmd.RefreshToken, cmd.IpAddress) {
			return nil, rescode.InvalidRefreshOrAccessTokens
		}
		user, err := userRepo.FindById(ctx, cmd.UserId)
		if err != nil {
			return nil, err
		}
		accessToken, err := token.Client().GenerateAccessToken(token.User{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
			Roles: user.Roles,
		})
		if err != nil {
			return nil, err
		}
		session.Refresh(accessToken)
		if err := sessionRepo.Save(ctx, user.Id, session); err != nil {
			return nil, err
		}
		return &AuthRefreshRes{
			AccessToken: accessToken,
		}, nil
	}
}

package commands

import (
	"context"

	"github.com/9ssi7/gopre-starter/internal/domain/abstracts"
	"github.com/9ssi7/gopre-starter/internal/domain/aggregates"
	"github.com/9ssi7/gopre-starter/internal/domain/valobj"
	"github.com/9ssi7/gopre-starter/pkg/cqrs"
	"github.com/9ssi7/gopre-starter/pkg/rescode"
	"github.com/9ssi7/gopre-starter/pkg/state"
	"github.com/9ssi7/gopre-starter/pkg/token"
	"github.com/9ssi7/gopre-starter/pkg/validation"
)

type AuthLogin struct {
	VerifyToken string         `json:"-"`
	Code        string         `json:"code" validate:"required,numeric,len=4"`
	Device      *valobj.Device `json:"-"`
}

type AuthLoginRes struct {
	AccessToken  string `json:"-"`
	RefreshToken string `json:"-"`
}

type AuthLoginHandler cqrs.HandlerFunc[AuthLogin, *AuthLoginRes]

func NewAuthLoginHandler(v validation.Service, userRepo abstracts.UserRepo, verifyRepo abstracts.VerifyRepo, sessionRepo abstracts.SessionRepo) AuthLoginHandler {
	return func(ctx context.Context, cmd AuthLogin) (*AuthLoginRes, error) {
		err := v.ValidateStruct(ctx, cmd)
		if err != nil {
			return nil, err
		}
		verify, err := verifyRepo.Find(ctx, cmd.VerifyToken, state.GetDeviceId(ctx))
		if err != nil {
			return nil, err
		}
		if verify.IsExpired() {
			return nil, rescode.VerificationExpired
		}
		if verify.IsExceeded() {
			return nil, rescode.VerificationExceeded
		}
		if cmd.Code != verify.Code {
			verify.IncTryCount()
			err = verifyRepo.Save(ctx, cmd.VerifyToken, verify)
			if err != nil {
				return nil, err
			}
			return nil, rescode.VerificationInvalid
		}
		err = verifyRepo.Delete(ctx, cmd.VerifyToken, state.GetDeviceId(ctx))
		if err != nil {
			return nil, err
		}
		user, err := userRepo.FindById(ctx, verify.UserId)
		if err != nil {
			return nil, err
		}
		accessToken, refreshToken, err := token.Client().Generate(token.User{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
			Roles: user.Roles,
		})
		if err != nil {
			return nil, rescode.Failed
		}
		ses := aggregates.NewSession(*cmd.Device, state.GetDeviceId(ctx), accessToken, refreshToken)
		if err = sessionRepo.Save(ctx, user.Id, ses); err != nil {
			return nil, err
		}
		return &AuthLoginRes{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, nil
	}
}

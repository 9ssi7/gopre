package commands

import (
	"context"

	"github.com/9ssi7/gopre-starter/assets"
	"github.com/9ssi7/gopre-starter/internal/app/messages"
	"github.com/9ssi7/gopre-starter/internal/domain/abstracts"
	"github.com/9ssi7/gopre-starter/internal/domain/aggregates"
	"github.com/9ssi7/gopre-starter/internal/domain/entities"
	"github.com/9ssi7/gopre-starter/internal/domain/valobj"
	"github.com/9ssi7/gopre-starter/internal/infra/mail"
	"github.com/9ssi7/gopre-starter/pkg/cqrs"
	"github.com/9ssi7/gopre-starter/pkg/rescode"
	"github.com/9ssi7/gopre-starter/pkg/state"
	"github.com/9ssi7/gopre-starter/pkg/validation"
	"github.com/google/uuid"
)

type AuthStart struct {
	Phone  string         `json:"phone" validate:"required_without=Email,omitempty,phone"`
	Email  string         `json:"email" validate:"required_without=Phone,omitempty,email"`
	Device *valobj.Device `json:"-"`
}

type AuthStartRes struct {
	VerifyToken string `json:"-"`
}

type AuthStartHandler cqrs.HandlerFunc[AuthStart, *AuthStartRes]

func NewAuthStartHandler(v validation.Service, verifyRepo abstracts.VerifyRepo, userRepo abstracts.UserRepo) AuthStartHandler {
	return func(ctx context.Context, cmd AuthStart) (*AuthStartRes, error) {
		err := v.ValidateStruct(ctx, cmd)
		if err != nil {
			return nil, err
		}
		var user *entities.User
		if cmd.Phone != "" {
			user, err = userRepo.FindByPhone(ctx, cmd.Phone)
			if err != nil {
				return nil, err
			}
		} else {
			user, err = userRepo.FindByEmail(ctx, cmd.Email)
			if err != nil {
				return nil, err
			}
		}
		if user == nil {
			return nil, rescode.NotFound
		}
		if !user.IsActive {
			return nil, rescode.UserDisabled
		}
		if user.TempToken != nil && *user.TempToken != "" {
			return nil, rescode.UserVerifyRequired
		}
		verifyToken := uuid.New().String()
		verify := aggregates.NewVerify(user.Id, state.GetDeviceId(ctx), state.GetLocale(ctx))
		err = verifyRepo.Save(ctx, verifyToken, verify)
		if err != nil {
			return nil, err
		}
		go func() {
			mail.GetClient().SendWithTemplate(mail.SendWithTemplateConfig{
				SendConfig: mail.SendConfig{
					To:      []string{user.Email},
					Subject: messages.AuthVerifySubject,
					Message: verify.Code,
				},
				Template: assets.Templates.AuthVerify,
				Data: map[string]interface{}{
					"Code":    verify.Code,
					"IP":      mail.GetField(cmd.Device.IP),
					"Browser": mail.GetField(cmd.Device.Name),
					"OS":      mail.GetField(cmd.Device.OS),
				},
			})
		}()
		return &AuthStartRes{
			VerifyToken: verifyToken,
		}, nil
	}
}

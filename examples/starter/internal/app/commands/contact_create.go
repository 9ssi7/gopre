package commands

import (
	"context"

	"github.com/9ssi7/gopre-starter/internal/domain/abstracts"
	"github.com/9ssi7/gopre-starter/internal/domain/entities"
	"github.com/9ssi7/gopre-starter/pkg/cqrs"
	"github.com/9ssi7/gopre-starter/pkg/validation"
)

type ContactCreate struct {
	Message string `json:"message" validate:"required,min=3,max=255"`
	Email   string `json:"email" validate:"required,email"`
}

type ContactCreateHandler cqrs.HandlerFunc[ContactCreate, *cqrs.Empty]

func NewContactCreateHandler(v validation.Service, repo abstracts.ContactRepo) ContactCreateHandler {
	return func(ctx context.Context, cmd ContactCreate) (*cqrs.Empty, error) {
		err := v.ValidateStruct(ctx, cmd)
		if err != nil {
			return nil, err
		}
		err = repo.Save(ctx, entities.NewContact(cmd.Message, cmd.Email))
		if err != nil {
			return nil, err
		}
		return &cqrs.Empty{}, nil
	}
}

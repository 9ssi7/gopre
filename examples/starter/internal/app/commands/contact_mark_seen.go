package commands

import (
	"context"

	"github.com/9ssi7/gopre-starter/internal/domain/abstracts"
	"github.com/9ssi7/gopre-starter/pkg/cqrs"
	"github.com/9ssi7/gopre-starter/pkg/validation"
	"github.com/google/uuid"
)

type ContactMarkSeen struct {
	Id uuid.UUID `json:"id" validate:"required,uuid"`
}

type ContactMarkSeenHandler cqrs.HandlerFunc[ContactMarkSeen, *cqrs.Empty]

func NewContactMarkSeenHandler(v validation.Service, repo abstracts.ContactRepo) ContactMarkSeenHandler {
	return func(ctx context.Context, cmd ContactMarkSeen) (*cqrs.Empty, error) {
		err := v.ValidateStruct(ctx, cmd)
		if err != nil {
			return nil, err
		}
		contact, err := repo.FindById(ctx, cmd.Id)
		if err != nil {
			return nil, err
		}
		contact.MarkSeen()
		err = repo.Save(ctx, contact)
		if err != nil {
			return nil, err
		}
		return &cqrs.Empty{}, nil
	}
}

package queries

import (
	"context"

	"github.com/9ssi7/gopre-starter/internal/app/dtos"
	"github.com/9ssi7/gopre-starter/internal/domain/abstracts"
	"github.com/9ssi7/gopre-starter/pkg/cqrs"
	"github.com/9ssi7/gopre-starter/pkg/list"
)

type ContactList struct {
	list.PagiRequest
}

type ContactListHandler cqrs.HandlerFunc[ContactList, *list.PagiResponse[*dtos.ContactListDto]]

func NewContactListHandler(repo abstracts.ContactRepo) ContactListHandler {
	return func(ctx context.Context, cmd ContactList) (*list.PagiResponse[*dtos.ContactListDto], error) {
		cmd.Default()
		list, err := repo.FindAll(ctx, &cmd.PagiRequest)
		if err != nil {
			return nil, err
		}
		return dtos.MapContactList(list), nil
	}
}

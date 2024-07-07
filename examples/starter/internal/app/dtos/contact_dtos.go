package dtos

import (
	"time"

	"github.com/9ssi7/gopre-starter/internal/domain/entities"
	"github.com/9ssi7/gopre-starter/pkg/list"
	"github.com/google/uuid"
)

type ContactListDto struct {
	Id        uuid.UUID `json:"id"`
	Message   string    `json:"message"`
	Email     string    `json:"email"`
	IsSeen    bool      `json:"is_seen"`
	CreatedAt time.Time `json:"created_at"`
}

func MapContactList(res *list.PagiResponse[*entities.Contact]) *list.PagiResponse[*ContactListDto] {
	var dtos []*ContactListDto
	for _, contact := range res.List {
		dtos = append(dtos, &ContactListDto{
			Id:        contact.Id,
			Message:   contact.Message,
			Email:     contact.Email,
			IsSeen:    contact.IsSeen,
			CreatedAt: contact.CreatedAt,
		})
	}
	return &list.PagiResponse[*ContactListDto]{
		List:      dtos,
		Page:      res.Page,
		Total:     res.Total,
		TotalPage: res.TotalPage,
		Limit:     res.Limit,
	}
}

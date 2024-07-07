package repos

import (
	"context"

	"github.com/9ssi7/gopre-starter/internal/domain/abstracts"
	"github.com/9ssi7/gopre-starter/internal/domain/entities"
	"github.com/9ssi7/gopre-starter/pkg/list"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type contactRepo struct {
	db *gorm.DB
}

func NewContactRepo(db *gorm.DB) abstracts.ContactRepo {
	return &contactRepo{
		db: db,
	}
}

func (r *contactRepo) Save(ctx context.Context, contact *entities.Contact) error {
	if err := r.db.WithContext(ctx).Save(contact).Error; err != nil {
		return err
	}
	return nil
}

func (r *contactRepo) FindById(ctx context.Context, id uuid.UUID) (*entities.Contact, error) {
	var contact entities.Contact
	if err := r.db.WithContext(ctx).Model(&entities.Contact{}).Where("id = ?", id).First(&contact).Error; err != nil {
		return nil, err
	}
	return &contact, nil
}

func (r *contactRepo) FindAll(ctx context.Context, req *list.PagiRequest) (*list.PagiResponse[*entities.Contact], error) {
	var contacts []*entities.Contact
	query := r.db.WithContext(ctx).Model(&entities.Contact{})
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	if err := query.Limit(*req.Limit).Offset(req.Offset()).Find(&contacts).Error; err != nil {
		return nil, err
	}
	return &list.PagiResponse[*entities.Contact]{
		List:          contacts,
		Total:         total,
		Limit:         *req.Limit,
		TotalPage:     req.TotalPage(total),
		FilteredTotal: total,
		Page:          *req.Page,
	}, nil
}

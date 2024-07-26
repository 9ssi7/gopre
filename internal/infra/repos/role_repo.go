package repos

import (
	"context"

	"github.com/9ssi7/gopre/internal/domain/abstracts"
	"github.com/9ssi7/gopre/internal/domain/entities"
	"github.com/9ssi7/gopre/pkg/list"
	"github.com/9ssi7/gopre/pkg/rescode"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type roleRepo struct {
	syncRepo
	txnGormRepo
	db *gorm.DB
}

func NewRoleRepo(db *gorm.DB) abstracts.RoleRepo {
	return &roleRepo{
		db:          db,
		txnGormRepo: newTxnGormRepo(db),
	}
}

func (r *roleRepo) Save(ctx context.Context, role *entities.Role) error {
	r.syncRepo.Lock()
	defer r.syncRepo.Unlock()
	if err := r.adapter.GetCurrent(ctx).Save(role).Error; err != nil {
		return err
	}
	return nil
}

func (r *roleRepo) FindById(ctx context.Context, id uuid.UUID) (*entities.Role, error) {
	var role entities.Role
	if err := r.adapter.GetCurrent(ctx).Model(&entities.Role{}).Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepo) Filter(ctx context.Context, req *list.PagiRequest, search string, isActive string) (*list.PagiResponse[*entities.Role], error) {
	var roles []*entities.Role
	query := r.adapter.GetCurrent(ctx).Model(&entities.Role{})
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, rescode.Failed(err)
	}
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}
	if isActive != "" {
		query = query.Where("is_active = ?", isActive)
	}
	var filteredTotal int64
	if err := query.Count(&filteredTotal).Error; err != nil {
		return nil, rescode.Failed(err)
	}
	if err := query.Offset(req.Offset()).Limit(*req.Limit).Find(&roles).Error; err != nil {
		return nil, rescode.Failed(err)
	}
	return &list.PagiResponse[*entities.Role]{
		List:          roles,
		Total:         total,
		Limit:         *req.Limit,
		TotalPage:     req.TotalPage(filteredTotal),
		FilteredTotal: filteredTotal,
		Page:          *req.Page,
	}, nil
}

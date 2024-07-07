package repos

import (
	"context"

	"github.com/9ssi7/gopre-starter/internal/domain/abstracts"
	"github.com/9ssi7/gopre-starter/internal/domain/entities"
	"github.com/9ssi7/gopre-starter/pkg/list"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) abstracts.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Save(ctx context.Context, user *entities.User) error {
	if err := r.db.WithContext(ctx).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepo) FindByToken(ctx context.Context, token string) (*entities.User, error) {
	var user entities.User
	if err := r.db.WithContext(ctx).Model(&entities.User{}).Where("temp_token = ? AND verified_at IS NULL", token).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindById(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	var user entities.User
	if err := r.db.WithContext(ctx).Model(&entities.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.WithContext(ctx).Model(&entities.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindByPhone(ctx context.Context, phone string) (*entities.User, error) {
	var user entities.User
	if err := r.db.WithContext(ctx).Model(&entities.User{}).Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) Filter(ctx context.Context, req *list.PagiRequest, search string, isActive string) (*list.PagiResponse[*entities.User], error) {
	var users []*entities.User
	query := r.db.WithContext(ctx).Model(&entities.User{})
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	if search != "" {
		query = query.Where("name LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	if isActive != "" {
		query = query.Where("is_active = ?", isActive)
	}
	var filteredTotal int64
	if err := query.Count(&filteredTotal).Error; err != nil {
		return nil, err
	}
	if err := query.Limit(*req.Limit).Offset(req.Offset()).Find(&users).Error; err != nil {
		return nil, err
	}
	return &list.PagiResponse[*entities.User]{
		List:          users,
		Total:         total,
		Limit:         *req.Limit,
		TotalPage:     req.TotalPage(filteredTotal),
		FilteredTotal: filteredTotal,
		Page:          *req.Page,
	}, nil
}

func (r *userRepo) FilterByRoles(ctx context.Context, req *list.PagiRequest, roles []string) (*list.PagiResponse[*entities.User], error) {
	var users []*entities.User
	query := r.db.WithContext(ctx).Model(&entities.User{})
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	if len(roles) > 0 {
		query = query.Where("roles && ?", pq.StringArray(roles))
	}
	var filteredTotal int64
	if err := query.Count(&filteredTotal).Error; err != nil {
		return nil, err
	}
	if err := query.Limit(*req.Limit).Offset(req.Offset()).Find(&users).Error; err != nil {
		return nil, err
	}
	return &list.PagiResponse[*entities.User]{
		List:          users,
		Total:         total,
		Limit:         *req.Limit,
		TotalPage:     req.TotalPage(filteredTotal),
		FilteredTotal: filteredTotal,
		Page:          *req.Page,
	}, nil
}

func (r *userRepo) ListByIds(ctx context.Context, userIds []uuid.UUID) ([]*entities.User, error) {
	var users []*entities.User
	if err := r.db.WithContext(ctx).Model(&entities.User{}).Where("id IN ?", userIds).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

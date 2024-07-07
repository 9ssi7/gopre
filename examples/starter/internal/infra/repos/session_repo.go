package repos

import (
	"context"
	"encoding/json"

	"github.com/9ssi7/gopre-starter/internal/domain/abstracts"
	"github.com/9ssi7/gopre-starter/internal/domain/aggregates"
	"github.com/9ssi7/gopre-starter/internal/infra/keyval"
	"github.com/9ssi7/gopre-starter/pkg/rescode"
	"github.com/google/uuid"
)

type sessionRepo struct {
	db keyval.DB
}

func NewSessionRepo(db keyval.DB) abstracts.SessionRepo {
	return &sessionRepo{
		db: db,
	}
}

func (s *sessionRepo) Save(ctx context.Context, userId uuid.UUID, session *aggregates.Session) error {
	key := s.calcKey(userId, session.DeviceId)
	bytes, err := json.Marshal(session)
	if err != nil {
		return rescode.Failed
	}
	if err := s.checkExistAndDel(ctx, key); err != nil {
		return rescode.Failed
	}
	if err := s.db.Set(ctx, key, bytes); err != nil {
		return rescode.Failed
	}
	return nil
}

func (s *sessionRepo) FindByIds(ctx context.Context, userId uuid.UUID, deviceId string) (*aggregates.Session, error) {
	key := s.calcKey(userId, deviceId)
	e, _, err := s.getByKey(ctx, key)
	if err != nil {
		return nil, err
	}
	if e == nil {
		return nil, rescode.NotFound
	}
	return e, nil
}

func (s *sessionRepo) FindAllByUserId(ctx context.Context, userId uuid.UUID) ([]*aggregates.Session, error) {
	keys, err := s.db.Keys(ctx, s.calcKey(userId, "*"))
	if err != nil {
		return nil, rescode.Failed
	}
	entities := make([]*aggregates.Session, len(keys))
	for i, k := range keys {
		e, _, err := s.getByKey(ctx, k)
		if err != nil {
			return nil, err
		}
		entities[i] = e
	}
	return entities, nil
}

func (s *sessionRepo) checkExistAndDel(ctx context.Context, key string) error {
	exist, err := s.db.Exist(ctx, key)
	if err != nil {
		return err
	}
	if exist {
		return s.db.Del(ctx, key)
	}
	return nil
}

func (s *sessionRepo) calcKey(userId uuid.UUID, deviceId string) string {
	return deviceId + "__" + userId.String()
}

func (s *sessionRepo) getByKey(ctx context.Context, key string) (*aggregates.Session, bool, error) {
	res, err := s.db.Get(ctx, key)
	if err != nil {
		return nil, true, rescode.Failed
	}
	var e aggregates.Session
	if err := json.Unmarshal([]byte(res), &e); err != nil {
		return nil, false, rescode.Failed
	}
	return &e, false, nil
}

func (s *sessionRepo) Destroy(ctx context.Context, userId uuid.UUID, deviceId string) error {
	key := s.calcKey(userId, deviceId)
	if err := s.db.Del(ctx, key); err != nil {
		return rescode.Failed
	}
	return nil
}

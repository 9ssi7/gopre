package repos

import (
	"context"
	"encoding/json"

	"github.com/9ssi7/gopre/internal/domain/abstracts"
	"github.com/9ssi7/gopre/internal/domain/aggregates"
	"github.com/9ssi7/gopre/pkg/rescode"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type sessionRepo struct {
	syncRepo
	db *redis.Client
}

func NewSessionRepo(db *redis.Client) abstracts.SessionRepo {
	return &sessionRepo{
		db: db,
	}
}

func (s *sessionRepo) Save(ctx context.Context, userId uuid.UUID, session *aggregates.Session) error {
	s.syncRepo.Lock()
	defer s.syncRepo.Unlock()
	key := s.calcKey(userId, session.DeviceId)
	bytes, err := json.Marshal(session)
	if err != nil {
		return rescode.Failed(err)
	}
	if err := s.checkExistAndDel(ctx, key); err != nil {
		return rescode.Failed(err)
	}
	if err := s.db.Set(ctx, key, bytes, 0).Err(); err != nil {
		return rescode.Failed(err)
	}
	return nil
}

func (s *sessionRepo) FindByIds(ctx context.Context, userId uuid.UUID, deviceId string) (*aggregates.Session, bool, error) {
	key := s.calcKey(userId, deviceId)
	e, notExists, err := s.getByKey(ctx, key)
	if err != nil {
		return nil, false, rescode.Failed(err)
	}
	if notExists {
		return nil, true, nil
	}
	return e, false, nil
}

func (s *sessionRepo) FindAllByUserId(ctx context.Context, userId uuid.UUID) ([]*aggregates.Session, error) {
	keys, err := s.db.Keys(ctx, s.calcKey(userId, "*")).Result()
	if err != nil {
		return nil, rescode.Failed(err)
	}
	entities := make([]*aggregates.Session, len(keys))
	for i, k := range keys {
		e, _, err := s.getByKey(ctx, k)
		if err != nil {
			return nil, rescode.Failed(err)
		}
		entities[i] = e
	}
	return entities, nil
}

func (s *sessionRepo) checkExistAndDel(ctx context.Context, key string) error {
	exist, err := s.db.Exists(ctx, key).Result()
	if err != nil {
		return rescode.Failed(err)
	}
	if exist == 1 {
		return s.db.Del(ctx, key).Err()
	}
	return nil
}

func (s *sessionRepo) calcKey(userId uuid.UUID, deviceId string) string {
	return deviceId + "__" + userId.String()
}

func (s *sessionRepo) getByKey(ctx context.Context, key string) (*aggregates.Session, bool, error) {
	res, err := s.db.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, true, nil
		}
		return nil, true, rescode.Failed(err)
	}
	var e aggregates.Session
	if err := json.Unmarshal([]byte(res), &e); err != nil {
		return nil, false, rescode.Failed(err)
	}
	return &e, false, nil
}

func (s *sessionRepo) Destroy(ctx context.Context, userId uuid.UUID, deviceId string) error {
	key := s.calcKey(userId, deviceId)
	if err := s.db.Del(ctx, key).Err(); err != nil {
		return rescode.Failed(err)
	}
	return nil
}

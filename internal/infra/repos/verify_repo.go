package repos

import (
	"context"
	"encoding/json"
	"time"

	"github.com/9ssi7/gopre/internal/domain/abstracts"
	"github.com/9ssi7/gopre/internal/domain/aggregates"
	"github.com/9ssi7/gopre/pkg/rescode"
	"github.com/redis/go-redis/v9"
)

type verifyRepo struct {
	syncRepo
	db *redis.Client
}

func NewVerifyRepo(db *redis.Client) abstracts.VerifyRepo {
	return &verifyRepo{
		db: db,
	}
}

func (r *verifyRepo) Save(ctx context.Context, token string, verify *aggregates.Verify) error {
	r.syncRepo.Lock()
	defer r.syncRepo.Unlock()
	b, err := json.Marshal(verify)
	if err != nil {
		return rescode.Failed(err)
	}
	if err = r.db.SetEx(ctx, r.calcKey(verify.DeviceId, token), b, 5*time.Minute).Err(); err != nil {
		return rescode.Failed(err)
	}
	return nil
}

func (r *verifyRepo) IsExists(ctx context.Context, token string, deviceId string) (bool, error) {
	res, err := r.db.Get(ctx, r.calcKey(deviceId, token)).Result()
	if err != nil {
		return false, rescode.Failed(err)
	}
	return res != "", nil
}

func (r *verifyRepo) Find(ctx context.Context, token string, deviceId string) (*aggregates.Verify, error) {
	res, err := r.db.Get(ctx, r.calcKey(deviceId, token)).Result()
	if err != nil {
		return nil, rescode.Failed(err)
	}
	var e aggregates.Verify
	if err = json.Unmarshal([]byte(res), &e); err != nil {
		return nil, rescode.Failed(err)
	}
	return &e, nil
}

func (r *verifyRepo) Delete(ctx context.Context, token string, deviceId string) error {
	if err := r.db.Del(ctx, r.calcKey(deviceId, token)).Err(); err != nil {
		return rescode.Failed(err)
	}
	return nil
}

func (r *verifyRepo) calcKey(deviceId string, token string) string {
	return "verify" + "__" + token + "__" + deviceId
}

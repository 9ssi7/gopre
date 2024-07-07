package repos

import (
	"context"
	"encoding/json"
	"time"

	"github.com/9ssi7/gopre-starter/internal/domain/abstracts"
	"github.com/9ssi7/gopre-starter/internal/domain/aggregates"
	"github.com/9ssi7/gopre-starter/internal/infra/keyval"
	"github.com/9ssi7/gopre-starter/pkg/rescode"
)

type verifyRepo struct {
	db keyval.DB
}

func NewVerifyRepo(db keyval.DB) abstracts.VerifyRepo {
	return &verifyRepo{
		db: db,
	}
}

func (r *verifyRepo) Save(ctx context.Context, token string, verify *aggregates.Verify) error {
	b, err := json.Marshal(verify)
	if err != nil {
		return rescode.Failed
	}
	if err = r.db.SetEx(ctx, r.calcKey(verify.DeviceId, token), b, 5*time.Minute); err != nil {
		return rescode.Failed
	}
	return nil
}

func (r *verifyRepo) IsExists(ctx context.Context, token string, deviceId string) (bool, error) {
	res, err := r.db.Get(ctx, r.calcKey(deviceId, token))
	if err != nil {
		return false, rescode.Failed
	}
	return res != "", nil
}

func (r *verifyRepo) Find(ctx context.Context, token string, deviceId string) (*aggregates.Verify, error) {
	res, err := r.db.Get(ctx, r.calcKey(deviceId, token))
	if err != nil {
		return nil, rescode.Failed
	}
	var e aggregates.Verify
	if err = json.Unmarshal([]byte(res), &e); err != nil {
		return nil, rescode.Failed
	}
	return &e, nil
}

func (r *verifyRepo) Delete(ctx context.Context, token string, deviceId string) error {
	if err := r.db.Del(ctx, r.calcKey(deviceId, token)); err != nil {
		return rescode.Failed
	}
	return nil
}

func (r *verifyRepo) calcKey(deviceId string, token string) string {
	return "verify" + "__" + token + "__" + deviceId
}

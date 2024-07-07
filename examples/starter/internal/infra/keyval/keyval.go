package keyval

import (
	"context"
	"time"

	"github.com/9ssi7/gopre-starter/config"
	"github.com/go-redis/redis/v8"
)

type DB interface {
	Set(ctx context.Context, k string, v interface{}) error
	HSet(ctx context.Context, k string, v ...interface{}) error
	Get(ctx context.Context, k string) (string, error)
	HGet(ctx context.Context, k string, field string) (string, error)
	HGetAll(ctx context.Context, k string) (map[string]string, error)
	SetEx(ctx context.Context, k string, v interface{}, d time.Duration) error
	Del(ctx context.Context, k ...string) error
	Exist(ctx context.Context, k string) (bool, error)
	Keys(ctx context.Context, prefix string) ([]string, error)
	Ping(ctx context.Context) error
}

var db DB

type redisClient struct {
	client *redis.Client
}

type Config struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func ConnectDB() DB {
	rClient := redis.NewClient(&redis.Options{
		Addr:     config.ReadValue().KeyValueDb.Host + ":" + config.ReadValue().KeyValueDb.Port,
		Password: config.ReadValue().KeyValueDb.Pw,
		DB:       config.ReadValue().KeyValueDb.Db,
	})
	db = &redisClient{
		client: rClient,
	}
	return db
}

func GetDB() DB {
	return db
}

func (r *redisClient) Set(ctx context.Context, k string, v interface{}) error {
	return r.client.Set(ctx, k, v, 0).Err()
}

func (r *redisClient) HSet(ctx context.Context, k string, v ...interface{}) error {
	return r.client.HSet(ctx, k, v).Err()
}

func (r *redisClient) Get(ctx context.Context, k string) (string, error) {
	return r.client.Get(ctx, k).Result()
}

func (r *redisClient) HGet(ctx context.Context, k string, field string) (string, error) {
	return r.client.HGet(ctx, k, field).Result()
}

func (r *redisClient) HGetAll(ctx context.Context, k string) (map[string]string, error) {
	return r.client.HGetAll(ctx, k).Result()
}

func (r *redisClient) SetEx(ctx context.Context, k string, v interface{}, d time.Duration) error {
	_, err := r.client.Set(ctx, k, v, d).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *redisClient) Del(ctx context.Context, k ...string) error {
	return r.client.Del(ctx, k...).Err()
}

func (r *redisClient) Exist(ctx context.Context, k string) (bool, error) {
	res, err := r.client.Exists(ctx, k).Result()
	if err != nil {
		return false, err
	}
	return res == 1, nil
}

func (r *redisClient) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

func (r *redisClient) Keys(ctx context.Context, pattern string) ([]string, error) {
	res, err := r.client.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, err
	}
	return res, nil
}

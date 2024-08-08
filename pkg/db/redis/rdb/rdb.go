package rdb

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	db "github.com/VadimGossip/platform_common/pkg/db/redis"
)

type rdb struct {
	dbc *redis.Client
}

func NewDB(dbc *redis.Client) db.DB {
	return &rdb{
		dbc: dbc,
	}
}
func (db *rdb) Get(ctx context.Context, key string) *redis.StringCmd {
	return db.dbc.Get(ctx, key)
}

func (db *rdb) GetWithDur(ctx context.Context, keys string) (string, time.Duration, error) {
	ts := time.Now()
	value, err := db.dbc.Get(ctx, keys).Result()
	return value, time.Since(ts), err
}

func (db *rdb) MGet(ctx context.Context, keys ...string) *redis.SliceCmd {
	return db.dbc.MGet(ctx, keys...)
}

func (db *rdb) MGetWithDur(ctx context.Context, keys ...string) ([]any, time.Duration, error) {
	ts := time.Now()
	values, err := db.dbc.MGet(ctx, keys...).Result()
	return values, time.Since(ts), err
}

func (db *rdb) HGetAll(ctx context.Context, key string, dest interface{}) error {
	return db.dbc.HGetAll(ctx, key).Scan(dest)
}

func (db *rdb) HSet(ctx context.Context, key string, values interface{}, expire time.Duration) error {
	if err := db.dbc.HSet(ctx, key, values).Err(); err != nil {
		return err
	}
	if expire > 0 {
		return db.dbc.Expire(ctx, key, expire).Err()
	}

	return nil
}

func (db *rdb) Del(ctx context.Context, keys ...string) error {
	return db.dbc.Del(ctx, keys...).Err()
}

func (db *rdb) Ping(ctx context.Context) error {
	return db.dbc.Ping(ctx).Err()
}

func (db *rdb) Close() error {
	return db.dbc.Close()
}

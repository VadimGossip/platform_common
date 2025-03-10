package rdb

import (
	"time"

	"github.com/redis/go-redis/v9"

	db "github.com/VadimGossip/platform_common/pkg/db/redis"
)

type rdbClient struct {
	masterDBC db.DB
}

type ClientOptions struct {
	Addr         string
	Username     string
	Password     string
	DB           int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func New(options *ClientOptions) db.Client {
	dbc := redis.NewClient(&redis.Options{
		Addr:         options.Addr,
		Username:     options.Username,
		Password:     options.Password,
		DB:           options.DB,
		ReadTimeout:  options.ReadTimeout,
		WriteTimeout: options.WriteTimeout,
	})
	return &rdbClient{
		masterDBC: NewDB(dbc),
	}
}

func (c *rdbClient) DB() db.DB {
	return c.masterDBC
}

func (c *rdbClient) Close() error {
	if c.masterDBC != nil {
		return c.masterDBC.Close()
	}

	return nil
}

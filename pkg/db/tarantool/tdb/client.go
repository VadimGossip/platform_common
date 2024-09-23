package tdb

import (
	"context"
	"time"

	"github.com/tarantool/go-tarantool/v2"

	db "github.com/VadimGossip/platform_common/pkg/db/tarantool"
)

type tdbClient struct {
	masterDBC db.DB
}

type ClientOptions struct {
	Addr     string
	Username string
	Password string
	Timeout  time.Duration
}

func New(ctx context.Context, options ClientOptions) (db.Client, error) {
	dialer := tarantool.NetDialer{
		Address:  options.Addr,
		User:     options.Username,
		Password: options.Password,
	}
	opts := tarantool.Opts{
		Timeout: options.Timeout,
	}

	conn, err := tarantool.Connect(ctx, dialer, opts)
	if err != nil {
		return nil, err
	}

	return &tdbClient{
		masterDBC: NewDB(conn),
	}, nil
}

func (c *tdbClient) DB() db.DB {
	return c.masterDBC
}

func (c *tdbClient) Close() error {
	if c.masterDBC != nil {
		return c.masterDBC.Close()
	}

	return nil
}

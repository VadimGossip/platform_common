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

func New(ctx context.Context, address, user, password string) (db.Client, error) {
	dialer := tarantool.NetDialer{
		Address:  address,
		User:     user,
		Password: password,
	}
	opts := tarantool.Opts{
		Timeout: 1 * time.Second,
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

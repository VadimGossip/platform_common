package chdb

import (
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/pkg/errors"

	db "github.com/VadimGossip/platform_common/pkg/db/clickhouse"
)

type chdbClient struct {
	masterDBC db.DB
}

func New(options *clickhouse.Options) (db.Client, error) {
	dbc, err := clickhouse.Open(options)
	if err != nil {
		return nil, errors.Errorf("failed to connect to db: %v", err)
	}

	return &chdbClient{
		masterDBC: NewDB(dbc),
	}, nil
}

func (c *chdbClient) DB() db.DB {
	return c.masterDBC
}

func (c *chdbClient) Close() error {
	if c.masterDBC != nil {
		return c.masterDBC.Close()
	}

	return nil
}

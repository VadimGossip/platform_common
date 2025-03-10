package chdb

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"

	db "github.com/VadimGossip/platform_common/pkg/db/clickhouse"
)

type chdb struct {
	dbc driver.Conn
}

func NewDB(dbc driver.Conn) db.DB {
	return &chdb{
		dbc: dbc,
	}
}

func (db *chdb) PrepareBatch(ctx context.Context, query string, opts ...driver.PrepareBatchOption) (driver.Batch, error) {
	return db.dbc.PrepareBatch(ctx, query, opts...)
}

func (db *chdb) SendBatch(batch driver.Batch) error {
	return batch.Send()
}

func (db *chdb) Ping(ctx context.Context) error {
	return db.dbc.Ping(ctx)
}

func (db *chdb) Close() error {
	return db.dbc.Close()
}

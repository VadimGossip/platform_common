package clickhouse

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Client interface {
	DB() DB
	Close() error
}

type Execer interface {
	PrepareBatch(ctx context.Context, query string, opts ...driver.PrepareBatchOption) (driver.Batch, error)
	SendBatch(batch driver.Batch) error
}

type Pinger interface {
	Ping(ctx context.Context) error
}

type DB interface {
	Execer
	Pinger
	Close() error
}

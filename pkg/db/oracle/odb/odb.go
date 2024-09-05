package odb

import (
	"context"
	"database/sql"

	db "github.com/VadimGossip/platform_common/pkg/db/oracle"
)

type key string

// TxKey key name for tx in context
const (
	TxKey key = "tx"
)

type odb struct {
	dbc *sql.DB
}

func NewDB(dbc *sql.DB) db.DB {
	return &odb{
		dbc: dbc,
	}
}

func (o *odb) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	tx, ok := ctx.Value(TxKey).(*sql.Tx)
	if ok {
		return tx.Query(query, args...)
	}

	return o.dbc.QueryContext(ctx, query, args...)
}

func (o *odb) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	tx, ok := ctx.Value(TxKey).(*sql.Tx)
	if ok {
		return tx.QueryRow(query, args...)
	}

	return o.dbc.QueryRowContext(ctx, query, args...)
}

func (o *odb) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	tx, ok := ctx.Value(TxKey).(*sql.Tx)
	if ok {
		return tx.Exec(query, args...)
	}

	return o.dbc.ExecContext(ctx, query, args...)
}

func (o *odb) Ping(ctx context.Context) error {
	return o.dbc.PingContext(ctx)
}

func (o *odb) Close() error {
	return o.dbc.Close()
}

func (o *odb) BeginTx(ctx context.Context, txOptions *sql.TxOptions) (*sql.Tx, error) {
	return o.dbc.BeginTx(ctx, txOptions)
}

func MakeContextTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}

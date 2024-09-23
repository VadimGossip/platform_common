package tdb

import (
	"github.com/tarantool/go-tarantool/v2"

	db "github.com/VadimGossip/platform_common/pkg/db/tarantool"
)

type tdb struct {
	dbc *tarantool.Connection
}

func NewDB(dbc *tarantool.Connection) db.DB {
	return &tdb{
		dbc: dbc,
	}
}

func (t *tdb) Do(req tarantool.Request) *tarantool.Future {
	return t.dbc.Do(req)
}

func (t *tdb) Close() error {
	return t.dbc.Close()
}

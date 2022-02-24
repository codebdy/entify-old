package sql

import (
	"context"
	"database/sql"
)

type ExecQuerier interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

type Conn struct {
	ExecQuerier
}

type Driver struct {
	Conn
	dialect string
}

func (d Driver) DB() *sql.DB {
	return d.ExecQuerier.(*sql.DB)
}

func Open(driver, source string) (*Driver, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		return nil, err
	}
	return &Driver{Conn{db}, driver}, nil
}

func (d *Driver) Close() error { return d.DB().Close() }

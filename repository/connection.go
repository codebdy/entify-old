package repository

import (
	"rxdrag.com/entity-engine/config"
	"rxdrag.com/entity-engine/db"
)

type Connection struct {
	dbx *db.Dbx
}

func Open() (*Connection, error) {
	dbx, err := db.Open(config.DRIVER_NAME, config.MYSQL_CONFIG)
	if err != nil {
		return nil, err
	}
	con := Connection{
		dbx: dbx,
	}
	return &con, err
}

func (c *Connection) Close() error {
	return c.dbx.Close()
}

func (c *Connection) BeginTx() error {
	return c.dbx.BeginTx()
}

func (c *Connection) Commit() error {
	return c.dbx.Commit()
}

func (c *Connection) ClearTx() {
	c.dbx.ClearTx()
}

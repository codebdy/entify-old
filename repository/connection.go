package repository

import (
	"rxdrag.com/entity-engine/config"
	"rxdrag.com/entity-engine/db"
)

type Connection struct {
	Dbx *db.Dbx
}

func Open() (*Connection, error) {
	dbx, err := db.Open(config.DRIVER_NAME, config.MYSQL_CONFIG)
	if err != nil {
		return nil, err
	}
	con := Connection{
		Dbx: dbx,
	}
	return &con, err
}

func (c *Connection) Close() error {
	return c.Dbx.Close()
}

func (c *Connection) BeginTx() error {
	return c.Dbx.BeginTx()
}

func (c *Connection) Commit() error {
	return c.Dbx.Commit()
}

func (c *Connection) ClearTx() {
	c.Dbx.ClearTx()
}

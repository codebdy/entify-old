package repository

import (
	"rxdrag.com/entify/config"
	"rxdrag.com/entify/db"
)

type Connection struct {
	idSeed int //use for sql join table
	Dbx    *db.Dbx
}

func Open() (*Connection, error) {
	dbx, err := db.Open(config.DRIVER_NAME, config.MYSQL_CONFIG)
	if err != nil {
		return nil, err
	}
	con := Connection{
		idSeed: 1,
		Dbx:    dbx,
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

//use for sql join table
func (c *Connection) CreateId() int {
	c.idSeed++
	return c.idSeed
}

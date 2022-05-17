package repository

import (
	"rxdrag.com/entify/config"
	"rxdrag.com/entify/db"
)

type Connection struct {
	idSeed int //use for sql join table
	Dbx    *db.Dbx
}

func OpenWithConfig(cfg config.DbConfig) (*Connection, error) {
	dbx, err := db.Open(cfg.Driver, DbString(cfg))
	if err != nil {
		return nil, err
	}
	con := Connection{
		idSeed: 1,
		Dbx:    dbx,
	}
	return &con, err
}

func Open() (*Connection, error) {
	cfg := config.GetDbConfig()
	return OpenWithConfig(cfg)
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

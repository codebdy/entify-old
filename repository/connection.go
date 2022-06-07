package repository

import (
	"rxdrag.com/entify/config"
	"rxdrag.com/entify/db"
)

type Connection struct {
	idSeed int //use for sql join table
	Dbx    *db.Dbx
	v      *AbilityVerifier
}

func openWithConfig(cfg config.DbConfig, v *AbilityVerifier) (*Connection, error) {
	dbx, err := db.Open(cfg.Driver, DbString(cfg))
	if err != nil {
		return nil, err
	}
	con := Connection{
		idSeed: 1,
		Dbx:    dbx,
		v:      v,
	}
	return &con, err
}

func Open(v *AbilityVerifier) (*Connection, error) {
	cfg := config.GetDbConfig()
	return openWithConfig(cfg, v)
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

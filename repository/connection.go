package repository

import (
	"fmt"

	"rxdrag.com/entify/config"
	"rxdrag.com/entify/db"
)

type Connection struct {
	idSeed int //use for sql join table
	Dbx    *db.Dbx
}

func Open(cfg config.DbConfig) (*Connection, error) {
	fmt.Println("呵呵", DbString(cfg), cfg.Driver)
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

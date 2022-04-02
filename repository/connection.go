package repository

import (
	"rxdrag.com/entity-engine/config"
	"rxdrag.com/entity-engine/dbx"
)

type Connection struct {
	dbx *db.Dbx
}

func Open() (*Connection, error) {
	dbx, err := dbx.Open(config.DRIVER_NAME, config.MYSQL_CONFIG)
	if err != nil {
		return nil, err
	}
	con := Connection{
		dbx: dbx,
	}
	return &con, err
}

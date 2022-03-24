package repository

import (
	"database/sql"
	"log"

	"rxdrag.com/entity-engine/config"
)

type Connection struct {
	db *sql.DB
	tx *sql.Tx
}

func (c *Connection) Begin() error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	c.tx = tx
	return nil
}

func (c *Connection) validateDb() {
	if c.db == nil {
		panic("Not init connection with db")
	}
}

func (c *Connection) validateTx() {
	if c.tx == nil {
		panic("Not init connection with tx")
	}
}

func (c *Connection) Exec(sql string, args ...interface{}) (sql.Result, error) {
	c.validateDb()
	if c.tx != nil {
		return c.tx.Exec(sql, args...)
	}
	return c.db.Exec(sql, args...)
}

func (c *Connection) Query(query string, args ...interface{}) (*sql.Rows, error) {
	c.validateDb()
	if c.tx != nil {
		return c.tx.Query(query, args...)
	} else {
		return c.db.Query(query, args...)
	}
}

func (c *Connection) QueryRow(query string, args ...interface{}) *sql.Row {
	c.validateDb()
	if c.tx != nil {
		return c.tx.QueryRow(query, args...)
	} else {
		return c.db.QueryRow(query, args...)
	}
}

func (c *Connection) Close() error {
	c.validateDb()
	return c.db.Close()
}

func (c *Connection) Commit() error {
	c.validateTx()
	return c.tx.Commit()
}
func (c *Connection) Rollback() error {
	c.validateTx()
	return c.tx.Rollback()
}

func (c *Connection) clearTx() {
	c.validateTx()
	err := c.Rollback()
	if err != sql.ErrTxDone && err != nil {
		log.Fatalln(err)
	}
}

func OpenConnection() (*Connection, error) {
	db, err := sql.Open(config.DRIVER_NAME, config.MYSQL_CONFIG)
	if err != nil {
		return nil, err
	}
	con := Connection{
		db: db,
	}
	return &con, err
}

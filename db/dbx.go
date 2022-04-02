package db

import (
	"database/sql"
	"log"
)

type Dbx struct {
	db *sql.DB
	tx *sql.Tx
}

func (c *Dbx) BeginTx() error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	c.tx = tx
	return nil
}

func (c *Dbx) ClearTx() {
	c.validateTx()
	err := c.Rollback()
	if err != sql.ErrTxDone && err != nil {
		log.Fatalln(err)
	}
}

func (c *Dbx) validateDb() {
	if c.db == nil {
		panic("Not init connection with db")
	}
}

func (c *Dbx) validateTx() {
	if c.tx == nil {
		panic("Not init connection with tx")
	}
}

func (c *Dbx) Exec(sql string, args ...interface{}) (sql.Result, error) {
	c.validateDb()
	if c.tx != nil {
		return c.tx.Exec(sql, args...)
	}
	return c.db.Exec(sql, args...)
}

func (c *Dbx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	c.validateDb()
	if c.tx != nil {
		return c.tx.Query(query, args...)
	} else {
		return c.db.Query(query, args...)
	}
}

func (c *Dbx) QueryRow(query string, args ...interface{}) *sql.Row {
	c.validateDb()
	if c.tx != nil {
		return c.tx.QueryRow(query, args...)
	} else {
		return c.db.QueryRow(query, args...)
	}
}

func (c *Dbx) Close() error {
	c.validateDb()
	return c.db.Close()
}

func (c *Dbx) Commit() error {
	c.validateTx()
	return c.tx.Commit()
}
func (c *Dbx) Rollback() error {
	c.validateTx()
	return c.tx.Rollback()
}

func Open(driver string, config string) (*Dbx, error) {
	db, err := sql.Open(driver, config)
	if err != nil {
		return nil, err
	}
	con := Dbx{
		db: db,
	}
	return &con, err
}

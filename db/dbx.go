package db

import (
	"database/sql"
	"log"
)

var openedDB *sql.DB

type Dbx struct {
	db *sql.DB
	tx *sql.Tx
}

func (d *Dbx) BeginTx() error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}

	d.tx = tx
	return nil
}

func (d *Dbx) ClearTx() {
	d.validateTx()
	err := d.Rollback()
	if err != sql.ErrTxDone && err != nil {
		log.Fatalln(err)
	}
}

func (d *Dbx) validateDb() {
	if d.db == nil {
		panic("Not init connection with db")
	}
}

func (d *Dbx) validateTx() {
	if d.tx == nil {
		panic("Not init connection with tx")
	}
}

func (d *Dbx) Exec(sql string, args ...interface{}) (sql.Result, error) {
	d.validateDb()
	if d.tx != nil {
		return d.tx.Exec(sql, args...)
	}
	return d.db.Exec(sql, args...)
}

func (d *Dbx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	d.validateDb()
	if d.tx != nil {
		return d.tx.Query(query, args...)
	} else {
		return d.db.Query(query, args...)
	}
}

func (d *Dbx) QueryRow(query string, args ...interface{}) *sql.Row {
	d.validateDb()
	if d.tx != nil {
		return d.tx.QueryRow(query, args...)
	} else {
		return d.db.QueryRow(query, args...)
	}
}

// func (c *Dbx) Close() error {
// 	c.validateDb()
// 	return c.db.Close()
// }

func (d *Dbx) Commit() error {
	d.validateTx()
	return d.tx.Commit()
}
func (c *Dbx) Rollback() error {
	c.validateTx()
	return c.tx.Rollback()
}

func Open(driver string, config string) (*Dbx, error) {
	if openedDB == nil {
		db, err := sql.Open(driver, config)
		openedDB = db
		if err != nil {
			return nil, err
		}
	}
	con := Dbx{
		db: openedDB,
	}
	return &con, nil
}

func Close() error {
	if openedDB != nil {
		err := openedDB.Close()
		openedDB = nil
		return err
	}

	return nil
}

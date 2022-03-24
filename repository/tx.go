package repository

import (
	"database/sql"
	"log"
)

type Tx struct {
	tx *sql.Tx
}

func NewTx(db *sql.DB) (*Tx, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	return &Tx{
		tx: tx,
	}, nil
}

func (tx *Tx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return tx.tx.Exec(query, args...)
}

func (tx *Tx) Commit() error {
	return tx.tx.Commit()
}
func (tx *Tx) Rollback() error {
	return tx.tx.Rollback()
}

func (tx *Tx) clear() {
	err := tx.Rollback()
	if err != sql.ErrTxDone && err != nil {
		log.Fatalln(err)
	}
}

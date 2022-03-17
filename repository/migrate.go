package repository

import (
	"database/sql"
	"fmt"
	"log"

	"rxdrag.com/entity-engine/config"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/repository/dialect"
)

func ExcuteDiff(d *meta.Diff) {
	var undoList []string
	db, err := sql.Open(config.DRIVER_NAME, config.MYSQL_CONFIG)
	defer db.Close()
	if err != nil {
		panic("Open db error:" + err.Error())
	}

	for _, table := range d.DeletedTables {
		err = DeleteTable(table, &undoList, db)
		if err != nil {
			rollback(undoList, db)
			panic("Delete table error:" + err.Error())
		}
	}

	for _, table := range d.AddedTables {
		err = CreateTable(table, &undoList, db)
		if err != nil {
			rollback(undoList, db)
			panic("Create table error:" + err.Error())
		}
	}

	for _, tableDiff := range d.ModifiedTables {
		err = ModifyTable(tableDiff, &undoList, db)
		if err != nil {
			rollback(undoList, db)
			panic("Modify table error:" + err.Error())
		}
	}

}

func UndoDiff(d *meta.Diff) {

}

func DeleteTable(table *meta.Table, undoList *[]string, db *sql.DB) error {
	sqlBuilder := dialect.GetSQLBuilder()
	excuteSQL := sqlBuilder.BuildDeleteTableSQL(table)
	undoSQL := sqlBuilder.BuildCreateTableSQL(table)
	*undoList = append(*undoList, undoSQL)
	_, err := db.Exec(excuteSQL)
	if err != nil {
		return err
	}
	log.Println("Delete Table SQL:", excuteSQL)
	return nil
}

func CreateTable(table *meta.Table, undoList *[]string, db *sql.DB) error {
	sqlBuilder := dialect.GetSQLBuilder()
	excuteSQL := sqlBuilder.BuildCreateTableSQL(table)
	undoSQL := sqlBuilder.BuildDeleteTableSQL(table)
	*undoList = append(*undoList, undoSQL)
	_, err := db.Exec(excuteSQL)
	if err != nil {
		return err
	}
	log.Println("Add Table SQL:", excuteSQL)

	return nil
}

func ModifyTable(entityDiff *meta.TableDiff, undoList *[]string, db *sql.DB) error {
	fmt.Println("Not implement ModifyEntity")
	return nil
}

func rollback(undoList []string, db *sql.DB) {
	for _, sql := range undoList {
		_, err := db.Exec(sql)
		if err != nil {
			log.Println("Rollaback failed:", sql)
		}
	}
}

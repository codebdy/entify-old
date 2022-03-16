package repository

import (
	"database/sql"
	"fmt"

	"rxdrag.com/entity-engine/config"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/repository/dialect"
)

func ExcuteDiff(d *meta.Diff) {
	var undoList []string
	db, err := sql.Open(config.DRIVER_NAME, config.MYSQL_CONFIG)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, table := range d.DeletedTables {
		DeleteTable(table.Name)
	}

	for _, table := range d.AddedTables {
		CreateTable(table, &undoList, db)
	}

	for _, tableDiff := range d.ModifiedTables {
		ModifyTable(tableDiff)
	}

}

func UndoDiff(d *meta.Diff) {

}

func DeleteTable(entityName string) {
	fmt.Println("Not implement DeleteEntity")
}

func CreateTable(table *meta.Table, undoList *[]string, db *sql.DB) {
	sqlBuilder := dialect.GetSQLBuilder()
	excuteSQL, undoSQL := sqlBuilder.BuildCreateTableSQL(table)
	*undoList = append(*undoList, undoSQL)
	_, err := db.Exec(excuteSQL)
	if err != nil {
		panic("Create table error:" + err.Error())
	}
	fmt.Println("AddEntity SQL:", excuteSQL)
}

func ModifyTable(entityDiff *meta.TableDiff) {
	fmt.Println("Not implement ModifyEntity")
}

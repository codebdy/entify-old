package repository

import (
	"fmt"
	"log"

	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/repository/dialect"
)

func ExcuteDiff(d *model.Diff) {
	var undoList []string
	con, err := OpenConnection()
	defer con.Close()
	if err != nil {
		panic("Open db error:" + err.Error())
	}

	for _, table := range d.DeletedTables {
		err = DeleteTable(table, &undoList, con)
		if err != nil {
			rollback(undoList, con)
			panic("Delete table error:" + err.Error())
		}
	}

	for _, table := range d.AddedTables {
		err = CreateTable(table, &undoList, con)
		if err != nil {
			rollback(undoList, con)
			panic("Create table error:" + err.Error())
		}
	}

	for _, tableDiff := range d.ModifiedTables {
		err = ModifyTable(tableDiff, &undoList, con)
		if err != nil {
			rollback(undoList, con)
			panic("Modify table error:" + err.Error())
		}
	}

}

func DeleteTable(table *model.Table, undoList *[]string, con *Connection) error {
	sqlBuilder := dialect.GetSQLBuilder()
	excuteSQL := sqlBuilder.BuildDeleteTableSQL(table)
	undoSQL := sqlBuilder.BuildCreateTableSQL(table)
	_, err := con.Exec(excuteSQL)
	if err != nil {
		return err
	}
	*undoList = append(*undoList, undoSQL)
	log.Println("Delete Table SQL:", excuteSQL)
	return nil
}

func CreateTable(table *model.Table, undoList *[]string, con *Connection) error {
	sqlBuilder := dialect.GetSQLBuilder()
	excuteSQL := sqlBuilder.BuildCreateTableSQL(table)
	undoSQL := sqlBuilder.BuildDeleteTableSQL(table)
	_, err := con.Exec(excuteSQL)
	if err != nil {
		return err
	}
	*undoList = append(*undoList, undoSQL)
	log.Println("Add Table SQL:", excuteSQL)

	return nil
}

func ModifyTable(tableDiff *model.TableDiff, undoList *[]string, con *Connection) error {
	sqlBuilder := dialect.GetSQLBuilder()
	atoms := sqlBuilder.BuildModifyTableAtoms(tableDiff)
	for _, atom := range atoms {
		_, err := con.Exec(atom.ExcuteSQL)
		if err != nil {
			fmt.Println("出错atom", atom.ExcuteSQL, err.Error())
			return err
		}
		*undoList = append(*undoList, atom.UndoSQL)
	}
	return nil
}

func rollback(undoList []string, con *Connection) {
	for _, sql := range undoList {
		_, err := con.Exec(sql)
		if err != nil {
			log.Println("Rollback failed:", sql)
		}
	}
}

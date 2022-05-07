package dialect

import (
	"testing"

	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/table"
)

func TestModifyTableName(t *testing.T) {
	var mysqlBuilder MySQLBuilder

	atoms := mysqlBuilder.BuildModifyTableAtoms(
		&model.TableDiff{
			OldTable: &table.Table{
				Name: "User",
			},
			NewTable: &table.Table{
				Name: "User2",
			},
		},
	)

	if len(atoms) != 1 {
		t.Error("Modify atoms number error")
	}

	if atoms[0].ExcuteSQL != "ALTER TABLE User RENAME TO User2 " {
		t.Error("Modify atom ExcuteSQL error:#" + atoms[0].ExcuteSQL + "#")
	}

	if atoms[0].UndoSQL != "ALTER TABLE User2 RENAME TO User " {
		t.Error("Modify atom UndoSQL error:#" + atoms[0].UndoSQL + "#")
	}
}

func TestModifyColumnName(t *testing.T) {
	var mysqlBuilder MySQLBuilder
	atoms := mysqlBuilder.BuildModifyTableAtoms(
		&model.TableDiff{
			OldTable: &table.Table{
				Uuid: "uuid1",
				Name: "User",
			},
			NewTable: &table.Table{
				Uuid: "uuid1",
				Name: "User",
			},
		},
	)

	if len(atoms) != 1 {
		t.Errorf("Modify atoms number error, number:%d", len(atoms))
	}

	if atoms[0].ExcuteSQL != "ALTER TABLE User CHANGE COLUMN newColumn1 nickname text" {
		t.Errorf("ExcuteSQL error:" + atoms[0].ExcuteSQL)
	}

	if atoms[0].UndoSQL != "ALTER TABLE User CHANGE COLUMN nickname newColumn1 text" {
		t.Errorf("UndoSQL error:" + atoms[0].UndoSQL)
	}
	t.Log("#" + atoms[0].ExcuteSQL + "#")
	t.Log(atoms[0].UndoSQL)
}

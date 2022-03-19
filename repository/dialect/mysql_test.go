package dialect

import (
	"testing"

	"rxdrag.com/entity-engine/meta"
)

func TestModifyTableName(t *testing.T) {
	var mysqlBuilder MySQLBuilder

	atoms := mysqlBuilder.BuildModifyTableAtoms(
		&meta.TableDiff{
			OldTable: &meta.Table{
				Name: "User",
			},
			NewTable: &meta.Table{
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
		&meta.TableDiff{
			OldTable: &meta.Table{
				MetaUuid: "uuid1",
				Name:     "User",
			},
			NewTable: &meta.Table{
				MetaUuid: "uuid1",
				Name:     "User",
			},
			ModifyColumns: []meta.ColumnDiff{
				{
					OldColumn: meta.Column{
						Name: "newColumn1",
						Uuid: "column1",
						Type: meta.COLUMN_STRING,
					},
					NewColumn: meta.Column{
						Name: "nickname",
						Uuid: "column1",
						Type: meta.COLUMN_STRING,
					},
				},
			},
		},
	)

	if len(atoms) != 1 {
		t.Errorf("Modify atoms number error, number:%d", len(atoms))
	}

}

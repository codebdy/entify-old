package meta

import (
	"encoding/json"
	"testing"
)

func TestModifyEntityName(t *testing.T) {
	oldData := `
	{
		"entities": [
			{
				"name": "User",
				"uuid": "3e9ae743-de18-4b0c-a77e-3726be4049a8",
				"columns": [
					{
						"name": "id",
						"type": "ID",
						"uuid": "6758ae89-1e2c-462d-907c-a54baf6cf6fd",
						"primary": true
					},
					{
						"name": "newColumn1",
						"type": "String",
						"uuid": "e1afb0c4-5eee-40f3-8c34-3ce15746877b"
					}
				],
				"entityType": "Normal"
			}
		],
		"relations": []
	}
	`

	newData := `
	{
		"entities": [
			{
				"name": "User2",
				"uuid": "3e9ae743-de18-4b0c-a77e-3726be4049a8",
				"columns": [
					{
						"name": "id",
						"type": "ID",
						"uuid": "6758ae89-1e2c-462d-907c-a54baf6cf6fd",
						"primary": true
					},
					{
						"name": "newColumn1",
						"type": "String",
						"uuid": "e1afb0c4-5eee-40f3-8c34-3ce15746877b"
					}
				],
				"entityType": "Normal"
			}
		],
		"relations": []
	}
	`

	oldM := MetaContent{}
	json.Unmarshal([]byte(oldData), &oldM)
	newM := MetaContent{}
	json.Unmarshal([]byte(newData), &newM)
	newM.Validate()
	diff := CreateDiff(&oldM, &newM)

	if len(diff.ModifiedTables) != 1 {
		t.Errorf("Diffent table number is %d ,not 1", len(diff.ModifiedTables))
	}

	if diff.oldContent.Tables()[0].Name != "user" {
		t.Errorf("Old name is %s, not expected user", diff.oldContent.Tables()[0].Name)
	}

	if diff.newContent.Tables()[0].Name != "user2" {
		t.Errorf("Old name is %s, not expected user2", diff.newContent.Tables()[0].Name)
	}
}

func TestModifiedTableName(t *testing.T) {
	diff := CreateDiff(
		&MetaContent{
			Entities: []EntityMeta{
				{
					Name: "OldName",
				},
			},
		},
		&MetaContent{
			Entities: []EntityMeta{
				{
					Name: "NewName",
				},
			},
		},
	)

	if len(diff.ModifiedTables) != 1 {
		t.Error("Cereate entity name modify diff error, diff number error")
	}

	if diff.ModifiedTables[0].OldTable.Name != "old_name" {
		t.Error("Cereate entity name modify diff error, old name error")
	}

	if diff.ModifiedTables[0].NewTable.Name != "new_name" {
		t.Error("Cereate entity name modify diff error, new name error")
	}

}

func TestColumnDifferent(t *testing.T) {
	diff := columnDifferent(
		&ColumnMeta{
			Name: "newColumn1",
			Uuid: "column1",
			Type: COLUMN_STRING,
		},
		&ColumnMeta{
			Name: "nickname",
			Uuid: "column1",
			Type: COLUMN_STRING},
	)

	if diff == nil {
		t.Errorf("columnDifferent return value is nil")
		return
	}

	if diff.OldColumn.Name != "newColumn1" {
		t.Errorf("expect old column newColumn1, but actual is %s", diff.OldColumn.Name)
	}

	if diff.NewColumn.Name != "nickname" {
		t.Errorf("expect new column nickname, but actual is %s", diff.NewColumn.Name)
	}
}

func TestChangeTableColumnName(t *testing.T) {
	diff := tableDifferent(
		&Table{
			Name:     "User",
			MetaUuid: "User-uuid",
			Columns: []ColumnMeta{
				{
					Name: "newColumn1",
					Uuid: "column1",
					Type: COLUMN_STRING,
				},
			},
		},
		&Table{
			Name:     "User",
			MetaUuid: "User-uuid",
			Columns: []ColumnMeta{
				{
					Name: "nickname",
					Uuid: "column1",
					Type: COLUMN_STRING,
				},
			},
		},
	)

	if diff == nil {
		t.Errorf("tableDifferent return value is nil")
		return
	}

	if len(diff.ModifyColumns) != 1 {
		t.Errorf("Column diff number is %d ,not 1", len(diff.ModifyColumns))
	}

	if diff.ModifyColumns[0].OldColumn.Name != "newColumn1" {
		t.Errorf("Column diff old column error: %s", diff.ModifyColumns[0].OldColumn.Name)
	}
}

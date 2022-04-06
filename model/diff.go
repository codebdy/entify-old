package model

import "rxdrag.com/entity-engine/model/table"

type ModifyAtom struct {
	ExcuteSQL string
	UndoSQL   string
}

type ColumnDiff struct {
	OldColumn *table.Column
	NewColumn *table.Column
}

type TableDiff struct {
	OldTable      *table.Table
	NewTable      *table.Table
	DeleteColumns []*table.Column
	AddColumns    []*table.Column
	ModifyColumns []ColumnDiff //删除列索引，并重建
}

type Diff struct {
	oldContent *Model
	newContent *Model

	DeletedTables  []*table.Table
	AddedTables    []*table.Table
	ModifiedTables []*TableDiff
}

func findTable(uuid string, tables []*table.Table) *table.Table {
	for i := range tables {
		if tables[i].Uuid == uuid {
			return tables[i]
		}
	}
	return nil
}

func findColumn(uuid string, columns []*table.Column) *table.Column {
	for _, column := range columns {
		if column.Uuid == uuid {
			return column
		}
	}

	return nil
}

func columnDifferent(oldColumn, newColumn *table.Column) *ColumnDiff {
	diff := ColumnDiff{
		OldColumn: oldColumn,
		NewColumn: newColumn,
	}
	if oldColumn.Name != newColumn.Name {
		return &diff
	}
	// if oldColumn.Generated != newColumn.Generated {
	// 	return &diff
	// }
	if oldColumn.Index != newColumn.Index {
		return &diff
	}
	if oldColumn.Nullable != newColumn.Nullable {
		return &diff
	}
	if oldColumn.Length != newColumn.Length {
		return &diff
	}
	if oldColumn.Primary != newColumn.Primary {
		return &diff
	}

	if oldColumn.Unique != newColumn.Unique {
		return &diff
	}

	if oldColumn.Type != newColumn.Type {
		return &diff
	}
	return nil
}
func tableDifferent(oldTable, newTable *table.Table) *TableDiff {
	var diff TableDiff
	modified := false
	diff.OldTable = oldTable
	diff.NewTable = newTable

	for _, column := range oldTable.Columns {
		foundCoumn := findColumn(column.Uuid, newTable.Columns)
		if foundCoumn == nil {
			diff.DeleteColumns = append(diff.DeleteColumns, column)
			modified = true
		}
	}

	for _, column := range newTable.Columns {
		foundColumn := findColumn(column.Uuid, oldTable.Columns)
		if foundColumn == nil {
			diff.AddColumns = append(diff.AddColumns, column)
			modified = true
		} else {
			columnDiff := columnDifferent(foundColumn, column)
			if columnDiff != nil {
				diff.ModifyColumns = append(diff.ModifyColumns, *columnDiff)
				modified = true
			}
		}
	}

	if diff.OldTable.Name != diff.NewTable.Name || modified {
		return &diff
	}
	return nil
}

func CreateDiff(published, next *Model) *Diff {
	diff := Diff{
		oldContent: published,
		newContent: next,
	}

	// publishedTables := published.Graph.Tables
	// nextTables := next.Graph.Tables

	// for _, table := range publishedTables {
	// 	foundTable := findTable(table.Uuid, nextTables)
	// 	//删除的Table
	// 	if foundTable == nil {
	// 		diff.DeletedTables = append(diff.DeletedTables, table)
	// 	}
	// }
	// for _, table := range nextTables {
	// 	foundTable := findTable(table.Uuid, publishedTables)
	// 	//添加的Table
	// 	if foundTable == nil {
	// 		diff.AddedTables = append(diff.AddedTables, table)
	// 	} else { //修改的Table
	// 		tableDiff := tableDifferent(foundTable, table)
	// 		if tableDiff != nil {
	// 			diff.ModifiedTables = append(diff.ModifiedTables, tableDiff)
	// 		}
	// 	}
	// }

	return &diff
}

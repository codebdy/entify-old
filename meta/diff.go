package meta

import "fmt"

type ModifyAtom struct {
	ExcuteSQL string
	UndoSQL   string
}

type ColumnDiff struct {
	OldColumn Column
	NewColumn Column
}

type TableDiff struct {
	OldTable      *Table
	NewTable      *Table
	DeleteColumns []Column
	AddColumns    []Column
	ModifyColumns []ColumnDiff //删除列索引，并重建
}

type Diff struct {
	oldContent *MetaContent
	newContent *MetaContent

	DeletedTables  []*Table
	AddedTables    []*Table
	ModifiedTables []*TableDiff
}

func findColumn(uuid string, columns []Column) *Column {
	for _, column := range columns {
		if column.Uuid == uuid {
			return &column
		}
	}

	return nil
}

func columnDifferent(oldColumn, newColumn *Column) *ColumnDiff {
	diff := ColumnDiff{
		OldColumn: *oldColumn,
		NewColumn: *newColumn,
	}
	if oldColumn.Name != newColumn.Name {
		return &diff
	}
	if oldColumn.Generated != newColumn.Generated {
		return &diff
	}
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
func tableDifferent(oldTable, newTable *Table) *TableDiff {
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
			columnDiff := columnDifferent(&column, foundColumn)
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

func CreateDiff(published, next *MetaContent) *Diff {
	diff := Diff{
		oldContent: published,
		newContent: next,
	}

	fmt.Println("进入 CreateDiff")
	publishedTables := published.Tables()
	nextTables := next.Tables()

	for _, table := range publishedTables {
		foundTable := FindTable(table.MetaUuid, nextTables)
		//删除的Table
		if foundTable == nil {
			diff.DeletedTables = append(diff.DeletedTables, table)
		}
	}
	for _, table := range nextTables {
		foundTable := FindTable(table.MetaUuid, publishedTables)
		//添加的Entity
		if foundTable == nil {
			diff.AddedTables = append(diff.AddedTables, table)
		} else {
			tableDiff := tableDifferent(table, foundTable)
			if tableDiff != nil {
				diff.ModifiedTables = append(diff.ModifiedTables, tableDiff)
			}
		}
	}

	return &diff
}

package model

type Table struct {
	MetaUuid string
	Name     string
	Columns  []ColumnMeta
}

type TableDelete struct {
	Table *Table
}

type TableAdd struct {
	Table *Table
}

type TableModify struct {
	OldTable *Table
	NewTable *Table
}
package meta

type Table struct {
	columns []*Column
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

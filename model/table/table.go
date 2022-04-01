package table

type Table struct {
	Uuid    string
	Name    string
	Columns []*Column
}

func FindTable(uuid string, tables []*Table) *Table {
	for i := range tables {
		if tables[i].Uuid == uuid {
			return tables[i]
		}
	}
	return nil
}

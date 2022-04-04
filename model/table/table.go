package table

type Table struct {
	Uuid          string
	Name          string
	EntityInnerId uint64
	Columns       []*Column
}

func FindTable(uuid string, tables []*Table) *Table {
	for i := range tables {
		if tables[i].Uuid == uuid {
			return tables[i]
		}
	}
	return nil
}

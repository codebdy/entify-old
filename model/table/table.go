package table

type Table struct {
	Uuid          string
	Name          string
	EntityInnerId uint64
	Columns       []*Column
	PKString      string
	Partial       bool
}

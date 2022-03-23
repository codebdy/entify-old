package meta

type Model struct {
	Entities  []*EntityMeta
	Relations []*RelationMeta
	Tables    []*Table
}

func NewModel(c *MetaContent) *Model {
	return nil
}

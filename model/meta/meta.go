package meta

type Model struct {
	Classes   []*ClassMeta
	Relations []*RelationMeta
}

func New(m *MetaContent) *Model {
	model := Model{
		Classes:   make([]*ClassMeta, len(m.Classes)),
		Relations: make([]*RelationMeta, len(m.Relations)),
	}

	for i := range m.Classes {
		model.Classes[i] = &m.Classes[i]
	}

	for i := range m.Relations {
		model.Relations[i] = &m.Relations[i]
	}
	return &model
}

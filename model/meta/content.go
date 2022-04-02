package meta

type MetaContent struct {
	Classes   []ClassMeta    `json:"entities"`
	Relations []RelationMeta `json:"relations"`
	Diagrams  []interface{}  `json:"diagrams"`
	X6Nodes   []interface{}  `json:"x6Nodes"`
	X6Edges   []interface{}  `json:"x6Edges"`
}

func (m *MetaContent) ToModel() *Model {
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

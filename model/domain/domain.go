package domain

import "rxdrag.com/entity-engine/model/meta"

type Model struct {
	Enums     []*Enum
	Classes   []*Class
	Relations []*Relation
}

func New(m *meta.Model) *Model {
	model := Model{}

	for i := range m.Classes {
		class := m.Classes[i]
		if class.StereoType == meta.CLASSS_ENUM {
			model.Enums[i] = NewEnum(class)
		} else {
			model.Classes[i] = NewClass(class)
		}
	}

	return &model
}

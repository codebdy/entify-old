package domain

/**
 * 在domain层，把有子类的实体，拆分成接口+实体
 * 比如A => A + AEntity
 */

import (
	"rxdrag.com/entity-engine/model/meta"
)

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

	for i := range m.Relations {
		relation := m.Relations[i]

		src := model.GetClassByUuid(relation.SourceId)
		tar := model.GetClassByUuid(relation.TargetId)
		if src == nil || tar == nil {
			panic("Meta is not integral, can not find class of relation:" + relation.Uuid)
		}
		if relation.RelationType == meta.INHERIT {
			src.Parents = append(src.Parents, tar)
			tar.Children = append(tar.Children, src)
		} else {
			r := NewRelation(relation, src, tar)
			model.Relations = append(model.Relations, r)
		}
	}
	return &model
}

func (m *Model) GetClassByUuid(uuid string) *Class {
	for i := range m.Classes {
		cls := m.Classes[i]
		if cls.Uuid == uuid {
			return cls
		}
	}

	return nil
}

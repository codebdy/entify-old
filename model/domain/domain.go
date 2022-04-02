package domain

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

		sourceClass := model.GetClassByUuid(relation.SourceId)
		targetClass := model.GetClassByUuid(relation.TargetId)
		if sourceClass == nil || targetClass == nil {
			panic("Meta is not integral, can not find class of relation:" + relation.Uuid)
		}
		if relation.RelationType == meta.INHERIT {
			sourceClass.Parents = append(sourceClass.Parents, targetClass)
			targetClass.Children = append(targetClass.Children, sourceClass)
		} else {
			model.Relations = append(model.Relations, NewRelation(relation, sourceClass, targetClass))
		}
	}

	//处理枚举属性

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

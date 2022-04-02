package domain

/**
 * 在domain层，把有子类的实体，拆分成接口+实体
 * 比如A => A + AEntity
 */

import (
	"rxdrag.com/entity-engine/consts"
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

	//把实体继承，拆分成接口+实体类
	newClases := []*Class{}
	for i := range model.Classes {
		cls := model.Classes[i]
		if cls.StereoType == meta.CLASSS_ENTITY && cls.HasChildren() {
			cls.StereoType = meta.CLASSS_ABSTRACT

			newCls := &Class{
				Uuid:        cls.Uuid + consts.ENTITY,
				StereoType:  meta.CLASSS_ENTITY,
				Name:        cls.Name + consts.ENTITY,
				Description: cls.Name + " entity class",
				Parents:     []*Class{cls},
			}

			cls.Children = append(cls.Children, newCls)
			newClases = append(newClases, newCls)
		}
	}

	model.Classes = append(model.Classes, newClases...)

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

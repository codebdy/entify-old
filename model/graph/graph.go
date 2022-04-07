package graph

import (
	"fmt"

	"rxdrag.com/entity-engine/model/domain"
	"rxdrag.com/entity-engine/model/meta"
	"rxdrag.com/entity-engine/model/table"
)

type Model struct {
	Enums      []*Enum
	Interfaces []*Interface
	Entities   []*Entity
	Relations  []*Relation
	Tables     []*table.Table
}

func New(m *domain.Model) *Model {
	model := Model{}

	for i := range m.Enums {
		model.Enums = append(model.Enums, NewEnum(m.Enums[i]))
	}

	//构建所有接口
	for i := range m.Classes {
		cls := m.Classes[i]
		if cls.StereoType == meta.CLASSS_ABSTRACT {
			model.Interfaces = append(model.Interfaces, NewInterface(cls))
		}
	}

	//构建所有实体
	for i := range m.Classes {
		cls := m.Classes[i]
		if cls.StereoType == meta.CLASSS_ENTITY ||
			cls.StereoType == meta.CLASS_VALUE_OBJECT ||
			cls.StereoType == meta.CLASS_SERVICE {
			newEntity := NewEntity(cls)
			model.Entities = append(model.Entities, newEntity)
			//构建接口实现关系
			allParents := cls.AllParents()
			for j := range allParents {
				parentInterface := model.GetInterfaceByUuid(allParents[j].Uuid)
				if parentInterface == nil {
					panic("Can not find interface by uuid:" + allParents[j].Uuid)
				}
				parentInterface.Children = append(parentInterface.Children, newEntity)
				newEntity.Interfaces = append(newEntity.Interfaces, parentInterface)
			}
		}
	}

	//处理关联
	for i := range m.Relations {
		relation := m.Relations[i]
		source := model.GetNodeByUuid(relation.Source.Uuid)
		if source.Entity() == nil && source.Interface() == nil {
			panic("Can not find souce by uuid:" + relation.Source.Uuid)
		}
		target := model.GetNodeByUuid(relation.Target.Uuid)
		if target.Entity() == nil && target.Interface() == nil {
			panic("Can not find target by uuid:" + relation.Target.Uuid)
		}
		r := NewRelation(relation, source, target)
		model.Relations = append(model.Relations, r)
		source.AddAssociation(NewAssociation(r, source.Uuid()))
		target.AddAssociation(NewAssociation(r, target.Uuid()))

		//增加派生关联
		sourceEntities := []*Entity{}
		targetEntities := []*Entity{}

		if source.isInterface() {
			sourceEntities = append(sourceEntities, source.Interface().Children...)
		} else {
			sourceEntities = append(sourceEntities, source.Entity())
		}

		if target.isInterface() {
			targetEntities = append(targetEntities, target.Interface().Children...)
		} else {
			targetEntities = append(targetEntities, target.Entity())
		}

		for i := range sourceEntities {
			s := sourceEntities[i]
			for j := range targetEntities {
				t := targetEntities[j]
				r.Children = append(r.Children, &DerivedRelation{
					Parent: r,
					Source: s,
					Target: t,
				})
			}
		}
	}

	//处理属性的实体类型跟枚举类型
	for i := range model.Interfaces {
		intf := model.Interfaces[i]
		for j := range intf.attributes {
			attr := intf.attributes[j]
			if attr.Type == meta.ENUM || attr.Type == meta.ENUM_ARRAY {
				attr.EumnType = model.GetEnumByUuid(attr.TypeUuid)
			}

			if attr.Type == meta.ENTITY || attr.Type == meta.ENTITY_ARRAY {
				attr.EnityType = model.GetEntityByUuid(attr.TypeUuid)
			}

			//这个代码并不成熟
			if attr.Type == meta.VALUE_OBJECT || attr.Type == meta.VALUE_OBJECT_ARRAY {
				attr.EnityType = model.GetEntityByUuid(attr.TypeUuid)
			}
		}
		for j := range intf.Methods {
			method := intf.Methods[j]
			if method.Method.Type == meta.ENUM || method.Method.Type == meta.ENUM_ARRAY {
				method.EumnType = model.GetEnumByUuid(method.Method.TypeUuid)
			}

			if method.Method.Type == meta.ENTITY || method.Method.Type == meta.ENTITY_ARRAY {
				method.EnityType = model.GetEntityByUuid(method.Method.TypeUuid)
			}

			//这个代码并不成熟
			if method.Method.Type == meta.VALUE_OBJECT || method.Method.Type == meta.VALUE_OBJECT_ARRAY {
				method.EnityType = model.GetEntityByUuid(method.Method.TypeUuid)
			}
		}

	}

	for i := range model.Entities {
		ent := model.Entities[i]
		for j := range ent.attributes {
			attr := ent.attributes[j]
			if attr.Type == meta.ENUM || attr.Type == meta.ENUM_ARRAY {
				attr.EumnType = model.GetEnumByUuid(attr.TypeUuid)
			}

			if attr.Type == meta.ENTITY || attr.Type == meta.ENTITY_ARRAY {
				attr.EnityType = model.GetEntityByUuid(attr.TypeUuid)
			}

			//这个代码并不成熟
			if attr.Type == meta.VALUE_OBJECT || attr.Type == meta.VALUE_OBJECT_ARRAY {
				attr.EnityType = model.GetEntityByUuid(attr.TypeUuid)
			}
		}
		for j := range ent.Methods {
			method := ent.Methods[j]
			if method.Method.Type == meta.ENUM || method.Method.Type == meta.ENUM_ARRAY {
				method.EumnType = model.GetEnumByUuid(method.Method.TypeUuid)
			}

			if method.Method.Type == meta.ENTITY || method.Method.Type == meta.ENTITY_ARRAY {
				method.EnityType = model.GetEntityByUuid(method.Method.TypeUuid)
			}

			//这个代码并不成熟
			if method.Method.Type == meta.VALUE_OBJECT || method.Method.Type == meta.VALUE_OBJECT_ARRAY {
				method.EnityType = model.GetEntityByUuid(method.Method.TypeUuid)
			}
		}
	}
	//处理Table
	for i := range model.Entities {
		ent := model.Entities[i]
		if ent.Domain.StereoType == meta.CLASSS_ENTITY {
			model.Tables = append(model.Tables, NewEntityTable(ent))
		}
	}

	for i := range model.Relations {
		relation := model.Relations[i]
		model.Tables = append(model.Tables, NewRelationTables(relation)...)
	}

	return &model
}

func (m *Model) Validate() {
	//检查空实体（除ID外没有属性跟关联）
	for _, entity := range m.Entities {
		if entity.IsEmperty() {
			panic(fmt.Sprintf("Entity %s should have one normal field at least", entity.Name()))
		}
	}
}

func (m *Model) RootEnities() []*Entity {
	entities := []*Entity{}
	for i := range m.Entities {
		ent := m.Entities[i]
		if ent.Domain.Root {
			entities = append(entities, ent)
		}
	}

	return entities
}

func (m *Model) RootInterfaces() []*Interface {
	interfaces := []*Interface{}
	for i := range m.Interfaces {
		intf := m.Interfaces[i]
		if intf.Domain.Root {
			interfaces = append(interfaces, intf)
		}
	}

	return interfaces
}

func (m *Model) GetNodeByUuid(uuid string) Node {
	intf := m.GetInterfaceByUuid(uuid)

	if intf != nil {
		return intf
	}

	return m.GetEntityByUuid(uuid)
}

func (m *Model) GetNodeByName(name string) Node {
	intf := m.GetInterfaceByName(name)

	if intf != nil {
		return intf
	}

	return m.GetEntityByName(name)
}

func (m *Model) GetInterfaceByUuid(uuid string) *Interface {
	for i := range m.Interfaces {
		intf := m.Interfaces[i]
		if intf.Uuid() == uuid {
			return intf
		}
	}
	return nil
}

func (m *Model) GetEntityByUuid(uuid string) *Entity {
	for i := range m.Entities {
		ent := m.Entities[i]
		if ent.Uuid() == uuid {
			return ent
		}
	}
	return nil
}

func (m *Model) GetInterfaceByName(name string) *Interface {
	for i := range m.Interfaces {
		intf := m.Interfaces[i]
		if intf.Name() == name {
			return intf
		}
	}
	return nil
}

func (m *Model) GetEntityByName(name string) *Entity {
	for i := range m.Entities {
		ent := m.Entities[i]
		if ent.Name() == name {
			return ent
		}
	}
	return nil
}

func (m *Model) GetMetaEntity() *Entity {
	return m.GetEntityByUuid(meta.MetaClass.Uuid)
}

func (m *Model) GetEnumByUuid(uuid string) *Enum {
	for i := range m.Enums {
		enum := m.Enums[i]
		if enum.Uuid == uuid {
			return enum
		}
	}
	return nil
}

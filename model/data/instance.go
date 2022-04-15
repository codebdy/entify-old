package data

import (
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/model/table"
)

type Field struct {
	Column *table.Column
	Value  interface{}
}

type Instance struct {
	Id                uint64
	Entity            *graph.Entity
	Fields            []*Field
	References        []*Reference
	DerivedReferences []*DerivedReference
}

func NewInstance(object map[string]interface{}, entity *graph.Entity) *Instance {
	instance := Instance{
		Entity: entity,
	}
	if object[consts.ID] != nil {
		instance.Id = object[consts.ID].(uint64)
	}

	columns := entity.Table.Columns
	for i := range columns {
		column := columns[i]
		if object[column.Name] != nil {
			instance.Fields = append(instance.Fields, &Field{
				Column: column,
				Value:  object[column.Name],
			})
		}
	}
	allAssociation := entity.AllAssociations()
	for i := range allAssociation {
		asso := allAssociation[i]
		if !asso.IsAbstract() {
			if object[asso.Name()] != nil {
				instance.References = append(instance.References, &Reference{
					Association: asso,
					Value:       object[asso.Name()].(map[string]interface{}),
				})
			}

		} else {

		}
	}
	return &instance
}

func (ins *Instance) IsInsert() bool {
	for i := range ins.Fields {
		field := ins.Fields[i]
		if field.Column.Name == consts.ID {
			if field.Value != nil {
				return false
			}
		}
	}
	return true
}

func (ins *Instance) Table() *table.Table {
	return ins.Entity.Table
}

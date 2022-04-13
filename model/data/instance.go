package data

import "rxdrag.com/entity-engine/model/graph"

type Field struct {
	Attribute *graph.Attribute
	Value     interface{}
}

type Instance struct {
	Entity     *graph.Entity
	Fields     []*Field
	References []*Reference
}

func New(object map[string]interface{}, entity *graph.Entity) *Instance {
	instance := Instance{
		Entity: entity,
	}
	allAttributes := entity.AllAttributes()
	for i := range allAttributes {
		attr := allAttributes[i]
		if object[attr.Name] != nil {
			instance.Fields = append(instance.Fields, &Field{
				Attribute: attr,
				Value:     object[attr.Name],
			})
		}
	}
	return &instance
}

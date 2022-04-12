package data

import "rxdrag.com/entity-engine/model/graph"

type Field struct {
	Attribute *graph.Attribute
	Value     interface{}
}

type Fields = map[string]Field

type Instance struct {
	Entity     *graph.Entity
	Fields     Fields
	References []Reference
}

func New(object map[string]interface{}, entity *graph.Entity) *Instance {
	instance := Instance{
		Entity: entity,
	}

	return &instance
}

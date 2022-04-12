package data

import "rxdrag.com/entity-engine/model/graph"

type Attributes = map[string]interface{}

type Instance struct {
	Entity       *graph.Entity
	Atrributes   Attributes
	Associations []Association
}

func New(object map[string]interface{}, entity *graph.Entity) *Instance {
	instance := Instance{
		Entity: entity,
	}

	return &instance
}

package data

import "rxdrag.com/entity-engine/model/graph"

type Fields = map[string]interface{}

type Instance struct {
	Entity *graph.Entity
	Fields Fields
	Edges  []Reference
}

func New(object map[string]interface{}, entity *graph.Entity) *Instance {
	instance := Instance{
		Entity: entity,
	}

	return &instance
}

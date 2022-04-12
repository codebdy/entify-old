package data

import "rxdrag.com/entity-engine/model/graph"

type Attributes = map[string]interface{}

type Node struct {
	Entity       *graph.Entity
	Atrributes   Attributes
	Associations []Edge
}

func New(object map[string]interface{}, entity *graph.Entity) *Node {
	instance := Node{
		Entity: entity,
	}

	return &instance
}

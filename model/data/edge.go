package data

import "rxdrag.com/entity-engine/model/graph"

type Edge struct {
	Association *graph.Association
	Single      Node
	Array       []Node
}

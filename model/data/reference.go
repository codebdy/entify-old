package data

import "rxdrag.com/entity-engine/model/graph"

type Reference struct {
	Association *graph.Association
	Single      Instance
	Array       []Instance
}

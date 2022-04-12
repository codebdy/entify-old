package data

import "rxdrag.com/entity-engine/model/graph"

type Reference struct {
	Association *graph.Association
}

type HasOne struct {
	Reference
	data Instance
}

type HasMany struct {
	Reference
	data []Instance
}

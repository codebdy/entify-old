package data

import "rxdrag.com/entity-engine/model/graph"

type HasOne struct {
	Cascade bool
	Add     Instance
	Delete  Instance
	Update  Instance
	Sync    Instance
}

type HasMany struct {
	Cascade bool
	Add     []Instance
	Delete  []Instance
	Update  []Instance
	Sync    []Instance
}

type Reference struct {
	Association *graph.Association
	Value       interface{}
}

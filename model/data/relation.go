package data

import (
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model/graph"
)

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
	Value       map[string]interface{}
}

func (r *Reference) Deleted() []*Instance {
	instances := []*Instance{}

	return instances
}

func (r *Reference) Added() []*Instance {
	instances := []*Instance{}

	return instances
}

func (r *Reference) updated() []*Instance {
	instances := []*Instance{}

	return instances
}

func (r *Reference) Cascade() bool {
	return r.Value[consts.ARG_CASCADE].(bool)
}

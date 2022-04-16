package data

import (
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/model/table"
)

// type HasOne struct {
// 	Cascade bool
// 	Add     Instance
// 	Delete  Instance
// 	Update  Instance
// 	Sync    Instance
// }

// type HasMany struct {
// 	Cascade bool
// 	Add     []Instance
// 	Delete  []Instance
// 	Update  []Instance
// 	Sync    []Instance
// }

type Associationer interface {
	Deleted() []*Instance
	Added() []*Instance
	Updated() []*Instance
	Synced() []*Instance
	Cascade() bool
	SourceColumn() *table.Column
	TargetColumn() *table.Column
	Table() *table.Table
	IsSource() bool
}

type Reference struct {
	Association *graph.Association
	Value       map[string]interface{}
}

type DerivedReference struct {
	Association *graph.DerivedAssociation
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

func (r *Reference) Updated() []*Instance {
	instances := []*Instance{}

	return instances
}

func (r *Reference) Synced() []*Instance {
	instances := []*Instance{}

	return instances
}

func (r *Reference) Cascade() bool {
	return r.Value[consts.ARG_CASCADE].(bool)
}

func (r *Reference) SourceColumn() *table.Column {
	for i := range r.Association.Relation.Table.Columns {
		column := r.Association.Relation.Table.Columns[i]
		if column.Name == r.Association.Relation.Source.TableName() {
			return column
		}
	}
	return nil
}

func (r *Reference) TargetColumn() *table.Column {
	for i := range r.Association.Relation.Table.Columns {
		column := r.Association.Relation.Table.Columns[i]
		if column.Name == r.Association.Relation.Target.TableName() {
			return column
		}
	}
	return nil
}

func (r *Reference) Table() *table.Table {
	return r.Association.Relation.Table
}

func (r *Reference) IsSource() bool {
	return r.IsSource()
}

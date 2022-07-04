package data

import (
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model/graph"
	"rxdrag.com/entify/model/meta"
	"rxdrag.com/entify/model/table"
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
	OwnerColumn() *table.Column
	TypeColumn() *table.Column
	//Entity, Partial and External
	TypeEntity() *graph.Entity
	IsCombination() bool
}

//没有继承关系的关联
type Reference struct {
	Association *graph.Association
	Value       map[string]interface{}
}

type DerivedReference struct {
	Association *graph.DerivedAssociation
	Value       map[string]interface{}
}

func doConvertToInstances(data interface{}, isArray bool, entity *graph.Entity) []*Instance {
	instances := []*Instance{}
	if data == nil {
		return []*Instance{}
	}
	if isArray {
		objects := data.([]interface{})
		for i := range objects {
			instances = append(instances, NewInstance(objects[i].(map[string]interface{}), entity))
		}
	} else {
		instances = append(instances, NewInstance(data.(map[string]interface{}), entity))
	}

	return instances
}

func (r *Reference) convertToInstances(data interface{}) []*Instance {
	return doConvertToInstances(data, r.Association.IsArray(), r.TypeEntity())
}

func (r *DerivedReference) convertToInstances(data interface{}) []*Instance {
	return doConvertToInstances(data, r.Association.DerivedFrom.IsArray(), r.TypeEntity())
}

func (r *Reference) Deleted() []*Instance {
	return r.convertToInstances(r.Value[consts.ARG_DELETE])
}

func (r *Reference) Added() []*Instance {
	return r.convertToInstances(r.Value[consts.ARG_ADD])
}

func (r *Reference) Updated() []*Instance {
	return r.convertToInstances(r.Value[consts.ARG_UPDATE])
}

func (r *Reference) Synced() []*Instance {
	return r.convertToInstances(r.Value[consts.ARG_SYNC])
}

func (r *Reference) Cascade() bool {
	if r.Value[consts.ARG_CASCADE] != nil {
		return r.Value[consts.ARG_CASCADE].(bool)
	}
	return false
}

func (r *Reference) SourceColumn() *table.Column {
	for i := range r.Association.Relation.Table.Columns {
		column := r.Association.Relation.Table.Columns[i]
		if column.Name == r.Association.Relation.SourceClass().TableName() {
			return column
		}
	}
	return nil
}

func (r *Reference) TargetColumn() *table.Column {
	for i := range r.Association.Relation.Table.Columns {
		column := r.Association.Relation.Table.Columns[i]
		if column.Name == r.Association.Relation.TargetClass().TableName() {
			return column
		}
	}
	return nil
}

func (r *Reference) Table() *table.Table {
	return r.Association.Relation.Table
}

func (r *Reference) IsSource() bool {
	return r.Association.IsSource()
}

func (r *Reference) OwnerColumn() *table.Column {
	if r.IsSource() {
		return r.SourceColumn()
	} else {
		return r.TargetColumn()
	}
}
func (r *Reference) TypeColumn() *table.Column {
	if !r.IsSource() {
		return r.SourceColumn()
	} else {
		return r.TargetColumn()
	}
}

func (r *Reference) TypeEntity() *graph.Entity {
	entity := r.Association.TypeEntity()
	if entity != nil {
		return entity
	}

	partial := r.Association.TypePartial()

	if partial != nil {
		return &partial.Entity
	}

	external := r.Association.TypeExternal()

	if external != nil {
		return &external.Entity
	}
	panic("Can not find reference entity")
}

func (r *Reference) IsCombination() bool {
	return r.IsSource() &&
		(r.Association.Relation.RelationType == meta.TWO_WAY_COMBINATION ||
			r.Association.Relation.RelationType == meta.ONE_WAY_COMBINATION)
}

//====derived
func (r *DerivedReference) Deleted() []*Instance {
	return r.convertToInstances(r.Value[consts.ARG_DELETE])
}

func (r *DerivedReference) Added() []*Instance {
	return r.convertToInstances(r.Value[consts.ARG_ADD])
}

func (r *DerivedReference) Updated() []*Instance {
	return r.convertToInstances(r.Value[consts.ARG_UPDATE])
}

func (r *DerivedReference) Synced() []*Instance {
	return r.convertToInstances(r.Value[consts.ARG_SYNC])
}

func (r *DerivedReference) Cascade() bool {
	return r.Value[consts.ARG_CASCADE].(bool)
}

func (r *DerivedReference) SourceColumn() *table.Column {
	for i := range r.Association.Relation.Table.Columns {
		column := r.Association.Relation.Table.Columns[i]
		if column.Name == r.Association.Relation.SourceClass().TableName() {
			return column
		}
	}
	return nil
}

func (r *DerivedReference) TargetColumn() *table.Column {
	for i := range r.Association.Relation.Table.Columns {
		column := r.Association.Relation.Table.Columns[i]
		if column.Name == r.Association.Relation.TargetClass().TableName() {
			return column
		}
	}
	return nil
}

func (r *DerivedReference) Table() *table.Table {
	return r.Association.Relation.Table
}

func (r *DerivedReference) IsSource() bool {
	return r.Association.DerivedFrom.IsSource()
}

func (r *DerivedReference) OwnerColumn() *table.Column {
	if r.IsSource() {
		return r.SourceColumn()
	} else {
		return r.TargetColumn()
	}
}
func (r *DerivedReference) TypeColumn() *table.Column {
	if !r.IsSource() {
		return r.SourceColumn()
	} else {
		return r.TargetColumn()
	}
}

func (r *DerivedReference) TypeEntity() *graph.Entity {
	entity := r.Association.TypeEntity()
	if entity != nil {
		return entity
	}

	partial := r.Association.TypePartial()

	if partial != nil {
		return &partial.Entity
	}

	external := r.Association.TypeExternal()

	if external != nil {
		return &external.Entity
	}
	panic("Can not find reference entity")
}

func (r *DerivedReference) IsCombination() bool {
	return r.IsSource() &&
		(r.Association.Relation.Parent.RelationType == meta.TWO_WAY_COMBINATION ||
			r.Association.Relation.Parent.RelationType == meta.ONE_WAY_COMBINATION)
}

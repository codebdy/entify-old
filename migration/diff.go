package migration

import "rxdrag.com/entity-engine/meta"

type ColumnDiff struct {
	OldColumn meta.Column
	NewColumn meta.Column
}

type EntityDiff struct {
	DeleteColumns []meta.Column
	AddColumns    []meta.Column
	ModifyColumns []ColumnDiff
}

type RelationDiff struct {
	OldeRelation meta.Relation
	NewRelation  meta.Relation
}

type Diff struct {
	DeleteEntities  []meta.Entity
	DeleteRelations []meta.Relation
	AddEntities     []meta.Entity
	AddRlations     []meta.Relation
	ModifyEntities  []EntityDiff
	ModifyRelations []RelationDiff
}

package meta

type ColumnDiff struct {
	OldColumn Column
	NewColumn Column
}

type EntityDiff struct {
	DeleteColumns []Column
	AddColumns    []Column
	ModifyColumns []ColumnDiff
}

type RelationDiff struct {
	OldeRelation Relation
	NewRelation  Relation
}

type Diff struct {
	DeleteRelations []Relation
	DeleteEntities  []Entity
	AddEntities     []Entity
	AddRlations     []Relation
	ModifyEntities  []EntityDiff
	ModifyRelations []RelationDiff
}

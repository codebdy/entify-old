package meta

type ColumnDiff struct {
	OldColumn Column
	NewColumn Column
}

type EntityDiff struct {
	OldEntity     *Entity
	NewEntity     *Entity
	DeleteColumns []Column
	AddColumns    []Column
	ModifyColumns []ColumnDiff
}

type RelationDiff struct {
	OldeRelation Relation
	NewRelation  Relation
}

type Diff struct {
	oldContent *MetaContent
	newContent *MetaContent

	DeletedRelations []Relation
	DeletedEntities  []Entity
	AddedEntities    []Entity
	AddedRlations    []Relation
	ModifiedEntities []EntityDiff
	ModifieRelations []RelationDiff
}

func findRelation(uuid string, relations []Relation) *Relation {
	for _, relation := range relations {
		if relation.Uuid == uuid {
			return &relation
		}
	}

	return nil
}

func findEntity(uuid string, entities []Entity) *Entity {
	for _, entity := range entities {
		if entity.Uuid == uuid {
			return &entity
		}
	}

	return nil
}

func findColumn(uuid string, columns []Column) *Column {
	for _, column := range columns {
		if column.Uuid == uuid {
			return &column
		}
	}

	return nil
}

func relationDifferent(oldRelation, newRelation *Relation) *RelationDiff {
	diff := RelationDiff{
		OldeRelation: *oldRelation,
		NewRelation:  *newRelation,
	}
	if oldRelation.RelationType != newRelation.RelationType {
		return &diff
	}
	if oldRelation.RoleOnSource != newRelation.RoleOnSource {
		return &diff
	}
	if oldRelation.RoleOnTarget != newRelation.RoleOnTarget {
		return &diff
	}
	if oldRelation.SourceId != newRelation.SourceId {
		return &diff
	}
	if oldRelation.TargetId != newRelation.TargetId {
		return &diff
	}
	return nil
}

func columnDifferent(oldColumn, newColumn *Column) *ColumnDiff {
	diff := ColumnDiff{
		OldColumn: *oldColumn,
		NewColumn: *newColumn,
	}
	if oldColumn.Name != newColumn.Name {
		return &diff
	}
	if oldColumn.Generated != newColumn.Generated {
		return &diff
	}
	if oldColumn.Index != newColumn.Index {
		return &diff
	}
	if oldColumn.Nullable != newColumn.Nullable {
		return &diff
	}
	if oldColumn.Length != newColumn.Length {
		return &diff
	}
	if oldColumn.Primary != newColumn.Primary {
		return &diff
	}

	if oldColumn.Unique != newColumn.Unique {
		return &diff
	}

	if oldColumn.Type != newColumn.Type {
		return &diff
	}
	return nil
}
func entityDifferent(oldEntity, newEntity *Entity) *EntityDiff {
	var diff EntityDiff
	modified := false
	diff.OldEntity = oldEntity
	diff.NewEntity = newEntity

	for _, column := range oldEntity.Columns {
		foundCoumn := findColumn(column.Uuid, newEntity.Columns)
		if foundCoumn == nil {
			diff.DeleteColumns = append(diff.DeleteColumns, column)
			modified = true
		}
	}

	for _, column := range newEntity.Columns {
		foundColumn := findColumn(column.Uuid, oldEntity.Columns)
		if foundColumn == nil {
			diff.AddColumns = append(diff.AddColumns, column)
			modified = true
		} else {
			columnDiff := columnDifferent(&column, foundColumn)
			if columnDiff != nil {
				diff.ModifyColumns = append(diff.ModifyColumns, *columnDiff)
				modified = true
			}
		}
	}

	if diff.OldEntity.GetTableName() != diff.NewEntity.GetTableName() ||
		diff.OldEntity.EntityType != diff.NewEntity.EntityType ||
		modified {
		return &diff
	}
	return nil
}

func CreateDiff(published, next *MetaContent) *Diff {
	diff := Diff{
		oldContent: published,
		newContent: next,
	}
	publishedRelations := published.Relations
	nextRelations := next.Relations

	for _, relation := range publishedRelations {
		foundRelation := findRelation(relation.Uuid, nextRelations)
		//删除的Relation
		if foundRelation == nil {
			diff.DeletedRelations = append(diff.DeletedRelations, relation)
		}
	}
	for _, relation := range nextRelations {
		foundRelation := findRelation(relation.Uuid, publishedRelations)
		//添加的Relation
		if foundRelation == nil {
			diff.AddedRlations = append(diff.AddedRlations, relation)
		} else {
			relationDiff := relationDifferent(&relation, foundRelation)
			if relationDiff != nil {
				diff.ModifieRelations = append(diff.ModifieRelations, *relationDiff)
			}
		}
	}

	publishedEntities := published.Entities
	nextEntities := next.Entities

	for _, entity := range publishedEntities {
		foundEntity := findEntity(entity.Uuid, nextEntities)
		//删除的Entity
		if foundEntity == nil {
			diff.DeletedEntities = append(diff.DeletedEntities, entity)
		}
	}
	for _, entity := range nextEntities {
		foundEntity := findEntity(entity.Uuid, publishedEntities)
		//添加的Entity
		if foundEntity == nil {
			diff.AddedEntities = append(diff.AddedEntities, entity)
		} else {
			entityDiff := entityDifferent(&entity, foundEntity)
			if entityDiff != nil {
				diff.ModifiedEntities = append(diff.ModifiedEntities, *entityDiff)
			}
		}
	}

	return &diff
}

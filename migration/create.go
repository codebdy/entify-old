package migration

import (
	"rxdrag.com/entity-engine/meta"
)

func findRelation(uuid string, relations []meta.Relation) *meta.Relation {
	for _, relation := range relations {
		if relation.Uuid == uuid {
			return &relation
		}
	}

	return nil
}

func findEntity(uuid string, entities []meta.Entity) *meta.Entity {
	for _, entity := range entities {
		if entity.Uuid == uuid {
			return &entity
		}
	}

	return nil
}

func findColumn(uuid string, columns []meta.Column) *meta.Column {
	for _, column := range columns {
		if column.Uuid == uuid {
			return &column
		}
	}

	return nil
}

func relationDifferent(oldRelation, newRelation *meta.Relation) *meta.RelationDiff {
	diff := meta.RelationDiff{
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

func columnDifferent(oldColumn, newColumn *meta.Column) *meta.ColumnDiff {
	diff := meta.ColumnDiff{
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
func entityDifferent(oldEntity, newEntity *meta.Entity) *meta.EntityDiff {
	var diff meta.EntityDiff
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

func CreateDiff(published, next *meta.MetaContent) *meta.Diff {
	var diff meta.Diff
	publishedRelations := published.Relations
	nextRelations := next.Relations

	for _, relation := range publishedRelations {
		foundRelation := findRelation(relation.Uuid, nextRelations)
		//删除的Relation
		if foundRelation == nil {
			diff.DeleteRelations = append(diff.DeleteRelations, relation)
		}
	}
	for _, relation := range nextRelations {
		foundRelation := findRelation(relation.Uuid, publishedRelations)
		//添加的Relation
		if foundRelation == nil {
			diff.AddRlations = append(diff.AddRlations, relation)
		} else {
			relationDiff := relationDifferent(&relation, foundRelation)
			if relationDiff != nil {
				diff.ModifyRelations = append(diff.ModifyRelations, *relationDiff)
			}
		}
	}

	publishedEntities := published.Entities
	nextEntities := next.Entities

	for _, entity := range publishedEntities {
		foundEntity := findEntity(entity.Uuid, nextEntities)
		//删除的Entity
		if foundEntity == nil {
			diff.DeleteEntities = append(diff.DeleteEntities, entity)
		}
	}
	for _, entity := range nextEntities {
		foundEntity := findEntity(entity.Uuid, publishedEntities)
		//添加的Entity
		if foundEntity == nil {
			diff.AddEntities = append(diff.AddEntities, entity)
		} else {
			entityDiff := entityDifferent(&entity, foundEntity)
			if entityDiff != nil {
				diff.ModifyEntities = append(diff.ModifyEntities, *entityDiff)
			}
		}
	}

	return &diff
}

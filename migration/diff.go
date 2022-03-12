package migration

import (
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/utils"
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

func relations(object utils.Object) []meta.Relation {
	var relations []meta.Relation
	if object[consts.META_RELATIONS] != nil {
		relations = object[consts.META_RELATIONS].([]meta.Relation)
	} else {
		relations = make([]meta.Relation, 0)
	}
	return relations
}

func entities(object utils.Object) []meta.Entity {
	var entities []meta.Entity
	if object[consts.META_ENTITIES] != nil {
		entities = object[consts.META_ENTITIES].([]meta.Entity)
	} else {
		entities = make([]meta.Entity, 0)
	}
	return entities
}

func relationDifferent(oldRelation, newRelation *meta.Relation) *meta.RelationDiff {
	diff := meta.RelationDiff{
		OldeRelation: oldRelation,
		NewRelation:  newRelation,
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

func entityDifferent(oldRelation, newRelation *meta.Entity) *meta.EntityDiff {
	var diff meta.EntityDiff
	return &diff
}

func CreateDiff(published, next utils.Object) *meta.Diff {
	var diff meta.Diff
	publishedRelations := relations(published)
	nextRelations := relations(next)

	for _, relation := range publishedRelations {
		foundRelation := findRelation(relation.Uuid, nextRelations)
		//删除的Relation
		if foundRelation == nil {
			diff.DeleteRelations = append(diff.DeleteRelations, &relation)
		}
	}
	for _, relation := range nextRelations {
		foundRelation := findRelation(relation.Uuid, publishedRelations)
		//添加的Relation
		if foundRelation == nil {
			diff.AddRlations = append(diff.AddRlations, &relation)
		} else {
			relationDiff := relationDifferent(&relation, foundRelation)
			if relationDiff != nil {
				diff.ModifyRelations = append(diff.ModifyRelations, relationDiff)
			}
		}
	}

	publishedEntities := entities(published)
	nextEntities := entities(next)

	for _, entity := range publishedEntities {
		foundEntity := findEntity(entity.Uuid, nextEntities)
		//删除的Entity
		if foundEntity == nil {
			diff.DeleteEntities = append(diff.DeleteEntities, &entity)
		}
	}
	for _, entity := range nextEntities {
		foundEntity := findEntity(entity.Uuid, publishedEntities)
		//添加的Entity
		if foundEntity == nil {
			diff.AddEntities = append(diff.AddEntities, &entity)
		} else {
			entityDiff := entityDifferent(&entity, foundEntity)
			if entityDiff != nil {
				diff.ModifyEntities = append(diff.ModifyEntities, entityDiff)
			}
		}
	}

	return &diff
}

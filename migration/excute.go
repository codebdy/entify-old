package migration

import (
	"fmt"

	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/repository"
)

func ExcuteDiff(d *meta.Diff) {
	fmt.Println("ExcuteDiff:", d)
	for _, relation := range d.DeleteRelations {
		repository.DeleteRelation(relation)
	}
	for _, entity := range d.DeleteEntities {
		repository.DeleteEntity(entity.Name)
	}

	for _, entity := range d.AddEntities {
		repository.AddEntity(entity)
	}

	for _, relation := range d.AddRlations {
		repository.AddRelation(relation)
	}

	for _, entityDiff := range d.ModifyEntities {
		repository.ModifyEntity(entityDiff)
	}

	for _, relationDiff := range d.ModifyRelations {
		repository.ModifyRelation(relationDiff)
	}
}

func UndoDiff(d *meta.Diff) {

}

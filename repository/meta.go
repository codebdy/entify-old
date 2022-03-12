package repository

import (
	"fmt"

	"rxdrag.com/entity-engine/meta"
)

func DeleteRelation(relatin *meta.Relation) {
	fmt.Println("Not implement DeleteRelation")
}

func DeleteEntity(entityName string) {
	fmt.Println("Not implement DeleteEntity")
}

func AddEntity(entity *meta.Entity) {
	fmt.Println("Not implement AddEntity")
}

func AddRelation(relation *meta.Relation) {
	fmt.Println("Not implement AddRelation")
}
func ModifyEntity(entityDiff *meta.EntityDiff) {
	fmt.Println("Not implement ModifyEntity")
}

func ModifyRelation(relationDiff *meta.RelationDiff) {
	fmt.Println("Not implement ModifyRelation")
}

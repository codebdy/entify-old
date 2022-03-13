package repository

import (
	"database/sql"
	"fmt"

	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/repository/dialect"
)

func DeleteRelation(relatin *meta.Relation) {
	fmt.Println("Not implement DeleteRelation")
}

func DeleteEntity(entityName string) {
	fmt.Println("Not implement DeleteEntity")
}

func AddEntity(entity *meta.Entity, undoList *[]string, db *sql.DB) {
	sqlBuilder := dialect.GetSQLBuilder()
	excuteSQL, undoSQL := sqlBuilder.BuildCreateEntitySQL(entity)
	*undoList = append(*undoList, undoSQL)
	db.Exec(excuteSQL)
	fmt.Println("AddEntity SQL:", excuteSQL)
}

func AddRelation(relation *meta.Relation) {
	fmt.Println("Not implement AddRelation", relation.RoleOnSource, relation.RoleOnTarget)
}
func ModifyEntity(entityDiff *meta.EntityDiff) {
	fmt.Println("Not implement ModifyEntity")
}

func ModifyRelation(relationDiff *meta.RelationDiff) {
	fmt.Println("Not implement ModifyRelation")
}

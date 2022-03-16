package migration

import (
	"database/sql"
	"fmt"

	"rxdrag.com/entity-engine/config"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/repository"
)

func ExcuteDiff(d *meta.Diff) {
	var undoList []string
	db, err := sql.Open(config.DRIVER_NAME, config.MYSQL_CONFIG)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, relation := range d.DeletedRelations {
		repository.DeleteRelation(&relation)
	}
	for _, entity := range d.DeletedEntities {
		repository.DeleteEntity(entity.Name)
	}

	for _, entity := range d.AddedEntities {
		repository.AddEntity(&entity, &undoList, db)
	}

	for _, relation := range d.AddedRlations {
		repository.AddRelation(&relation)
	}

	for _, entityDiff := range d.ModifiedEntities {
		repository.ModifyEntity(&entityDiff)
	}

	for _, relationDiff := range d.ModifieRelations {
		repository.ModifyRelation(&relationDiff)
	}
}

func UndoDiff(d *meta.Diff) {

}

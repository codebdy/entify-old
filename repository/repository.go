package repository

import (
	"database/sql"
	"fmt"

	"rxdrag.com/entity-engine/config"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/utils"
)

func SaveOneEntity(object utils.SimpleJSON, entity *meta.Entity) (interface{}, error) {
	fmt.Println(object)
	db, err := sql.Open("mysql", config.MYSQL_CONFIG)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return nil, nil
	//fmt.Println(p.Args)
	//fmt.Println(p.Conte
}

package repository

import (
	"database/sql"
	"fmt"
	"log"

	"rxdrag.com/entity-engine/config"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/repository/dialect"
	"rxdrag.com/entity-engine/utils"
)

func clearTransaction(tx *sql.Tx) {
	err := tx.Rollback()
	if err != sql.ErrTxDone && err != nil {
		log.Fatalln(err)
	}
}

func InsertOne(object map[string]interface{}, entity *meta.Entity) (interface{}, error) {
	db, err := sql.Open(config.DRIVER_NAME, config.MYSQL_CONFIG)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	tx, err := db.Begin()
	defer clearTransaction(tx)

	if err != nil {
		return nil, err
	}
	sqlBuilder := dialect.GetSQLBuilder()
	saveStr, values := sqlBuilder.BuildInsertSQL(object, entity)

	_, err = tx.Exec(saveStr, values...)
	if err != nil {
		fmt.Println("Insert data failed:", err.Error())
		return nil, err
	}
	tx.Commit()

	id := object[consts.META_ID]
	savedObject, err := QueryOneById(entity, id)
	if err != nil {
		fmt.Println("QueryOneById failed:", err.Error())
		return nil, err
	}
	//affectedRows, err := result.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected failed:", err.Error())
		return nil, err
	}

	return savedObject, nil
}

func UpdateOne(object map[string]interface{}, entity *meta.Entity) (interface{}, error) {
	db, err := sql.Open(config.DRIVER_NAME, config.MYSQL_CONFIG)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	sqlBuilder := dialect.GetSQLBuilder()

	tx, err := db.Begin()
	defer clearTransaction(tx)

	saveStr, values := sqlBuilder.BuildUpdateSQL(object, entity)
	if err != nil {
		return nil, err
	}
	fmt.Println(saveStr)
	_, err = tx.Exec(saveStr, values...)
	if err != nil {
		fmt.Println("Update data failed:", err.Error())
		return nil, err
	}
	tx.Commit()

	id := object[consts.META_ID]

	savedObject, err := QueryOneById(entity, id)
	if err != nil {
		fmt.Println("QueryOneById failed:", err.Error())
		return nil, err
	}
	return savedObject, nil
}

func SaveOne(object map[string]interface{}, entity *meta.Entity) (interface{}, error) {
	if object[consts.META_ID] == nil {
		object[consts.META_ID] = utils.CreateId()
		return InsertOne(object, entity)
	} else {
		return UpdateOne(object, entity)
	}
}

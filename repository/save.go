package repository

import (
	"database/sql"
	"fmt"

	"rxdrag.com/entity-engine/config"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/repository/dialect"
	"rxdrag.com/entity-engine/utils"
)

func SaveOne(object map[string]interface{}, entity *model.Entity) (interface{}, error) {
	db, err := sql.Open(config.DRIVER_NAME, config.MYSQL_CONFIG)
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	tx, err := NewTx(db)
	defer tx.clear()
	if err != nil {
		panic(err.Error())
	}

	obj, err := tx.doSaveOne(object, entity)
	if err != nil {
		panic(err.Error())
	}
	err = tx.Commit()
	if err != nil {
		panic(err.Error())
	}
	return obj, nil
}

func InsertOne(object map[string]interface{}, entity *model.Entity) (interface{}, error) {
	db, err := sql.Open(config.DRIVER_NAME, config.MYSQL_CONFIG)
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	tx, err := NewTx(db)
	defer tx.clear()
	if err != nil {
		panic(err.Error())
	}

	obj, err := tx.doInsertOne(object, entity)
	if err != nil {
		panic(err.Error())
	}
	err = tx.Commit()
	if err != nil {
		panic(err.Error())
	}
	return obj, nil
}

func (tx *Tx) doInsertOne(object map[string]interface{}, entity *model.Entity) (interface{}, error) {
	sqlBuilder := dialect.GetSQLBuilder()
	saveStr, values := sqlBuilder.BuildInsertSQL(object, entity)

	_, err := tx.Exec(saveStr, values...)
	if err != nil {
		fmt.Println("Insert data failed:", err.Error())
		return nil, err
	}

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

func (tx *Tx) doUpdateOne(object map[string]interface{}, entity *model.Entity) (interface{}, error) {

	sqlBuilder := dialect.GetSQLBuilder()

	saveStr, values := sqlBuilder.BuildUpdateSQL(object, entity)
	fmt.Println(saveStr)
	_, err := tx.Exec(saveStr, values...)
	if err != nil {
		fmt.Println("Update data failed:", err.Error())
		return nil, err
	}

	id := object[consts.META_ID]

	savedObject, err := QueryOneById(entity, id)
	if err != nil {
		fmt.Println("QueryOneById failed:", err.Error())
		return nil, err
	}
	return savedObject, nil
}

func (tx *Tx) doSaveOne(object map[string]interface{}, entity *model.Entity) (interface{}, error) {
	if object[consts.META_ID] == nil {
		object[consts.META_ID] = utils.CreateId()
		return tx.doInsertOne(object, entity)
	} else {
		return tx.doUpdateOne(object, entity)
	}
}

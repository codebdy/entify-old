package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"rxdrag.com/entity-engine/config"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/utils"
)

func dataFields(object map[string]interface{}) []string {
	return utils.MapStringKeys(object, "")
}

func insertFields(fields []string) string {
	return strings.Join(fields, ",")
}

func insertValueSymbols(fields []string) string {
	array := make([]string, len(fields))
	for i := range array {
		array[i] = "?"
	}
	return strings.Join(array, ",")
}

func values(object map[string]interface{}, entity *meta.Entity) []interface{} {
	objValues := make([]interface{}, 0, len(object))
	for key := range object {
		value := object[key]
		column := entity.GetColumn(key)

		if column == nil {
			panic("Can not find column:" + key)
		}

		if column.Type == meta.COLUMN_SIMPLE_JSON ||
			column.Type == meta.COLUMN_SIMPLE_ARRAY ||
			column.Type == meta.COLUMN_JSON_ARRAY {
			value, _ = json.Marshal(value)
		}
		fmt.Println("Make Field", key)
		objValues = append(objValues, value)
	}
	return objValues
}

func insertString(object map[string]interface{}, entity *meta.Entity) string {
	keys := dataFields(object)
	return fmt.Sprintf("INSERT INTO `%s`(%s) VALUES(%s)", entity.GetTableName(), insertFields(keys), insertValueSymbols(keys))
}

func updateSetFields(object map[string]interface{}) string {
	keys := dataFields(object)
	if len(keys) == 0 {
		panic("No update fields")
	}
	for i, key := range keys {
		keys[i] = key + "=?"
	}
	return strings.Join(keys, ",")
}

func updateString(object map[string]interface{}, entity *meta.Entity) string {
	return fmt.Sprintf("UPDATE `%s` SET %s WHERE ID = %s", entity.GetTableName(), updateSetFields(object), object["id"])
}

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

	saveStr := insertString(object, entity)
	fmt.Println("INSERT", saveStr)

	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(saveStr, values(object, entity)...)
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

	tx, err := db.Begin()
	defer clearTransaction(tx)

	saveStr := updateString(object, entity)
	if err != nil {
		return nil, err
	}
	values := values(object, entity)
	fmt.Println(saveStr, values)
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

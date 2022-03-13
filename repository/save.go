package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"rxdrag.com/entity-engine/config"
	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/utils"
)

func dataFields(object map[string]interface{}) []string {
	return utils.StringFilter(utils.MapStringKeys(object, ""), func(value string) bool {
		return value != consts.META_ID
	})
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
		if key == consts.META_ID {
			continue
		}
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
		fmt.Println("Insert Field", key)
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
	return fmt.Sprintf("UPDATE `%s` SET %s WHERE ID = '%s'", entity.GetTableName(), updateSetFields(object), object["id"])
}

func InsertOne(object map[string]interface{}, entity *meta.Entity) (interface{}, error) {
	fmt.Println(object)
	db, err := sql.Open(config.DRIVER_NAME, config.MYSQL_CONFIG)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	tx, err := db.Begin()

	saveStr := insertString(object, entity)
	fmt.Println("INSERT", saveStr)

	if err != nil {
		return nil, err
	}

	result, err := tx.Exec(saveStr, values(object, entity)...)
	if err != nil {
		fmt.Println("save data failed:", err.Error())
		return nil, err
	}
	id, err := result.LastInsertId()

	tx.Commit()
	fmt.Println("insert new record", id)
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
	fmt.Println(object)
	db, err := sql.Open(config.DRIVER_NAME, config.MYSQL_CONFIG)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	tx, err := db.Begin()

	saveStr := updateString(object, entity)

	fmt.Println(saveStr)

	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(saveStr, values(object, entity)...)
	if err != nil {
		fmt.Println("save data failed:", err.Error())
		return nil, err
	}

	id := object["id"]

	tx.Commit()
	fmt.Println("insert new record", id)
	savedObject, err := QueryOneById(entity, id)
	if err != nil {
		tx.Rollback()
		fmt.Println("QueryOneById failed:", err.Error())
		return nil, err
	}
	return savedObject, nil
}

func SaveOne(object map[string]interface{}, entity *meta.Entity) (interface{}, error) {
	if object["id"] == nil {
		return InsertOne(object, entity)
	} else {
		return UpdateOne(object, entity)
	}
	// return map[string]interface{}{
	// 	consts.RESPONSE_AFFECTEDROWS: affectedRows,
	// 	consts.RESPONSE_RETURNING:    insertedObject,
	// }, nil
}

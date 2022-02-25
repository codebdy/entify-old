package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"rxdrag.com/entity-engine/config"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/utils"
)

func fields(object map[string]interface{}) string {
	keys := utils.MapStringKeys(object, "`")
	return strings.Join(keys, ",")
}

func valueSymbols(object map[string]interface{}) string {
	array := make([]string, len(object))
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
		objValues = append(objValues, value)
	}
	return objValues
}

func SaveOneEntity(object map[string]interface{}, entity *meta.Entity) (interface{}, error) {
	fmt.Println(object)
	db, err := sql.Open("mysql", config.MYSQL_CONFIG)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	tx, err := db.Begin()

	saveStr := fmt.Sprintf("INSERT INTO `%s`(%s) VALUES(%s)", entity.GetTableName(), fields(object), valueSymbols(object))

	fmt.Println(saveStr)

	if err != nil {
		return nil, err
	}

	result, err := tx.Exec(saveStr, values(object, entity)...)
	if err != nil {
		fmt.Println("insert data failed:", err.Error())
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("fetch last insert id failed:", err.Error())
		return nil, err
	}
	tx.Commit()
	fmt.Println("insert new record", id)
	return nil, nil
}

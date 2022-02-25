package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"rxdrag.com/entity-engine/config"
	"rxdrag.com/entity-engine/meta"
)

func objectFields(object map[string]interface{}) string {
	var fieldsStr string
	for key := range object {
		fieldsStr = fieldsStr + "`" + key + "`"
	}

	return fieldsStr
}

func SaveOneEntity(object map[string]interface{}, entity *meta.Entity) (interface{}, error) {
	fmt.Println(object)
	db, err := sql.Open("mysql", config.MYSQL_CONFIG)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	saveStr := fmt.Sprintf("INSERT INTO `%s`(%s) VALUES(?)", entity.GetTableName(), objectFields(object))

	fmt.Println(saveStr)

	content, err := json.Marshal(object["content"])
	if err != nil {
		return nil, err
	}

	result, err := db.Exec(saveStr, content)
	if err != nil {
		fmt.Println("insert data failed:", err.Error())
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("fetch last insert id failed:", err.Error())
		return nil, err
	}
	fmt.Println("insert new record", id)
	return nil, nil
	//fmt.Println(p.Args)
	//fmt.Println(p.Conte
}

package repository

import (
	"database/sql"
	"fmt"

	"rxdrag.com/entity-engine/config"
	"rxdrag.com/entity-engine/meta"
)

func SaveOneEntity(object map[string]interface{}, entity *meta.Entity) (interface{}, error) {
	fmt.Println(object)
	db, err := sql.Open("mysql", config.MYSQL_CONFIG)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result, err := db.Exec("INSERT INTO `user`(`name`,`password`) VALUES('tom', 'tom')")
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

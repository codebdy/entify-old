package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/config"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/utils"
)

func makeValues(entity *meta.Entity) []interface{} {
	names := entity.ColumnNames()
	values := make([]interface{}, len(names))
	for i, columnName := range names {
		column := entity.GetColumn(columnName)
		if column.Type == meta.COLUMN_SIMPLE_JSON {
			var value utils.SimpleJSON
			values[i] = &value
		} else {
			var value sql.NullString
			values[i] = &value
		}
	}

	return values
}

func Query(entity *meta.Entity, queryStr string) ([]interface{}, error) {
	db, err := sql.Open(config.DRIVER_NAME, config.MYSQL_CONFIG)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	rows, err := db.Query(queryStr)
	columns, err := rows.Columns()
	var instances []interface{}
	for rows.Next() {
		row := make(map[string]interface{})
		values := makeValues(entity)
		err = rows.Scan(values...)
		for i, value := range values {
			row[columns[i]] = value
		}
		instances = append(instances, row)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	//fmt.Println(p.Args)
	//fmt.Println(p.Context.Value("data"))
	return instances, nil
}

func QueryOneById(entity *meta.Entity, id int64) (interface{}, error) {
	db, err := sql.Open(config.DRIVER_NAME, config.MYSQL_CONFIG)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	names := entity.ColumnNames()

	queryStr := "select %s from %s where id = ?"
	queryStr = fmt.Sprintf(queryStr, strings.Join(names, ","), entity.GetTableName())

	object := make(map[string]interface{})
	values := makeValues(entity)
	err = db.QueryRow(queryStr, id).Scan(values...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for i, value := range values {
		object[names[i]] = value
	}

	fmt.Println("Query one entity:" + entity.Name)
	return object, nil
}

func QueryResolveFn(entity *meta.Entity) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		queryStr := "select * from %s"

		queryStr = fmt.Sprintf(queryStr, entity.GetTableName())
		//err = db.Select(&instances, queryStr)
		return Query(entity, queryStr)
	}
}

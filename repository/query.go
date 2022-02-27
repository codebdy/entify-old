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
	fmt.Println("names", names)
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
	fmt.Println("呵呵")
	rows, err := db.Query(queryStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var instances []interface{}
	for rows.Next() {
		row := make(map[string]interface{})
		values := makeValues(entity)
		err = rows.Scan(values...)
		fmt.Println("呵呵2", values, columns)
		for i, value := range values {
			if nullValue, ok := value.(sql.NullString); ok {
				row[columns[i]] = nullValue.String
			} else {
				row[columns[i]] = value
			}

		}
		instances = append(instances, row)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(instances)
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
		names := entity.ColumnNames()
		queryStr := "select %s from %s "
		queryStr = fmt.Sprintf(queryStr, strings.Join(names, ","), entity.GetTableName())
		//err = db.Select(&instances, queryStr)
		return Query(entity, queryStr)
	}
}

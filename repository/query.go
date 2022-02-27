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
		switch column.Type {
		case meta.COLUMN_INT:
			var value sql.NullInt32
			values[i] = &value
			break
		case meta.COLUMN_FLOAT:
			var value sql.NullFloat64
			values[i] = &value
			break
		case meta.COLUMN_BOOLEAN:
			var value sql.NullBool
			values[i] = &value
			break
		case meta.COLUMN_DATE:
			var value sql.NullTime
			values[i] = &value
			break
		case meta.COLUMN_SIMPLE_JSON:
			var value utils.SimpleJSON
			values[i] = &value
			break
			// COLUMN_SIMPLE_ARRAY string = "simpleArray" ##待添加代码
			// COLUMN_JSON_ARRAY   string = "JsonArray"
		default:
			var value sql.NullString
			values[i] = &value
		}
	}

	return values
}

func convertValuesToObject(values []interface{}, entity *meta.Entity) map[string]interface{} {
	object := make(map[string]interface{})
	names := entity.ColumnNames()
	for i, value := range values {
		columnName := names[i]
		column := entity.GetColumn(columnName)
		switch column.Type {
		case meta.COLUMN_INT:
			nullValue := value.(*sql.NullInt64)
			if nullValue.Valid {
				object[columnName] = nullValue.Int64
			}
			break
		case meta.COLUMN_FLOAT:
			nullValue := value.(*sql.NullFloat64)
			if nullValue.Valid {
				object[columnName] = nullValue.Float64
			}
			break
		case meta.COLUMN_BOOLEAN:
			nullValue := value.(*sql.NullBool)
			if nullValue.Valid {
				object[columnName] = nullValue.Bool
			}
			break
		case meta.COLUMN_DATE:
			nullValue := value.(*sql.NullTime)
			if nullValue.Valid {
				object[columnName] = nullValue.Time
			}
			break
		case meta.COLUMN_SIMPLE_JSON:
			object[columnName] = value
			break
		case meta.COLUMN_JSON_ARRAY:
			object[columnName] = value
			break
		case meta.COLUMN_SIMPLE_ARRAY:
			object[columnName] = value
			break
		default:
			nullValue := value.(*sql.NullString)
			if nullValue.Valid {
				object[columnName] = nullValue.String
			}
		}

	}
	return object
}

func Query(entity *meta.Entity, queryStr string) ([]interface{}, error) {
	db, err := sql.Open(config.DRIVER_NAME, config.MYSQL_CONFIG)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	rows, err := db.Query(queryStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var instances []interface{}
	for rows.Next() {
		values := makeValues(entity)
		err = rows.Scan(values...)
		instances = append(instances, convertValuesToObject(values, entity))
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

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

	values := makeValues(entity)
	err = db.QueryRow(queryStr, id).Scan(values...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Query one entity:" + entity.Name)
	return convertValuesToObject(values, entity), nil
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

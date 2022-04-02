package repository

import (
	"database/sql"
	"fmt"

	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/oldmeta"
	"rxdrag.com/entity-engine/repositoryold/dialectold"
	"rxdrag.com/entity-engine/utils"
)

func makeValues(entity *model.Entity) []interface{} {
	names := entity.ColumnNames()
	values := make([]interface{}, len(names))
	for i, columnName := range names {
		column := entity.GetColumn(columnName)
		switch column.Type {
		case oldmeta.COLUMN_INT:
			var value sql.NullInt32
			values[i] = &value
			break
		case oldmeta.COLUMN_FLOAT:
			var value sql.NullFloat64
			values[i] = &value
			break
		case oldmeta.COLUMN_BOOLEAN:
			var value sql.NullBool
			values[i] = &value
			break
		case oldmeta.COLUMN_DATE:
			var value sql.NullTime
			values[i] = &value
			break
		case oldmeta.COLUMN_SIMPLE_JSON:
			var value utils.JSON
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

func convertValuesToObject(values []interface{}, entity *model.Entity) map[string]interface{} {
	object := make(map[string]interface{})
	names := entity.ColumnNames()
	for i, value := range values {
		columnName := names[i]
		column := entity.GetColumn(columnName)
		switch column.Type {
		case oldmeta.COLUMN_INT:
			nullValue := value.(*sql.NullInt64)
			if nullValue.Valid {
				object[columnName] = nullValue.Int64
			}
			break
		case oldmeta.COLUMN_FLOAT:
			nullValue := value.(*sql.NullFloat64)
			if nullValue.Valid {
				object[columnName] = nullValue.Float64
			}
			break
		case oldmeta.COLUMN_BOOLEAN:
			nullValue := value.(*sql.NullBool)
			if nullValue.Valid {
				object[columnName] = nullValue.Bool
			}
			break
		case oldmeta.COLUMN_DATE:
			nullValue := value.(*sql.NullTime)
			if nullValue.Valid {
				object[columnName] = nullValue.Time
			}
			break
		case oldmeta.COLUMN_SIMPLE_JSON:
			object[columnName] = value
			break
		case oldmeta.COLUMN_JSON_ARRAY:
			object[columnName] = value
			break
		case oldmeta.COLUMN_SIMPLE_ARRAY:
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

func (con *Connection) doQueryEntity(entity *model.Entity, args map[string]interface{}) ([]interface{}, error) {
	builder := dialectold.GetSQLBuilder()
	queryStr, params := builder.BuildQuerySQL(entity, args)
	rows, err := con.dbx.Query(queryStr, params...)
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

func (con *Connection) QueryOneById(entity *model.Entity, id interface{}) (interface{}, error) {
	return con.doQueryOne(entity, QueryArg{
		consts.ARG_WHERE: QueryArg{
			consts.ID: QueryArg{
				consts.AEG_EQ: id,
			},
		},
	})
}

func (con *Connection) doQueryOne(entity *model.Entity, args map[string]interface{}) (interface{}, error) {

	builder := dialectold.GetSQLBuilder()

	queryStr, params := builder.BuildQuerySQL(entity, args)

	values := makeValues(entity)
	err := con.dbx.QueryRow(queryStr, params...).Scan(values...)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Query one entity:" + entity.Name)
	return convertValuesToObject(values, entity), nil
}

func (con *Connection) doInsertOne(object map[string]interface{}, entity *model.Entity) (interface{}, error) {
	sqlBuilder := dialectold.GetSQLBuilder()
	saveStr, values := sqlBuilder.BuildInsertSQL(object, entity)

	for _, association := range entity.Associations {
		if object[association.Name] == nil {
			continue
		}
	}

	result, err := con.dbx.Exec(saveStr, values...)
	if err != nil {
		fmt.Println("Insert data failed:", err.Error())
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId failed:", err.Error())
		return nil, err
	}
	savedObject, err := con.QueryOneById(entity, id)
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

func (con *Connection) doUpdateOne(object map[string]interface{}, entity *model.Entity) (interface{}, error) {

	sqlBuilder := dialectold.GetSQLBuilder()

	saveStr, values := sqlBuilder.BuildUpdateSQL(object, entity)
	fmt.Println(saveStr)
	_, err := con.dbx.Exec(saveStr, values...)
	if err != nil {
		fmt.Println("Update data failed:", err.Error())
		return nil, err
	}

	id := object[consts.META_ID]

	savedObject, err := con.QueryOneById(entity, id)
	if err != nil {
		fmt.Println("QueryOneById failed:", err.Error())
		return nil, err
	}
	return savedObject, nil
}

func (con *Connection) doSaveOne(object map[string]interface{}, entity *model.Entity) (interface{}, error) {
	if object[consts.META_ID] == nil {
		return con.doInsertOne(object, entity)
	} else {
		return con.doUpdateOne(object, entity)
	}
}

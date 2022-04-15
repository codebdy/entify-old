package repository

import (
	"database/sql"
	"fmt"

	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/db"
	"rxdrag.com/entity-engine/db/dialect"
	"rxdrag.com/entity-engine/model/data"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/model/meta"
	"rxdrag.com/entity-engine/utils"
)

func makeValues(node graph.Node) []interface{} {
	names := node.AllAttributeNames()
	values := make([]interface{}, len(names))
	for i, attrName := range names {
		attr := node.GetAttributeByName(attrName)
		switch attr.Type {
		case meta.ID:
			var value db.NullUint64
			values[i] = &value
			break
		case meta.INT:
			var value sql.NullInt64
			values[i] = &value
			break
		case meta.FLOAT:
			var value sql.NullFloat64
			values[i] = &value
			break
		case meta.BOOLEAN:
			var value sql.NullBool
			values[i] = &value
			break
		case meta.DATE:
			var value sql.NullTime
			values[i] = &value
			break
		case meta.CLASS_VALUE_OBJECT,
			meta.ID_ARRAY,
			meta.INT_ARRAY,
			meta.FLOAT_ARRAY,
			meta.STRING_ARRAY,
			meta.DATE_ARRAY,
			meta.ENUM_ARRAY,
			meta.VALUE_OBJECT_ARRAY,
			meta.ENTITY_ARRAY:
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

func convertValuesToObject(values []interface{}, node graph.Node) map[string]interface{} {
	object := make(map[string]interface{})
	names := node.AllAttributeNames()
	for i, value := range values {
		attrName := names[i]
		column := node.GetAttributeByName(attrName)
		switch column.Type {
		case meta.ID:
			nullValue := value.(*db.NullUint64)
			if nullValue.Valid {
				object[attrName] = nullValue.Uint64
			}
			break
		case meta.INT:
			nullValue := value.(*sql.NullInt64)
			if nullValue.Valid {
				object[attrName] = nullValue.Int64
			}
			break
		case meta.FLOAT:
			nullValue := value.(*sql.NullFloat64)
			if nullValue.Valid {
				object[attrName] = nullValue.Float64
			}
			break
		case meta.BOOLEAN:
			nullValue := value.(*sql.NullBool)
			if nullValue.Valid {
				object[attrName] = nullValue.Bool
			}
			break
		case meta.DATE:
			nullValue := value.(*sql.NullTime)
			if nullValue.Valid {
				object[attrName] = nullValue.Time
			}
			break
		case meta.VALUE_OBJECT,
			meta.ID_ARRAY,
			meta.INT_ARRAY,
			meta.FLOAT_ARRAY,
			meta.STRING_ARRAY,
			meta.DATE_ARRAY,
			meta.ENUM_ARRAY,
			meta.VALUE_OBJECT_ARRAY,
			meta.ENTITY_ARRAY:
			object[attrName] = value
			break
		default:
			nullValue := value.(*sql.NullString)
			if nullValue.Valid {
				object[attrName] = nullValue.String
			}
		}

	}
	return object
}

func (con *Connection) doQueryEntity(node graph.Node, args map[string]interface{}) ([]interface{}, error) {
	builder := dialect.GetSQLBuilder()
	queryStr, params := builder.BuildQuerySQL(node, args)
	rows, err := con.Dbx.Query(queryStr, params...)
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
		values := makeValues(node)
		err = rows.Scan(values...)
		instances = append(instances, convertValuesToObject(values, node))
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	//fmt.Println(p.Context.Value("data"))
	return instances, nil
}

func (con *Connection) QueryOneById(node graph.Node, id interface{}) (interface{}, error) {
	return con.doQueryOne(node, QueryArg{
		consts.ARG_WHERE: QueryArg{
			consts.ID: QueryArg{
				consts.AEG_EQ: id,
			},
		},
	})
}

func (con *Connection) doQueryOne(node graph.Node, args map[string]interface{}) (interface{}, error) {

	builder := dialect.GetSQLBuilder()

	queryStr, params := builder.BuildQuerySQL(node, args)

	values := makeValues(node)
	err := con.Dbx.QueryRow(queryStr, params...).Scan(values...)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Query one entity:" + node.Name())
	return convertValuesToObject(values, node), nil
}

func (con *Connection) doInsertOne(instance *data.Instance) (interface{}, error) {
	sqlBuilder := dialect.GetSQLBuilder()
	saveStr := sqlBuilder.BuildInsertSQL(instance.Fields, instance.Table())
	values := data.MakeValues(instance.Fields)
	// for _, association := range entity.AllAssociations() {
	// 	if object[association.Name()] == nil {
	// 		continue
	// 	}
	// }

	result, err := con.Dbx.Exec(saveStr, values...)
	if err != nil {
		fmt.Println("Insert data failed:", err.Error())
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId failed:", err.Error())
		return nil, err
	}
	for _, ref := range instance.References {
		con.doSaveReference(ref, uint64(id))
	}

	savedObject, err := con.QueryOneById(instance.Entity, id)
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

func (con *Connection) doUpdateOne(instance *data.Instance) (interface{}, error) {

	sqlBuilder := dialect.GetSQLBuilder()

	saveStr := sqlBuilder.BuildUpdateSQL(instance.Id, instance.Fields, instance.Table())
	values := data.MakeValues(instance.Fields)
	fmt.Println(saveStr)
	_, err := con.Dbx.Exec(saveStr, values...)
	if err != nil {
		fmt.Println("Update data failed:", err.Error())
		return nil, err
	}

	for _, ref := range instance.References {
		con.doSaveReference(ref, instance.Id)
	}

	savedObject, err := con.QueryOneById(instance.Entity, instance.Id)

	if err != nil {
		fmt.Println("QueryOneById failed:", err.Error())
		return nil, err
	}
	return savedObject, nil
}

func (con *Connection) doSaveReference(r *data.Reference, ownerId uint64) {

}

func (con *Connection) doSaveOne(instance *data.Instance) (interface{}, error) {
	if instance.IsInsert() {
		return con.doInsertOne(instance)
	} else {
		return con.doUpdateOne(instance)
	}
}

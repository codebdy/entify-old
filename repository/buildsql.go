package repository

import (
	"fmt"
	"strings"

	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/repository/dialect"
)

func BuildQuerySQL(entity *model.Entity, args map[string]interface{}) (string, []interface{}) {
	var params []interface{}
	names := entity.ColumnNames()
	sqlBuilder := dialect.GetSQLBuilder()
	queryStr := "select %s from %s WHERE true "
	queryStr = fmt.Sprintf(queryStr, strings.Join(names, ","), entity.GetTableName())
	if args[consts.ARG_WHERE] != nil {
		whereStr, whereParams := sqlBuilder.BuildBoolExp(args[consts.ARG_WHERE].(map[string]interface{}))
		queryStr = queryStr + " " + whereStr
		params = append(params, whereParams...)
	}

	queryStr = queryStr + " order by id desc"
	fmt.Println("查询SQL:", queryStr)
	return queryStr, params
}

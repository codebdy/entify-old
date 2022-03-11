package repository

import (
	"fmt"
	"strings"

	"rxdrag.com/entity-engine/meta"
)

func BuildQuerySQL(entity *meta.Entity, args map[string]interface{}) (string, []interface{}) {
	var params []interface{}
	names := entity.ColumnNames()
	//sqlBuilder := dialect.GetSQLBuilder()

	//whereStr, whereParams := sqlBuilder.BuildBoolExp(args[consts.ARG_WHERE])

	queryStr := "select %s from %s order by id desc"
	queryStr = fmt.Sprintf(queryStr, strings.Join(names, ","), entity.GetTableName())

	return queryStr, params
}

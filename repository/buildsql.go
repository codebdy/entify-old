package repository

import (
	"fmt"
	"strings"

	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/repository/dialect"
)

func BuildQuerySQL(entity *meta.Entity, args map[string]interface{}) (string, []interface{}) {
	var params []interface{}
	names := entity.ColumnNames()
	sqlBuilder := dialect.GetSQLBuilder()
	queryStr := "select %s from %s order by id desc "
	if args[consts.ARG_WHERE] != nil {
		whereStr, whereParams := sqlBuilder.BuildBoolExp(args[consts.ARG_WHERE].(map[string]interface{}))
		queryStr = queryStr + " " + whereStr
		params = append(params, whereParams...)
	}

	queryStr = fmt.Sprintf(queryStr, strings.Join(names, ","), entity.GetTableName())

	return queryStr, params
}

package repository

import (
	"fmt"
	"strings"

	"rxdrag.com/entity-engine/meta"
)

type SQLBuilder interface {
	BoolExp(where map[string]interface{}) (string, []interface{})
}

func BuildQuerySQL(entity *meta.Entity, args map[string]interface{}) (string, []interface{}) {
	var params []interface{}
	names := entity.ColumnNames()

	queryStr := "select %s from %s order by id desc"
	queryStr = fmt.Sprintf(queryStr, strings.Join(names, ","), entity.GetTableName())

	return queryStr, params
}

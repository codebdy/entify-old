package dialect

import (
	"encoding/json"
	"fmt"
	"strings"

	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/utils"
)

type MySQLBuilder struct {
}

func (*MySQLBuilder) BuildFieldExp(fieldName string, fieldArgs map[string]interface{}) (string, []interface{}) {
	var params []interface{}
	queryStr := "true "
	for key, value := range fieldArgs {
		switch key {
		case consts.ARG_EQ:
			queryStr = queryStr + " AND " + fieldName + "=?"
			params = append(params, value)
			break
		case consts.ARG_ISNULL:
			if value == true {
				queryStr = queryStr + " AND ISNULL(" + fieldName + ")"
			}
			break
		default:
			panic("Can not find token:" + key)
		}
	}
	return "(" + queryStr + ")", params
}

func (b *MySQLBuilder) BuildBoolExp(where map[string]interface{}) (string, []interface{}) {
	var params []interface{}
	queryStr := ""
	for key, value := range where {
		switch key {
		case consts.ARG_AND:
			break
		case consts.ARG_NOT:
			break
		case consts.ARG_OR:
			break
		default:
			fiedleStr, fieldParam := b.BuildFieldExp(key, value.(map[string]interface{}))
			queryStr = queryStr + " AND " + fiedleStr
			params = append(params, fieldParam...)
		}
	}
	return queryStr, params
}

func (b *MySQLBuilder) ColumnTypeSQL(column *model.Column) string {
	typeStr := "text"
	switch column.Type {
	case meta.COLUMN_ID:
		typeStr = "bigint(64)"
		break
	case meta.COLUMN_INT:
		typeStr = "int"
		if column.Length == 1 {
			typeStr = "tinyint"
		} else if column.Length == 2 {
			typeStr = "smallint"
		} else if column.Length == 3 {
			typeStr = "mediumint"
		} else if column.Length == 4 {
			typeStr = "int"
		} else if column.Length > 4 {
			length := column.Length
			if length > 64 {
				length = 64
			}
			typeStr = fmt.Sprintf("bigint(%d)", length)
		}
		if column.Unsigned {
			typeStr = typeStr + " UNSIGNED"
		}
		break
	case meta.COLUMN_FLOAT:
		if column.Length > 4 {
			typeStr = "double"
		} else {
			typeStr = "float"
		}
		if column.FloatM > 0 && column.FloatD > 0 && column.FloatM >= column.FloatD {
			typeStr = fmt.Sprint(typeStr+"(%d,%d)", column.FloatM, column.FloatD)
		}
		if column.Unsigned {
			typeStr = typeStr + " UNSIGNED"
		}
		break
	case meta.COLUMN_BOOLEAN:
		typeStr = "tinyint(1)"
		break
	case meta.COLUMN_STRING:
		typeStr = "text"
		if column.Length > 0 {
			if column.Length <= 255 {
				typeStr = fmt.Sprintf("varchar(%d)", column.Length)
			} else if column.Length <= 65535 {
				typeStr = "text"
			} else if column.Length <= 16777215 {
				typeStr = "mediumtext"
			} else {
				typeStr = "longtext"
			}
		}
		break
	case meta.COLUMN_DATE:
		typeStr = "datetime"
		break
	case meta.COLUMN_SIMPLE_JSON:
		typeStr = "json"
		break
	case meta.COLUMN_SIMPLE_ARRAY:
		typeStr = "json"
		break
	case meta.COLUMN_JSON_ARRAY:
		typeStr = "json"
		break
	case meta.COLUMN_ENUM:
		typeStr = "tinytext"
		break
	}
	return typeStr
}

func (b *MySQLBuilder) BuildColumnSQL(column *model.Column) string {
	sql := "`" + column.Name + "` " + b.ColumnTypeSQL(column)
	if column.Generated {
		sql = sql + " AUTO_INCREMENT"
	}
	return sql
}

func (b *MySQLBuilder) BuildCreateTableSQL(table *model.Table) string {
	sql := "CREATE TABLE `%s` (%s)"
	fieldSqls := make([]string, len(table.Columns))
	for i := range table.Columns {
		columnSql := b.BuildColumnSQL(table.Columns[i])
		if table.Columns[i].Nullable {
			columnSql = columnSql + " NULL"
		} else {
			columnSql = columnSql + " NOT NULL"
		}
		fieldSqls[i] = columnSql
	}
	for _, column := range table.Columns {
		if column.Primary {
			fieldSqls = append(fieldSqls, fmt.Sprintf("PRIMARY KEY (`%s`)", column.Name))
		}
	}

	//建索引
	for _, column := range table.Columns {
		if column.Index {
			indexSql := "INDEX %s ( `%s`)"
			fieldSqls = append(fieldSqls, fmt.Sprintf(indexSql, column.Name+consts.INDEX_SUFFIX, column.Name))
		}
	}

	sql = fmt.Sprintf(sql, table.Name, strings.Join(fieldSqls, ","))

	return sql
}

func (b *MySQLBuilder) BuildDeleteTableSQL(table *model.Table) string {
	return "DROP TABLE " + table.Name
}

func (b *MySQLBuilder) BuildModifyTableAtoms(diff *model.TableDiff) []model.ModifyAtom {
	var atoms []model.ModifyAtom
	if diff.OldTable.Name != diff.NewTable.Name {
		//修改表名
		atoms = append(atoms, model.ModifyAtom{
			ExcuteSQL: fmt.Sprintf("ALTER TABLE %s RENAME TO %s ", diff.OldTable.Name, diff.NewTable.Name),
			UndoSQL:   fmt.Sprintf("ALTER TABLE %s RENAME TO %s ", diff.NewTable.Name, diff.OldTable.Name),
		})
	}
	b.appendDeleteColumnAtoms(diff, &atoms)
	b.appendAddColumnAtoms(diff, &atoms)
	b.appendModifyColumnAtoms(diff, &atoms)
	return atoms
}

func (b *MySQLBuilder) BuildInsertSQL(object map[string]interface{}, entity *model.Entity) (string, []interface{}) {
	keys := utils.MapStringKeys(object, "")
	sql := fmt.Sprintf("INSERT INTO `%s`(%s) VALUES(%s)", entity.GetTableName(), insertFields(keys), insertValueSymbols(keys))

	values := makeValues(keys, object, entity)

	return sql, values
}
func (b *MySQLBuilder) BuildUpdateSQL(object map[string]interface{}, entity *model.Entity) (string, []interface{}) {
	keys := utils.MapStringKeys(object, "")
	sql := fmt.Sprintf(
		"UPDATE `%s` SET %s WHERE ID = %s",
		entity.GetTableName(),
		updateSetFields(keys),
		object[consts.ID],
	)
	return sql, makeValues(keys, object, entity)
}

func updateSetFields(keys []string) string {
	if len(keys) == 0 {
		panic("No update fields")
	}
	newKeys := make([]string, len(keys))
	for i, key := range keys {
		newKeys[i] = key + "=?"
	}
	return strings.Join(newKeys, ",")
}

func insertFields(fields []string) string {
	return strings.Join(fields, ",")
}

func insertValueSymbols(fields []string) string {
	array := make([]string, len(fields))
	for i := range array {
		array[i] = "?"
	}
	return strings.Join(array, ",")
}

func makeValues(keys []string, object map[string]interface{}, entity *model.Entity) []interface{} {
	objValues := make([]interface{}, 0, len(keys))
	for _, key := range keys {
		value := object[key]
		column := entity.GetColumn(key)
		if column == nil {
			panic("Can not find column:" + key)
		}

		if column.Type == meta.COLUMN_SIMPLE_JSON ||
			column.Type == meta.COLUMN_SIMPLE_ARRAY ||
			column.Type == meta.COLUMN_JSON_ARRAY {
			value, _ = json.Marshal(value)
		}
		fmt.Println("Make Field", key)
		objValues = append(objValues, value)
	}
	return objValues
}

func (b *MySQLBuilder) appendDeleteColumnAtoms(diff *model.TableDiff, atoms *[]model.ModifyAtom) {
	for _, column := range diff.DeleteColumns {
		//删除索引
		if column.Index {
			indexName := column.Name + consts.INDEX_SUFFIX
			*atoms = append(*atoms, model.ModifyAtom{
				ExcuteSQL: fmt.Sprintf("DROP INDEX %s ON %s ", indexName, diff.NewTable.Name),
				UndoSQL:   fmt.Sprintf("CREATE INDEX %s ON %s (%s)", indexName, diff.NewTable.Name, column.Name),
			})
		}
		//删除列
		*atoms = append(*atoms, model.ModifyAtom{
			ExcuteSQL: fmt.Sprintf("ALTER TABLE %s DROP  %s ", diff.NewTable.Name, column.Name),
			UndoSQL:   fmt.Sprintf("ALTER TABLE %s ADD COLUMN  %s %s", diff.NewTable.Name, column.Name, b.ColumnTypeSQL(column)),
		})
	}
}

func (b *MySQLBuilder) appendAddColumnAtoms(diff *model.TableDiff, atoms *[]model.ModifyAtom) {
	for _, column := range diff.AddColumns {
		//添加列
		*atoms = append(*atoms, model.ModifyAtom{
			ExcuteSQL: fmt.Sprintf("ALTER TABLE %s ADD COLUMN  %s %s", diff.NewTable.Name, column.Name, b.ColumnTypeSQL(column)),
			UndoSQL:   fmt.Sprintf("ALTER TABLE %s DROP  %s ", diff.NewTable.Name, column.Name),
		})
		//添加索引
		if column.Index {
			indexName := column.Name + consts.INDEX_SUFFIX
			*atoms = append(*atoms, model.ModifyAtom{
				ExcuteSQL: fmt.Sprintf("CREATE INDEX %s ON %s (%s)", indexName, diff.NewTable.Name, column.Name),
				UndoSQL:   fmt.Sprintf("DROP INDEX %s ON %s ", indexName, diff.NewTable.Name),
			})
		}
	}
}

func (b *MySQLBuilder) appendModifyColumnAtoms(diff *model.TableDiff, atoms *[]model.ModifyAtom) {
	for _, columnDiff := range diff.ModifyColumns {

		//删除索引
		if columnDiff.OldColumn.Index {
			indexName := columnDiff.OldColumn.Name + consts.INDEX_SUFFIX
			*atoms = append(*atoms, model.ModifyAtom{
				ExcuteSQL: fmt.Sprintf("DROP INDEX %s ON %s ", indexName, diff.OldTable.Name),
				UndoSQL:   fmt.Sprintf("CREATE INDEX %s ON %s (%s)", indexName, diff.OldTable.Name, columnDiff.OldColumn.Name),
			})
		}
		//更改列
		if columnDiff.OldColumn.Name != columnDiff.NewColumn.Name ||
			columnDiff.OldColumn.Type != columnDiff.NewColumn.Type ||
			columnDiff.OldColumn.Length != columnDiff.NewColumn.Length ||
			columnDiff.OldColumn.FloatD != columnDiff.NewColumn.FloatD ||
			columnDiff.OldColumn.FloatM != columnDiff.NewColumn.FloatM ||
			columnDiff.OldColumn.Unsigned != columnDiff.NewColumn.Unsigned {
			*atoms = append(*atoms, model.ModifyAtom{
				ExcuteSQL: fmt.Sprintf(
					"ALTER TABLE %s CHANGE COLUMN %s %s %s",
					diff.NewTable.Name,
					columnDiff.OldColumn.Name,
					columnDiff.NewColumn.Name, b.ColumnTypeSQL(columnDiff.NewColumn),
				),
				UndoSQL: fmt.Sprintf(
					"ALTER TABLE %s CHANGE COLUMN %s %s %s",
					diff.NewTable.Name,
					columnDiff.NewColumn.Name,
					columnDiff.OldColumn.Name, b.ColumnTypeSQL(columnDiff.OldColumn),
				),
			})
		}
		//添加索引
		if columnDiff.NewColumn.Index {
			indexName := columnDiff.NewColumn.Name + consts.INDEX_SUFFIX
			*atoms = append(*atoms, model.ModifyAtom{
				ExcuteSQL: fmt.Sprintf("CREATE INDEX %s ON %s (%s)", indexName, diff.NewTable.Name, columnDiff.NewColumn.Name),
				UndoSQL:   fmt.Sprintf("DROP INDEX %s ON %s ", indexName, diff.NewTable.Name),
			})
		}
	}
}

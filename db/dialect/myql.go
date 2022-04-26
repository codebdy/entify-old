package dialect

import (
	"fmt"
	"strings"
	"time"

	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/model"
	"rxdrag.com/entity-engine/model/data"
	"rxdrag.com/entity-engine/model/graph"
	"rxdrag.com/entity-engine/model/meta"
	"rxdrag.com/entity-engine/model/table"
	"rxdrag.com/entity-engine/utils"
)

type MySQLBuilder struct {
}

func (*MySQLBuilder) BuildFieldExp(fieldName string, fieldArgs map[string]interface{}) (string, []interface{}) {
	var params []interface{}
	queryStr := ""
	for key, value := range fieldArgs {
		switch key {
		case consts.ARG_EQ:
			queryStr = fieldName + "=?"
			params = append(params, value)
			break
		case consts.ARG_ISNULL:
			if value == true {
				queryStr = "ISNULL(" + fieldName + ")"
			}
			break
		default:
			panic("Can not find token:" + key)
		}
	}
	return "(" + queryStr + ")", params
}

func (b *MySQLBuilder) BuildBoolExp(argClass graph.ArgClass, where map[string]interface{}) (string, []interface{}) {
	var params []interface{}
	querys := []string{}
	for key, value := range where {
		switch key {
		case consts.ARG_AND:
			break
		case consts.ARG_NOT:
			break
		case consts.ARG_OR:
			break
		default:
			sqls := []string{}
			for i := range argClass.Children {
				child := argClass.Children[i]
				fiedleStr, fieldParam := b.BuildFieldExp(child.Alise()+"."+key, value.(map[string]interface{}))
				sqls = append(sqls, fiedleStr)
				params = append(params, fieldParam...)
			}
			querys = append(querys, fmt.Sprintf("(%s)", strings.Join(sqls, " OR ")))
		}
	}
	return strings.Join(querys, " AND "), params
}

func (b *MySQLBuilder) ColumnTypeSQL(column *table.Column) string {
	typeStr := "text"
	switch column.Type {
	case meta.ID:
		typeStr = "bigint(64)"
		break
	case meta.INT:
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
	case meta.FLOAT:
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
	case meta.BOOLEAN:
		typeStr = "tinyint(1)"
		break
	case meta.STRING:
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
	case meta.DATE:
		typeStr = "datetime"
		break
	case meta.ENUM:
		typeStr = "tinytext"
		break
	case meta.VALUE_OBJECT,
		meta.ID_ARRAY,
		meta.INT_ARRAY,
		meta.FLOAT_ARRAY,
		meta.STRING_ARRAY,
		meta.DATE_ARRAY,
		meta.ENUM_ARRAY,
		meta.VALUE_OBJECT_ARRAY:
		typeStr = "json"
		break
	}
	return typeStr
}

func (b *MySQLBuilder) BuildColumnSQL(column *table.Column) string {
	sql := "`" + column.Name + "` " + b.ColumnTypeSQL(column)
	if column.Name == consts.ID {
		sql = fmt.Sprintf(sql + " AUTO_INCREMENT")
	}
	return sql
}

func (b *MySQLBuilder) BuildCreateTableSQL(table *table.Table) string {
	sql := "CREATE TABLE `%s` (%s)"
	fieldSqls := make([]string, len(table.Columns))
	for i := range table.Columns {
		columnSql := b.BuildColumnSQL(table.Columns[i])
		columnSql = columnSql + nullableString(table.Columns[i].Nullable)
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
	fmt.Println("Create table sql:", sql)

	if table.EntityInnerId > 0 {
		sql = sql + fmt.Sprintf(" AUTO_INCREMENT = %d", utils.EncodeBaseId(table.EntityInnerId))
	}
	return sql
}

func (b *MySQLBuilder) BuildDeleteTableSQL(table *table.Table) string {
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

func buildArgAssociation(association *graph.ArgAssociation, owner *graph.ArgEntity) string {
	var sql string
	derivedAssociations := association.Association.DerivedAssociations()
	for i := range derivedAssociations {
		derivedAsso := derivedAssociations[i]
		if owner != nil {
			typeEntity := association.GetTypeEntity(derivedAsso.TypeEntity().Uuid())
			povitTableAlias := fmt.Sprintf("%s_%d_%d", graph.PREFIX_T, owner.Id, typeEntity.Id)
			sql = sql + fmt.Sprintf(
				" LEFT JOIN %s %s ON %s=%s LEFT JOIN %s %s ON %s=%s ",
				derivedAsso.Relation.Table.Name,
				povitTableAlias,
				owner.Alise()+"."+consts.ID,
				povitTableAlias+"."+owner.Entity.Table.Name,
				typeEntity.Entity.TableName(),
				typeEntity.Alise(),
				povitTableAlias+"."+typeEntity.Entity.Table.Name,
				typeEntity.Alise()+"."+consts.ID,
			)

			for i := range typeEntity.FromClass.Associations {
				sql = sql + buildArgAssociation(typeEntity.FromClass.Associations[i], typeEntity)
			}
		}
	}
	return sql
}

func (b *MySQLBuilder) BuildQuerySQLBody(argEntity *graph.ArgEntity, fields []*graph.Attribute) string {
	names := make([]string, len(fields))
	for i := range fields {
		names[i] = argEntity.Alise() + "." + fields[i].Name
	}
	queryStr := "select %s from %s %s "
	queryStr = fmt.Sprintf(queryStr, strings.Join(names, ","), argEntity.Entity.TableName(), argEntity.Alise())

	for i := range argEntity.FromClass.Associations {
		association := argEntity.FromClass.Associations[i]
		queryStr = queryStr + " " + buildArgAssociation(association, argEntity)
	}
	return queryStr
}

func (b *MySQLBuilder) BuildWhereSQL(
	argEntity *graph.ArgEntity,
	fields []*graph.Attribute,
	where map[string]interface{},
) (string, []interface{}) {
	whereStr := ""
	var params []interface{}
	if where != nil {
		boolStr, whereParams := b.BuildBoolExp(*argEntity.FromClass, where)
		whereStr = " WHERE " + boolStr
		params = append(params, whereParams...)
	}
	return whereStr, params
}

func (b *MySQLBuilder) BuildOrderBySQL(
	argEntity *graph.ArgEntity,
	orderBy interface{},
) string {
	if _, ok := orderBy.(graph.QueryArg); ok {

	}
	return fmt.Sprintf(" ORDER BY %s.id DESC", argEntity.Alise())
}

// func (b *MySQLBuilder) BuildQuerySQL(
// 	tableName string,
// 	fields []*graph.Attribute,
// 	args map[string]interface{},
// ) (string, []interface{}) {
// 	var params []interface{}
// 	names := make([]string, len(fields))
// 	for i := range fields {
// 		names[i] = fields[i].Name
// 	}
// 	queryStr := "select %s from %s WHERE true "
// 	queryStr = fmt.Sprintf(queryStr, strings.Join(names, ","), tableName)
// 	if args[consts.ARG_WHERE] != nil {
// 		whereStr, whereParams := b.BuildBoolExp(args[consts.ARG_WHERE].(map[string]interface{}))
// 		queryStr = queryStr + " " + whereStr
// 		params = append(params, whereParams...)
// 	}

// 	queryStr = queryStr + " order by id desc"
// 	fmt.Println("查询SQL:", queryStr)
// 	return queryStr, params
// }

func associationFieldSQL(node graph.Noder) string {
	names := node.AllAttributeNames()
	for i := range names {
		names[i] = "a." + names[i]
	}
	return strings.Join(names, ",")
}

func (b *MySQLBuilder) BuildQueryByIdsSQL(entity *graph.Entity, idCounts int) string {
	parms := make([]string, idCounts)

	for i := range parms {
		parms[i] = "?"
	}
	queryStr := "select %s from %s WHERE id in(%s) "
	names := entity.AllAttributeNames()
	queryStr = fmt.Sprintf(queryStr,
		strings.Join(names, ","),
		entity.TableName(),
		strings.Join(parms, ","),
	)

	fmt.Println("BuildQueryByIdsSQL:", queryStr)
	return queryStr
}

func (b *MySQLBuilder) BuildQueryAssociatedInstancesSQL(
	node graph.Noder,
	ownerId uint64,
	povitTableName string,
	ownerFieldName string,
	typeFieldName string,
) string {
	queryStr := "select %s from %s a INNER JOIN %s b ON a.id = b.%s WHERE b.%s=%d "
	queryStr = fmt.Sprintf(queryStr,
		associationFieldSQL(node),
		node.Entity().TableName(),
		povitTableName,
		typeFieldName,
		ownerFieldName,
		ownerId)

	fmt.Println("BuildQueryAssociatedInstancesSQL:", queryStr)
	return queryStr
}

func (b *MySQLBuilder) BuildBatchAssociationSQL(
	tableName string,
	fields []*graph.Attribute,
	ids []uint64,
	povitTableName string,
	ownerFieldName string,
	typeFieldName string,
) string {
	queryStr := "select %s, b.%s as %s from %s a INNER JOIN %s b ON a.id = b.%s WHERE b.%s in (%s) "
	parms := make([]string, len(ids))
	names := make([]string, len(fields))
	for i := range parms {
		parms[i] = fmt.Sprintf("%d", ids[i])
	}
	for i := range fields {
		names[i] = "a." + fields[i].Name
	}

	queryStr = fmt.Sprintf(queryStr,
		strings.Join(names, ","),
		ownerFieldName,
		consts.ASSOCIATION_OWNER_ID,
		tableName,
		povitTableName,
		typeFieldName,
		ownerFieldName,
		strings.Join(parms, ","),
	)

	return queryStr
}

func (b *MySQLBuilder) BuildInsertSQL(fields []*data.Field, table *table.Table) string {
	sql := fmt.Sprintf("INSERT INTO `%s`(%s) VALUES(%s)", table.Name, insertFields(fields), insertValueSymbols(fields))

	return sql
}
func (b *MySQLBuilder) BuildUpdateSQL(id uint64, fields []*data.Field, table *table.Table) string {
	sql := fmt.Sprintf(
		"UPDATE `%s` SET %s WHERE ID = %d",
		table.Name,
		updateSetFields(fields),
		id,
	)

	return sql
}

func (b *MySQLBuilder) BuildClearAssociationSQL(ownerId uint64, tableName string, ownerFieldName string) string {
	sql := fmt.Sprintf(
		"DELETE FROM `%s` WHERE (`%s` = '%d')",
		tableName,
		ownerFieldName,
		ownerId,
	)
	return sql
}

func (b *MySQLBuilder) BuildDeleteSQL(id uint64, tableName string) string {
	sql := fmt.Sprintf(
		"DELETE FROM `%s` WHERE (`%s` = '%d')",
		tableName,
		"id",
		id,
	)
	return sql
}

func (b *MySQLBuilder) BuildSoftDeleteSQL(id uint64, tableName string) string {
	sql := fmt.Sprintf(
		"UPDATE `%s` SET `%s` = '%s' WHERE (`%s` = %d)",
		tableName,
		consts.DELETED_AT,
		time.Now(),
		"id",
		id,
	)
	return sql
}

func (b *MySQLBuilder) BuildQueryPovitSQL(povit *data.AssociationPovit) string {
	return fmt.Sprintf(
		"SELECT * FROM `%s` WHERE (`%s` = %d AND `%s` = %d)",
		povit.Table().Name,
		povit.Source.Column.Name,
		povit.Source.Value,
		povit.Target.Column.Name,
		povit.Target.Value,
	)
}
func (b *MySQLBuilder) BuildInsertPovitSQL(povit *data.AssociationPovit) string {
	return fmt.Sprintf(
		"INSERT INTO `%s`(%s,%s) VALUES(%d, %d)",
		povit.Table().Name,
		povit.Source.Column.Name,
		povit.Target.Column.Name,
		povit.Source.Value,
		povit.Target.Value,
	)
}

func (b *MySQLBuilder) BuildDeletePovitSQL(povit *data.AssociationPovit) string {
	return fmt.Sprintf(
		"DELETE FROM `%s` WHERE (`%s` = %d AND `%s` = %d)",
		povit.Table().Name,
		povit.Source.Column.Name,
		povit.Source.Value,
		povit.Target.Column.Name,
		povit.Target.Value,
	)
}

func updateSetFields(fields []*data.Field) string {
	if len(fields) == 0 {
		panic("No update fields")
	}
	newKeys := make([]string, len(fields))
	for i, field := range fields {
		newKeys[i] = field.Column.Name + "=?"
	}
	return strings.Join(newKeys, ",")
}

func insertFields(fields []*data.Field) string {
	names := make([]string, len(fields))
	for i := range fields {
		names[i] = fields[i].Column.Name
	}
	return strings.Join(names, ",")
}

func insertValueSymbols(fields []*data.Field) string {
	array := make([]string, len(fields))
	for i := range array {
		array[i] = "?"
	}
	return strings.Join(array, ",")
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
			ExcuteSQL: fmt.Sprintf("ALTER TABLE %s ADD COLUMN  %s %s %s", diff.NewTable.Name, column.Name, b.ColumnTypeSQL(column), nullableString(column.Nullable)),
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
				ExcuteSQL: fmt.Sprintf("DROP INDEX %s ON %s ", indexName, diff.NewTable.Name), //表名已在前面的步骤中被修改，这里用新表名
				UndoSQL:   fmt.Sprintf("CREATE INDEX %s ON %s (%s)", indexName, diff.NewTable.Name, columnDiff.OldColumn.Name),
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
		if columnDiff.OldColumn.Nullable != columnDiff.NewColumn.Nullable {
			*atoms = append(*atoms, model.ModifyAtom{
				ExcuteSQL: fmt.Sprintf(
					"ALTER TABLE %s MODIFY  %s %s %s",
					diff.NewTable.Name,
					columnDiff.NewColumn.Name,
					b.ColumnTypeSQL(columnDiff.NewColumn),
					nullableString(columnDiff.NewColumn.Nullable),
				),
				UndoSQL: fmt.Sprintf(
					"ALTER TABLE %s MODIFY  %s %s %s",
					diff.NewTable.Name,
					columnDiff.NewColumn.Name,
					b.ColumnTypeSQL(columnDiff.OldColumn),
					nullableString(!columnDiff.NewColumn.Nullable),
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

func nullableString(nullable bool) string {
	if nullable {
		return " NULL "
	}
	return " NOT NULL "
}

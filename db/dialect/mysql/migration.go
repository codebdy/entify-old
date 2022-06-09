package mysql

import (
	"fmt"
	"strings"

	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model"
	"rxdrag.com/entify/model/meta"
	"rxdrag.com/entify/model/table"
	"rxdrag.com/entify/utils"
)

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
			fieldSqls = append(fieldSqls, fmt.Sprintf("PRIMARY KEY (%s)", column.Name))
		}
	}

	if table.PKString != "" {
		fieldSqls = append(fieldSqls, fmt.Sprintf("PRIMARY KEY (%s)", table.PKString))
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
	//主键
	atoms = append(atoms, model.ModifyAtom{
		ExcuteSQL: fmt.Sprintf("ALTER TABLE %s DROP  PRIMARY  KEY, ADD PRIMARY KEY (%s)", diff.OldTable.Name, diff.OldTable.PKString),
		UndoSQL:   fmt.Sprintf("ALTER TABLE %s DROP  PRIMARY  KEY,ADD PRIMARY KEY (%s) ", diff.NewTable.Name, diff.NewTable.PKString),
	})

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

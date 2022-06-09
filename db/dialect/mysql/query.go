package mysql

import (
	"fmt"
	"strings"

	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/model/data"
	"rxdrag.com/entify/model/graph"
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
		case consts.ARG_IN:
			values := value.([]string)
			placeHolders := []string{}
			for i := range values {
				placeHolders = append(placeHolders, "?")
				params = append(params, values[i])
			}
			if len(placeHolders) > 0 {
				queryStr = fieldName + fmt.Sprintf(" in(%s)", strings.Join(placeHolders, ","))
			}
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

func (b *MySQLBuilder) BuildBoolExp(argEntity *graph.ArgEntity, where map[string]interface{}) (string, []interface{}) {
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
			asso := argEntity.Entity.GetAssociationByName(key)
			if asso == nil {
				fieldStr, fieldParam := b.BuildFieldExp(argEntity.Alise()+"."+key, value.(map[string]interface{}))
				params = append(params, fieldParam...)
				querys = append(querys, fmt.Sprintf("(%s)", fieldStr))
			} else {
				argAsso := argEntity.GetAssociation(key)
				var associStrs []string
				var associParams []interface{}
				for i := range argAsso.ArgEntities {
					assoStr, assoParam := b.BuildBoolExp(argAsso.ArgEntities[i], value.(map[string]interface{}))
					if assoStr != "" {
						assoStr = fmt.Sprintf("(%s)", assoStr)
						associStrs = append(associStrs, assoStr)
						associParams = append(associParams, assoParam...)
					}
				}
				querys = append(querys, strings.Join(associStrs, " OR "))
				params = append(params, associParams...)
			}
		}
	}
	return strings.Join(querys, " AND "), params
}

func buildArgAssociation(argAssociation *graph.ArgAssociation, owner *graph.ArgEntity) string {
	var sql string

	if !argAssociation.Association.IsAbstract() {
		if owner != nil {
			typeEntity := argAssociation.GetTypeEntity(argAssociation.Association.TypeClass().Uuid())
			povitTableAlias := fmt.Sprintf("%s_%d_%d", graph.PREFIX_T, owner.Id, typeEntity.Id)
			sql = sql + fmt.Sprintf(
				" LEFT JOIN %s %s ON %s=%s LEFT JOIN %s %s ON %s=%s ",
				argAssociation.Association.Relation.Table.Name,
				povitTableAlias,
				owner.Alise()+"."+consts.ID,
				povitTableAlias+"."+owner.Entity.Table.Name,
				typeEntity.Entity.TableName(),
				typeEntity.Alise(),
				povitTableAlias+"."+typeEntity.Entity.Table.Name,
				typeEntity.Alise()+"."+consts.ID,
			)

			for i := range typeEntity.Associations {
				sql = sql + buildArgAssociation(typeEntity.Associations[i], typeEntity)
			}
		}
		return sql
	}
	derivedAssociations := argAssociation.Association.DerivedAssociations()
	for i := range derivedAssociations {
		derivedAsso := derivedAssociations[i]
		if owner != nil {
			typeEntity := argAssociation.GetTypeEntity(derivedAsso.TypeEntity().Uuid())
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

			for i := range typeEntity.Associations {
				sql = sql + buildArgAssociation(typeEntity.Associations[i], typeEntity)
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

	for i := range argEntity.Associations {
		association := argEntity.Associations[i]
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
		boolStr, whereParams := b.BuildBoolExp(argEntity, where)
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

func (b *MySQLBuilder) BuildClearAssociationSQL(ownerId uint64, tableName string, ownerFieldName string) string {
	sql := fmt.Sprintf(
		"DELETE FROM `%s` WHERE (`%s` = '%d')",
		tableName,
		ownerFieldName,
		ownerId,
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

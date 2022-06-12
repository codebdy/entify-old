package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"rxdrag.com/entify/config"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/db"
	"rxdrag.com/entify/db/dialect"
	"rxdrag.com/entify/model/data"
	"rxdrag.com/entify/model/graph"
)

type InsanceData = map[string]interface{}

func (con *Connection) buildQueryInterfaceSQL(intf *graph.Interface, args map[string]interface{}) (string, []interface{}) {
	var (
		sqls       []string
		paramsList []interface{}
	)
	builder := dialect.GetSQLBuilder()
	for i := range intf.Children {
		entity := intf.Children[i]
		argEntity := graph.BuildArgEntity(
			entity,
			con.v.WeaveAuthInArgs(entity.Uuid(), args[consts.ARG_WHERE]),
			con,
		)
		queryStr := builder.BuildQuerySQLBody(argEntity, intf.AllAttributes())
		if where, ok := args[consts.ARG_WHERE].(graph.QueryArg); ok {
			whereSQL, params := builder.BuildWhereSQL(argEntity, intf.AllAttributes(), where)
			if whereSQL != "" {
				queryStr = queryStr + " WHERE " + whereSQL
			}

			paramsList = append(paramsList, params...)
		}
		queryStr = queryStr + builder.BuildOrderBySQL(argEntity, args[consts.ARG_ORDERBY])

		sqls = append(sqls, queryStr)
	}

	return strings.Join(sqls, " UNION "), paramsList
}

func (con *Connection) buildQueryEntitySQL(entity *graph.Entity, args map[string]interface{}) (string, []interface{}) {
	var paramsList []interface{}
	argEntity := graph.BuildArgEntity(
		entity,
		con.v.WeaveAuthInArgs(entity.Uuid(), args[consts.ARG_WHERE]),
		con,
	)
	builder := dialect.GetSQLBuilder()
	queryStr := builder.BuildQuerySQLBody(argEntity, entity.AllAttributes())
	if where, ok := args[consts.ARG_WHERE].(graph.QueryArg); ok {
		whereSQL, params := builder.BuildWhereSQL(argEntity, entity.AllAttributes(), where)
		if whereSQL != "" {
			queryStr = queryStr + " WHERE " + whereSQL
		}
		paramsList = append(paramsList, params...)
	}

	queryStr = queryStr + builder.BuildOrderBySQL(argEntity, args[consts.ARG_ORDERBY])

	return queryStr, paramsList
}

func (con *Connection) doQueryInterface(intf *graph.Interface, args map[string]interface{}) []InsanceData {
	sql, params := con.buildQueryInterfaceSQL(intf, args)

	rows, err := con.Dbx.Query(sql, params...)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}
	var instances []InsanceData
	for rows.Next() {
		values := makeQueryValues(intf)
		err = rows.Scan(values...)
		if err != nil {
			panic(err.Error())
		}
		instances = append(instances, convertValuesToObject(values, intf))
	}

	instancesIds := make([]interface{}, len(instances))
	for i := range instances {
		instancesIds[i] = instances[i][consts.ID]
	}

	for i := range intf.Children {
		child := intf.Children[i]
		oneEntityInstances := con.doQueryByIds(child, instancesIds)
		merageInstances(instances, oneEntityInstances)
	}

	return instances
}

func (con *Connection) doQueryEntity(entity *graph.Entity, args map[string]interface{}) []InsanceData {
	sql, params := con.buildQueryEntitySQL(entity, args)
	fmt.Println("doQueryEntity SQL:", sql, params)
	rows, err := con.Dbx.Query(sql, params...)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}
	var instances []InsanceData
	for rows.Next() {
		values := makeQueryValues(entity)
		err = rows.Scan(values...)
		if err != nil {
			panic(err.Error())
		}
		instances = append(instances, convertValuesToObject(values, entity))
	}

	return instances
}

func (con *Connection) QueryOneEntityById(entity *graph.Entity, id interface{}) interface{} {
	return con.doQueryOneEntity(entity, graph.QueryArg{
		consts.ARG_WHERE: graph.QueryArg{
			consts.ID: graph.QueryArg{
				consts.ARG_EQ: id,
			},
		},
	})
}
func (con *Connection) doQueryOneInterface(intf *graph.Interface, args map[string]interface{}) interface{} {
	querySql, params := con.buildQueryInterfaceSQL(intf, args)

	values := makeQueryValues(intf)
	err := con.Dbx.QueryRow(querySql, params...).Scan(values...)

	switch {
	case err == sql.ErrNoRows:
		return nil
	case err != nil:
		panic(err.Error())
	}

	instance := convertValuesToObject(values, intf)
	for i := range intf.Children {
		child := intf.Children[i]
		oneEntityInstances := con.doQueryByIds(child, []interface{}{instance[consts.ID]})
		if len(oneEntityInstances) > 0 {
			return oneEntityInstances[0]
		}
	}
	return nil
}

func (con *Connection) doQueryOneEntity(entity *graph.Entity, args map[string]interface{}) interface{} {
	queryStr, params := con.buildQueryEntitySQL(entity, args)

	values := makeQueryValues(entity)
	fmt.Println("doQueryOneEntity SQL:", queryStr)
	err := con.Dbx.QueryRow(queryStr, params...).Scan(values...)
	switch {
	case err == sql.ErrNoRows:
		return nil
	case err != nil:
		panic(err.Error())
	}

	return convertValuesToObject(values, entity)
}

func (con *Connection) doInsertOne(instance *data.Instance) (interface{}, error) {
	sqlBuilder := dialect.GetSQLBuilder()
	saveStr := sqlBuilder.BuildInsertSQL(instance.Fields, instance.Table())
	values := makeSaveValues(instance.Fields)
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
	for _, asso := range instance.Associations {
		err = con.doSaveAssociation(asso, uint64(id))
		if err != nil {
			fmt.Println("Save reference failed:", err.Error())
			return nil, err
		}
	}

	savedObject := con.QueryOneEntityById(instance.Entity, id)

	//affectedRows, err := result.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected failed:", err.Error())
		return nil, err
	}

	return savedObject, nil
}

func (con *Connection) doQueryAssociatedInstances(r data.Associationer, ownerId uint64) []InsanceData {
	var instances []InsanceData
	builder := dialect.GetSQLBuilder()
	entity := r.TypeEntity()
	queryStr := builder.BuildQueryAssociatedInstancesSQL(entity, ownerId, r.Table().Name, r.OwnerColumn().Name, r.TypeColumn().Name)
	rows, err := con.Dbx.Query(queryStr)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		values := makeQueryValues(entity)
		err = rows.Scan(values...)
		if err != nil {
			panic(err.Error())
		}
		instances = append(instances, convertValuesToObject(values, entity))
	}

	return instances
}

func (con *Connection) doQueryByIds(entity *graph.Entity, ids []interface{}) []InsanceData {
	var instances []map[string]interface{}
	builder := dialect.GetSQLBuilder()
	sql := builder.BuildQueryByIdsSQL(entity, len(ids))
	rows, err := con.Dbx.Query(sql, ids...)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		values := makeQueryValues(entity)
		err = rows.Scan(values...)
		if err != nil {
			panic(err.Error())
		}
		instances = append(instances, convertValuesToObject(values, entity))
	}

	return instances
}

func (con *Connection) doBatchAbstractRealAssociations(
	association *graph.Association,
	ids []uint64,
	args graph.QueryArg,
	v *AbilityVerifier,
) []InsanceData {
	var (
		instances  []InsanceData
		sqls       []string
		paramsList []interface{}
	)
	builder := dialect.GetSQLBuilder()
	abstractTypeClass := association.TypeClass()

	derivedAssociations := association.DerivedAssociations()
	for i := range derivedAssociations {
		derivedAsso := derivedAssociations[i]
		entity := derivedAsso.TypeEntity()
		argEntity := graph.BuildArgEntity(
			entity,
			con.v.WeaveAuthInArgs(entity.Uuid(), args[consts.ARG_WHERE]),
			con,
		)
		queryStr := builder.BuildBatchAssociationBodySQL(argEntity,
			abstractTypeClass.AllAttributes(),
			derivedAsso.Relation.Table.Name,
			derivedAsso.Owner().TableName(),
			derivedAsso.TypeEntity().TableName(),
			ids,
		)
		if where, ok := args[consts.ARG_WHERE].(graph.QueryArg); ok {
			whereSQL, params := builder.BuildWhereSQL(argEntity, entity.AllAttributes(), where)
			if whereSQL != "" {
				queryStr = queryStr + " AND " + whereSQL
			}

			paramsList = append(paramsList, params...)
		}
		queryStr = queryStr + builder.BuildOrderBySQL(argEntity, args[consts.ARG_ORDERBY])
		sqls = append(sqls, queryStr)
	}
	sql := strings.Join(sqls, " UNION ")
	fmt.Println("doBatchAbstractRealAssociations SQL:" + sql)
	rows, err := con.Dbx.Query(sql, paramsList...)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		values := makeQueryValues(abstractTypeClass)
		var idValue db.NullUint64
		values = append(values, &idValue)
		err = rows.Scan(values...)
		if err != nil {
			panic(err.Error())
		}
		instance := convertValuesToObject(values, abstractTypeClass)
		instance[consts.ASSOCIATION_OWNER_ID] = values[len(values)-1].(*db.NullUint64).Uint64
		instances = append(instances, instance)
	}

	instancesIds := make([]interface{}, len(instances))
	for i := range instances {
		instancesIds[i] = instances[i][consts.ID]
	}

	for i := range derivedAssociations {
		derivedAsso := derivedAssociations[i]
		oneEntityInstances := con.doQueryByIds(derivedAsso.TypeEntity(), instancesIds)
		merageInstances(instances, oneEntityInstances)
	}

	return instances
}

func (con *Connection) doBatchRealAssociations(
	association *graph.Association,
	ids []uint64,
	args graph.QueryArg,
	v *AbilityVerifier,
) []InsanceData {
	var instances []map[string]interface{}
	var paramsList []interface{}

	builder := dialect.GetSQLBuilder()
	typeClass := association.TypeClass()
	entity := association.TypeClass().Entity()
	argEntity := graph.BuildArgEntity(
		entity,
		con.v.WeaveAuthInArgs(entity.Uuid(), args[consts.ARG_WHERE]),
		con,
	)

	queryStr := builder.BuildBatchAssociationBodySQL(argEntity,
		typeClass.AllAttributes(),
		association.Relation.Table.Name,
		association.Owner().TableName(),
		association.TypeClass().TableName(),
		ids,
	)

	if where, ok := args[consts.ARG_WHERE].(graph.QueryArg); ok {
		whereSQL, params := builder.BuildWhereSQL(argEntity, entity.AllAttributes(), where)
		if whereSQL != "" {
			queryStr = queryStr + " AND " + whereSQL
		}
		paramsList = append(paramsList, params...)
	}

	queryStr = queryStr + builder.BuildOrderBySQL(argEntity, args[consts.ARG_ORDERBY])
	fmt.Println("doBatchRealAssociations SQL:	", queryStr)
	rows, err := con.Dbx.Query(queryStr, paramsList...)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		values := makeQueryValues(typeClass)
		var idValue db.NullUint64
		values = append(values, &idValue)
		err = rows.Scan(values...)
		if err != nil {
			panic(err.Error())
		}
		instance := convertValuesToObject(values, typeClass)
		instance[consts.ASSOCIATION_OWNER_ID] = values[len(values)-1].(*db.NullUint64).Uint64
		instances = append(instances, instance)
	}

	return instances
}

func merageInstances(source []InsanceData, target []InsanceData) {
	for i := range source {
		souceObj := source[i]
		for j := range target {
			targetObj := target[j]
			if souceObj[consts.ID] == targetObj[consts.ID] {
				targetObj[consts.ASSOCIATION_OWNER_ID] = souceObj[consts.ASSOCIATION_OWNER_ID]
				source[i] = targetObj
			}
		}
	}
}

func (con *Connection) doUpdateOne(instance *data.Instance) (interface{}, error) {

	sqlBuilder := dialect.GetSQLBuilder()

	saveStr := sqlBuilder.BuildUpdateSQL(instance.Id, instance.Fields, instance.Table())
	values := makeSaveValues(instance.Fields)
	fmt.Println(saveStr)
	_, err := con.Dbx.Exec(saveStr, values...)
	if err != nil {
		fmt.Println("Update data failed:", err.Error())
		return nil, err
	}

	for _, ref := range instance.Associations {
		con.doSaveAssociation(ref, instance.Id)
	}

	savedObject := con.QueryOneEntityById(instance.Entity, instance.Id)

	return savedObject, nil
}

func newAssociationPovit(r data.Associationer, ownerId uint64, tarId uint64) *data.AssociationPovit {
	sourceId := ownerId
	targetId := tarId

	if !r.IsSource() {
		sourceId = targetId
		targetId = ownerId
	}

	return data.NewAssociationPovit(r, sourceId, targetId)

}

func (con *Connection) doSaveAssociation(r data.Associationer, ownerId uint64) error {
	for _, ins := range r.Deleted() {
		if r.Cascade() {
			con.doDeleteInstance(ins)
		} else {
			povit := newAssociationPovit(r, ownerId, ins.Id)
			con.doDeleteAssociationPovit(povit)
		}
	}

	for _, ins := range r.Added() {
		saved, err := con.doSaveOne(ins)
		if err != nil {
			return err
		}

		if savedIns, ok := saved.(InsanceData); ok {
			tarId := savedIns[consts.ID].(uint64)
			relationInstance := newAssociationPovit(r, ownerId, tarId)
			con.doSaveAssociationPovit(relationInstance)
		} else {
			panic("Save Association error")
		}

	}

	for _, ins := range r.Updated() {
		if ins.Id == 0 {
			panic("Can not add new instance when update")
		}
		saved, err := con.doSaveOne(ins)
		if err != nil {
			return err
		}

		if savedIns, ok := saved.(InsanceData); ok {
			tarId := savedIns[consts.ID].(uint64)
			relationInstance := newAssociationPovit(r, ownerId, tarId)

			con.doSaveAssociationPovit(relationInstance)
		} else {
			panic("Save Association error")
		}
	}

	synced := r.Synced()
	if len(synced) == 0 {
		return nil
	}

	con.clearAssociation(r, ownerId)

	for _, ins := range synced {
		if ins.Id == 0 {
			panic("Can not add new instance when update")
		}
		saved, err := con.doSaveOne(ins)
		if err != nil {
			return err
		}

		if savedIns, ok := saved.(InsanceData); ok {
			tarId := savedIns[consts.ID].(uint64)
			relationInstance := newAssociationPovit(r, ownerId, tarId)

			con.doSaveAssociationPovit(relationInstance)
		} else {
			panic("Save Association error")
		}
	}

	return nil
}

func (con *Connection) clearAssociation(r data.Associationer, ownerId uint64) {
	con.deleteAssociationPovit(r, ownerId)

	if r.Cascade() {
		con.deleteAssociatedInstances(r, ownerId)
	}
}

func (con *Connection) deleteAssociationPovit(r data.Associationer, ownerId uint64) {
	sqlBuilder := dialect.GetSQLBuilder()
	sql := sqlBuilder.BuildClearAssociationSQL(ownerId, r.Table().Name, r.OwnerColumn().Name)
	_, err := con.Dbx.Exec(sql)
	if err != nil {
		panic(err.Error())
	}
}

func (con *Connection) deleteAssociatedInstances(r data.Associationer, ownerId uint64) {
	typeEntity := r.TypeEntity()
	associatedInstances := con.doQueryAssociatedInstances(r, ownerId)
	for i := range associatedInstances {
		ins := data.NewInstance(associatedInstances[i], typeEntity)
		con.doDeleteInstance(ins)
	}
}

func (con *Connection) doSaveAssociationPovit(povit *data.AssociationPovit) {
	sqlBuilder := dialect.GetSQLBuilder()
	sql := sqlBuilder.BuildQueryPovitSQL(povit)
	rows, err := con.Dbx.Query(sql)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}
	if !rows.Next() {
		sql = sqlBuilder.BuildInsertPovitSQL(povit)
		_, err := con.Dbx.Exec(sql)
		if err != nil {
			panic(err.Error())
		}
	}
}

func (con *Connection) doDeleteAssociationPovit(povit *data.AssociationPovit) {
	sqlBuilder := dialect.GetSQLBuilder()
	sql := sqlBuilder.BuildDeletePovitSQL(povit)
	_, err := con.Dbx.Exec(sql)
	if err != nil {
		panic(err.Error())
	}
}

func (con *Connection) doSaveOne(instance *data.Instance) (interface{}, error) {
	if instance.IsInsert() {
		return con.doInsertOne(instance)
	} else {
		return con.doUpdateOne(instance)
	}
}

func (con *Connection) doDeleteInstance(instance *data.Instance) {
	var sql string
	sqlBuilder := dialect.GetSQLBuilder()
	tableName := instance.Table().Name
	if instance.Entity.IsSoftDelete() {
		sql = sqlBuilder.BuildSoftDeleteSQL(instance.Id, tableName)
	} else {
		sql = sqlBuilder.BuildDeleteSQL(instance.Id, tableName)
	}

	_, err := con.Dbx.Exec(sql)
	if err != nil {
		panic(err.Error())
	}

	associstions := instance.Associations
	for i := range associstions {
		asso := associstions[i]
		if asso.IsCombination() {
			if !asso.TypeEntity().IsSoftDelete() {
				con.deleteAssociationPovit(asso, instance.Id)
			}
			con.deleteAssociatedInstances(asso, instance.Id)
		}
	}
}

func (con *Connection) doCheckEntity(name string) bool {
	sqlBuilder := dialect.GetSQLBuilder()
	var count int
	err := con.Dbx.QueryRow(sqlBuilder.BuildTableCheckSQL(name, config.GetDbConfig().Database)).Scan(&count)
	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		panic(err.Error())
	}
	return count > 0
}

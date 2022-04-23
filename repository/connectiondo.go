package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"rxdrag.com/entity-engine/consts"
	"rxdrag.com/entity-engine/db"
	"rxdrag.com/entity-engine/db/dialect"
	"rxdrag.com/entity-engine/model/data"
	"rxdrag.com/entity-engine/model/graph"
)

func (con *Connection) doQueryEntity(node graph.Noder, args map[string]interface{}) ([]interface{}, error) {
	builder := dialect.GetSQLBuilder()
	queryStr, params := builder.BuildQuerySQL(node, args)
	rows, err := con.Dbx.Query(queryStr, params...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var instances []interface{}
	for rows.Next() {
		values := makeQueryValues(node)
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

func (con *Connection) QueryOneById(node graph.Noder, id interface{}) (map[string]interface{}, error) {
	return con.doQueryOne(node, QueryArg{
		consts.ARG_WHERE: QueryArg{
			consts.ID: QueryArg{
				consts.AEG_EQ: id,
			},
		},
	})
}

func (con *Connection) doQueryOne(node graph.Noder, args map[string]interface{}) (map[string]interface{}, error) {

	builder := dialect.GetSQLBuilder()

	queryStr, params := builder.BuildQuerySQL(node, args)

	values := makeQueryValues(node)
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

func (con *Connection) doInsertOne(instance *data.Instance) (map[string]interface{}, error) {
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

func (con *Connection) doQueryAssociatedInstances(r data.Associationer, ownerId uint64) []map[string]interface{} {
	var instances []map[string]interface{}
	builder := dialect.GetSQLBuilder()
	entity := r.TypeEntity()
	queryStr := builder.BuildQueryAssociatedInstancesSQL(entity, ownerId, r.Table().Name, r.OwnerColumn().Name, r.TypeColumn().Name)
	rows, err := con.Dbx.Query(queryStr)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		values := makeQueryValues(entity)
		err = rows.Scan(values...)
		instances = append(instances, convertValuesToObject(values, entity))
	}
	if err != nil {
		panic(err.Error())
	}

	return instances
}

func (con *Connection) doBatchQueryAssociations(association *graph.Association, ids []uint64) []map[string]interface{} {
	var (
		instances []map[string]interface{}
		sqls      []string
	)
	builder := dialect.GetSQLBuilder()
	abstractTypeClass := association.TypeClass()
	if association.IsAbstract() {
		derivedAssociations := association.DerivedAssociations()
		for i := range derivedAssociations {
			derivedAsso := derivedAssociations[i]
			queryStr := builder.BuildBatchAssociationSQL(derivedAsso.TypeEntity().TableName(),
				abstractTypeClass.AllAttributes(),
				ids,
				derivedAsso.Relation.Table.Name,
				derivedAsso.Owner().TableName(),
				derivedAsso.TypeEntity().TableName(),
			)
			sqls = append(sqls, queryStr)
		}
	} else {
		queryStr := builder.BuildBatchAssociationSQL(association.TypeClass().TableName(),
			abstractTypeClass.AllAttributes(),
			ids,
			association.Relation.Table.Name,
			association.Owner().TableName(),
			association.TypeClass().TableName(),
		)
		sqls = append(sqls, queryStr)
	}
	sql := strings.Join(sqls, " UNION ")
	rows, err := con.Dbx.Query(sql)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		values := makeQueryValues(abstractTypeClass)
		var idValue db.NullUint64
		values = append(values, &idValue)
		err = rows.Scan(values...)
		instance := convertValuesToObject(values, abstractTypeClass)
		instance[consts.ASSOCIATION_OWNER_ID] = values[len(values)-1].(*db.NullUint64).Uint64
		instances = append(instances, instance)
	}
	if err != nil {
		panic(err.Error())
	}

	return instances
}

func (con *Connection) doUpdateOne(instance *data.Instance) (map[string]interface{}, error) {

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

	savedObject, err := con.QueryOneById(instance.Entity, instance.Id)

	if err != nil {
		fmt.Println("QueryOneById failed:", err.Error())
		return nil, err
	}
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

		tarId := saved[consts.ID].(uint64)
		relationInstance := newAssociationPovit(r, ownerId, tarId)

		con.doSaveAssociationPovit(relationInstance)
	}

	for _, ins := range r.Updated() {
		if ins.Id == 0 {
			panic("Can not add new instance when update")
		}
		saved, err := con.doSaveOne(ins)
		if err != nil {
			return err
		}

		tarId := saved[consts.ID].(uint64)
		relationInstance := newAssociationPovit(r, ownerId, tarId)

		con.doSaveAssociationPovit(relationInstance)
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

		tarId := saved[consts.ID].(uint64)
		relationInstance := newAssociationPovit(r, ownerId, tarId)

		con.doSaveAssociationPovit(relationInstance)
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

func (con *Connection) doSaveOne(instance *data.Instance) (map[string]interface{}, error) {
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

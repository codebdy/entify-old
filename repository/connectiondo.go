package repository

import (
	"database/sql"
	"fmt"

	"rxdrag.com/entity-engine/consts"
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
	// for _, association := range entity.AllAssociations() {
	// 	if object[association.Name()] == nil {
	// 		continue
	// 	}
	// }

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

func newAssociationInstance(r data.Associationer, ownerId uint64, tarId uint64) *data.AssociationInstance {
	sourceId := ownerId
	targetId := tarId

	if !r.IsSource() {
		sourceId = targetId
		targetId = ownerId
	}

	return data.NewAssociationInstance(r, sourceId, targetId)
}

func (con *Connection) doSaveAssociation(r data.Associationer, ownerId uint64) error {
	for _, ins := range r.Deleted() {
		if r.Cascade() {
			con.doDeleteInstance(ins)
		} else {
			relationInstance := newAssociationInstance(r, ownerId, ins.Id)
			err := con.doDeleteAssociationInstance(relationInstance)
			if err != nil {
				panic(err.Error())
			}
		}
	}

	for _, ins := range r.Added() {
		saved, err := con.doSaveOne(ins)
		if err != nil {
			return err
		}

		tarId := saved[consts.ID].(uint64)
		relationInstance := newAssociationInstance(r, ownerId, tarId)

		con.doSaveAssociationInstance(relationInstance)
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
		relationInstance := newAssociationInstance(r, ownerId, tarId)

		con.doSaveAssociationInstance(relationInstance)
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
		relationInstance := newAssociationInstance(r, ownerId, tarId)

		con.doSaveAssociationInstance(relationInstance)
	}

	return nil
}

func (con *Connection) clearAssociation(r data.Associationer, ownerId uint64) {
	sqlBuilder := dialect.GetSQLBuilder()
	sql := sqlBuilder.BuildClearAssociationSQL(ownerId, r.Table().Name, r.OwnerColumn().Name)
	_, err := con.Dbx.Exec(sql)
	if err != nil {
		panic(err.Error())
	}

	if r.Cascade() {
		con.deleteAssociatedInstances(r, ownerId)
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

func (con *Connection) doSaveAssociationInstance(instance *data.AssociationInstance) (interface{}, error) {
	return nil, nil
}

func (con *Connection) doDeleteAssociationInstance(instance *data.AssociationInstance) error {
	return nil
}

func (con *Connection) doSaveOne(instance *data.Instance) (map[string]interface{}, error) {
	if instance.IsInsert() {
		return con.doInsertOne(instance)
	} else {
		return con.doUpdateOne(instance)
	}
}

func (con *Connection) doDeleteInstance(instance *data.Instance) error {
	return nil
}

package repository

import (
	"fmt"

	"rxdrag.com/entify/db/dialect"
	"rxdrag.com/entify/model/data"
	"rxdrag.com/entify/model/graph"
)

type QueryArg = map[string]interface{}

func QueryInterface(intf *graph.Interface, args QueryArg) []InsanceData {
	con, err := Open()
	if err != nil {
		panic(err.Error())
	}
	return con.doQueryInterface(intf, args)
}

func QueryOneInterface(intf *graph.Interface, args QueryArg) interface{} {
	con, err := Open()
	if err != nil {
		panic(err.Error())
	}
	return con.doQueryOneInterface(intf, args)
}

func QueryEntity(entity *graph.Entity, args QueryArg) []InsanceData {
	con, err := Open()
	if err != nil {
		panic(err.Error())
	}
	return con.doQueryEntity(entity, args)
}

func QueryOneEntity(entity *graph.Entity, args QueryArg) interface{} {
	con, err := Open()
	if err != nil {
		panic(err.Error())
	}
	return con.doQueryOneEntity(entity, args)
}

func SaveOne(instance *data.Instance) (interface{}, error) {
	con, err := Open()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	err = con.BeginTx()
	defer con.ClearTx()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	obj, err := con.doSaveOne(instance)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	err = con.Commit()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return obj, nil
}

func InsertOne(instance *data.Instance) (interface{}, error) {
	con, err := Open()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer con.ClearTx()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	obj, err := con.doInsertOne(instance)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	err = con.Commit()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return obj, nil
}

func BatchQueryAssociations(association *graph.Association, ids []uint64) []map[string]interface{} {
	con, err := Open()
	if err != nil {
		panic(err.Error())
	}
	return con.doBatchAssociations(association, ids)
}

func IsEntityExists(name string) bool {
	con, err := Open()
	if err != nil {
		panic(err.Error())
	}
	return con.doCheckEntity(name)
}

func Install() error {
	sqlBuilder := dialect.GetSQLBuilder()
	con, err := Open()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = con.BeginTx()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	_, err = con.Dbx.Exec(sqlBuilder.BuildCreateMetaSQL())
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	_, err = con.Dbx.Exec(sqlBuilder.BuildCreateAbilitySQL())
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	_, err = con.Dbx.Exec(sqlBuilder.BuildCreateEntityAuthSettingsSQL())
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = con.Commit()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

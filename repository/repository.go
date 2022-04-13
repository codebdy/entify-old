package repository

import (
	"rxdrag.com/entity-engine/model/data"
	"rxdrag.com/entity-engine/model/graph"
)

type QueryArg = map[string]interface{}

func Query(node graph.Node, args map[string]interface{}) ([]interface{}, error) {
	con, err := Open()
	defer con.Close()
	if err != nil {
		panic(err.Error())
	}
	return con.doQueryEntity(node, args)
}

func QueryOne(node graph.Node, args map[string]interface{}) (interface{}, error) {
	con, err := Open()
	defer con.Close()
	if err != nil {
		panic(err.Error())
	}
	return con.doQueryOne(node, args)
}

func SaveOne(instance *data.Instance) (interface{}, error) {
	con, err := Open()
	defer con.Close()
	if err != nil {
		panic(err.Error())
	}
	err = con.BeginTx()
	defer con.ClearTx()
	if err != nil {
		panic(err.Error())
	}

	obj, err := con.doSaveOne(instance)
	if err != nil {
		panic(err.Error())
	}
	err = con.Commit()
	if err != nil {
		panic(err.Error())
	}
	return obj, nil
}

func InsertOne(object map[string]interface{}, entity *graph.Entity) (interface{}, error) {
	con, err := Open()
	defer con.Close()
	if err != nil {
		panic(err.Error())
	}
	err = con.Close()
	defer con.ClearTx()
	if err != nil {
		panic(err.Error())
	}

	obj, err := con.doInsertOne(object, entity)
	if err != nil {
		panic(err.Error())
	}
	err = con.Commit()
	if err != nil {
		panic(err.Error())
	}
	return obj, nil
}

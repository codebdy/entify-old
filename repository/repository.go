package repository

import (
	"rxdrag.com/entity-engine/model"
)

type QueryArg = map[string]interface{}

func Query(entity *model.Entity, args map[string]interface{}) ([]interface{}, error) {
	con, err := OpenConnection()
	defer con.Close()
	if err != nil {
		panic(err.Error())
	}
	return con.doQueryEntity(entity, args)
}

func QueryOne(entity *model.Entity, args map[string]interface{}) (interface{}, error) {
	con, err := OpenConnection()
	defer con.Close()
	if err != nil {
		panic(err.Error())
	}
	return con.doQueryOne(entity, args)
}

func SaveOne(object map[string]interface{}, entity *model.Entity) (interface{}, error) {
	con, err := OpenConnection()
	defer con.Close()
	if err != nil {
		panic(err.Error())
	}
	err = con.Begin()
	defer con.clearTx()
	if err != nil {
		panic(err.Error())
	}

	obj, err := con.doSaveOne(object, entity)
	if err != nil {
		panic(err.Error())
	}
	err = con.Commit()
	if err != nil {
		panic(err.Error())
	}
	return obj, nil
}

func InsertOne(object map[string]interface{}, entity *model.Entity) (interface{}, error) {
	con, err := OpenConnection()
	defer con.Close()
	if err != nil {
		panic(err.Error())
	}
	err = con.Close()
	defer con.clearTx()
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

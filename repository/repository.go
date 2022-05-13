package repository

import (
	"fmt"

	"rxdrag.com/entify/config"
	"rxdrag.com/entify/model/data"
	"rxdrag.com/entify/model/graph"
)

type QueryArg = map[string]interface{}

func Query(node graph.Noder, args QueryArg) []InsanceData {
	con, err := Open(config.GetDbConfig())
	defer con.Close()
	if err != nil {
		panic(err.Error())
	}
	return con.doQueryNode(node, args)
}

func QueryOne(node graph.Noder, args QueryArg) interface{} {
	con, err := Open(config.GetDbConfig())
	defer con.Close()
	if err != nil {
		panic(err.Error())
	}
	return con.doQueryOneNode(node, args)
}

func SaveOne(instance *data.Instance) (interface{}, error) {
	con, err := Open(config.GetDbConfig())
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

func InsertOne(instance *data.Instance) (interface{}, error) {
	con, err := Open(config.GetDbConfig())
	defer con.Close()
	if err != nil {
		panic(err.Error())
	}
	err = con.Close()
	defer con.ClearTx()
	if err != nil {
		panic(err.Error())
	}

	obj, err := con.doInsertOne(instance)
	if err != nil {
		panic(err.Error())
	}
	err = con.Commit()
	if err != nil {
		panic(err.Error())
	}
	return obj, nil
}

func BatchQueryAssociations(association *graph.Association, ids []uint64) []map[string]interface{} {
	con, err := Open(config.GetDbConfig())
	defer con.Close()
	if err != nil {
		panic(err.Error())
	}
	return con.doBatchAssociations(association, ids)
}

func Install(cfg config.DbConfig) {
	con, err := Open(cfg)
	defer con.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = con.BeginTx()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	sql := `CREATE TABLE meta (
		id bigint NOT NULL AUTO_INCREMENT,
		content json DEFAULT NULL,
		publishedAt datetime DEFAULT NULL,
		createdAt datetime DEFAULT NULL,
		updatedAt datetime DEFAULT NULL,
		status varchar(45) DEFAULT NULL,
		PRIMARY KEY (id)
	) ENGINE=InnoDB AUTO_INCREMENT=1507236403010867251 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
	`
	_, err = con.Dbx.Exec(sql)
	if err != nil {
		panic(err.Error())
	}

	err = con.Commit()

	if err != nil {
		panic(err.Error())
	}
}

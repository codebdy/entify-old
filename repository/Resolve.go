package repository

import (
	"database/sql"
	"fmt"

	"github.com/graphql-go/graphql"
	"rxdrag.com/entity-engine/config"
	"rxdrag.com/entity-engine/meta"
	"rxdrag.com/entity-engine/utils"
)

func QueryResolve(entity *meta.EntityMeta) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		db, err := sql.Open("mysql", config.MYSQL_CONFIG)
		defer db.Close()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		queryStr := "select * from %s"

		queryStr = fmt.Sprintf(queryStr, entity.GetTableName())
		//err = db.Select(&instances, queryStr)
		rows, err := db.Query(queryStr)
		columns, err := rows.Columns()
		var instances []utils.SimpleJSON
		for rows.Next() {
			row := make(map[string]interface{})
			values := make([]interface{}, len(columns))
			for i, columnName := range columns {
				if columnName == "content" {
					var value utils.SimpleJSON
					values[i] = &value
				} else {
					var value string
					values[i] = &value
				}

			}
			err = rows.Scan(values...)
			for i, value := range values {
				row[columns[i]] = value
			}
			instances = append(instances, row)
		}
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		fmt.Println("Resolve entity:" + entity.Name)
		//fmt.Println(p.Args)
		//fmt.Println(p.Context.Value("data"))
		return instances, nil
	}
}

func PostOneResolve(entity *meta.EntityMeta) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		fmt.Println(p.Args["object"])
		db, err := sql.Open("mysql", config.MYSQL_CONFIG)
		defer db.Close()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		queryStr := "select * from %s"

		queryStr = fmt.Sprintf(queryStr, entity.GetTableName())
		//err = db.Select(&instances, queryStr)
		rows, err := db.Query(queryStr)
		columns, err := rows.Columns()
		var instances []utils.SimpleJSON
		for rows.Next() {
			row := make(map[string]interface{})
			values := make([]interface{}, len(columns))
			for i, columnName := range columns {
				if columnName == "content" {
					var value utils.SimpleJSON
					values[i] = &value
				} else {
					var value string
					values[i] = &value
				}

			}
			err = rows.Scan(values...)
			for i, value := range values {
				row[columns[i]] = value
			}
			instances = append(instances, row)
		}
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		fmt.Println("Resolve entity:" + entity.Name)
		//fmt.Println(p.Args)
		//fmt.Println(p.Context.Value("data"))
		return instances, nil
	}
}

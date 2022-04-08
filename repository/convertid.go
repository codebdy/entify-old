package repository

import (
	"strconv"

	"rxdrag.com/entity-engine/consts"
)

func ConvertId(object map[string]interface{}) map[string]interface{} {
	switch object[consts.ID].(type) {
	case string:
		id, err := strconv.ParseInt(object[consts.ID].(string), 10, 64)
		if err != nil {
			panic("Convert id error:" + err.Error())
		}

		object[consts.ID] = id
	}

	return object
}
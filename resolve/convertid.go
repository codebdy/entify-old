package resolve

import (
	"strconv"

	"rxdrag.com/entify/consts"
)

func ConvertId(object map[string]interface{}) map[string]interface{} {
	if object[consts.ID] == nil {
		return object
	}
	switch object[consts.ID].(type) {
	case string:
		id, err := strconv.ParseUint(object[consts.ID].(string), 10, 64)
		if err != nil {
			panic("Convert id error:" + err.Error())
		}

		object[consts.ID] = id
	}

	return object
}

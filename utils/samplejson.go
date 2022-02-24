package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type SimpleJSON map[string]interface{}

func (m SimpleJSON) Value() (driver.Value, error) {
	if len(m) == 0 {
		return nil, nil
	}
	j, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return driver.Value([]byte(j)), nil
}

func (m *SimpleJSON) Scan(src interface{}) error {
	var source []byte
	_m := make(map[string]interface{})

	switch src.(type) {
	case []uint8:
		source = []byte(src.([]uint8))
	case nil:
		return nil
	default:
		return errors.New("incompatible type for JSON")
	}
	err := json.Unmarshal(source, &_m)
	if err != nil {
		return err
	}
	*m = SimpleJSON(_m)
	return nil
}

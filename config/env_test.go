package config

import (
	"testing"

	"rxdrag.com/entify/consts"
)

func TestSetString(t *testing.T) {
	oldValue := GetString(consts.DB_DRIVER)
	SetString(consts.DB_DRIVER, "test_value")
	if GetString(consts.DB_DRIVER) != "test_value" {
		t.Error("Error SetString,expected 'test_value', but is:" + GetString(consts.DB_DRIVER))
	}
	SetString(consts.DB_DRIVER, oldValue)
}

func TestGetString(t *testing.T) {
	if GetString(consts.DB_DRIVER) != "mysql" {
		t.Error("Getstring Error:" + GetString(consts.DB_DRIVER))
	}
}

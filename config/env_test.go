package config

import (
	"testing"

	"rxdrag.com/entify/consts"
)

func TestSetString(t *testing.T) {
	Init()
	SetString(consts.DB_DRIVER, "test_value")
	if GetString(consts.DB_DRIVER) != "test_value" {
		t.Error("Error SetString,expected 'test_value', but is:" + GetString(consts.DB_DRIVER))
	}
}

func TestGetString(t *testing.T) {
	Init()
	if GetString(consts.DB_DRIVER) != "mysql" {
		t.Error("Getstring Error:" + GetString(consts.DB_DRIVER))
	}
}

package config

import (
	"testing"

	"rxdrag.com/entify/consts"
)

func TestGetString(t *testing.T) {
	if GetString(consts.DB_DRIVER) != "mysql" {
		t.Error("Getstring Error:" + GetString(consts.DB_DRIVER))
	}
}

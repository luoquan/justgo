package dao

import (
	"testing"
)

func TestGetRDb(t *testing.T) {
	db := GetRDb()
	if db == nil {
		t.Error("conn't get Read Connnection")
	}
	db.Close()
}

func TestGetWDb(t *testing.T) {
	db := GetWDb()
	if db == nil {
		t.Error("conn't get Write Connnection")
	}
}

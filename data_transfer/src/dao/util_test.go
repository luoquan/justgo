/**
 * Created with IntelliJ IDEA.
 * User: luoquan
 * Date: 13-5-21
 * Time: 下午7:42
 * To change this template use File | Settings | File Templates.
 */
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

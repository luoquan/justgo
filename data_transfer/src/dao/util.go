/**
 * Created with IntelliJ IDEA.
 * User: luoquan
 * Date: 13-5-21
 * Time: 上午12:01
 * To change this template use File | Settings | File Templates.
 */
package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func getDb(dbName string) *sql.DB {
	db, err := sql.Open("mysql", "root:luoquan@tcp(localhost:3306)/"+dbName+"?charset=utf8")
	if err != nil {
		panic("can not connect to mysql")
	}
	return db
}

func GetRDb() *sql.DB {
	return getDb("finance")
}

func GetWDb() *sql.DB {
	return getDb("test")
}

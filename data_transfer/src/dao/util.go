package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type Wf struct {
	Id       int
	Version  int
	Type     int
	Title    string
	Org_name string
	Status   string
}

func (wf Wf) String() string {
	return strconv.Itoa(wf.Id) + "-" + strconv.Itoa(wf.Version) + "-" + strconv.Itoa(wf.Type) +
		"-" + wf.Title + "-" + wf.Org_name + "-" + wf.Status
}

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

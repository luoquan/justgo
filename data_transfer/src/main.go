package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//select id,version,type,title,organization_full_name,status,has_opinion,create_time from common_workflow_task where id<10;

func main() {
	db, err := sql.Open("mysql", "root:luoquan@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic("can not connect to mysql")
	}
	defer db.Close()

	if rows, err2 := db.Query("select name,password,age from student"); err2 == nil {
		for rows.Next() {
			var name, password string
			var age int
			rows.Scan(&name, &password, &age)
			fmt.Println(name, password, age)
		}
	}
}

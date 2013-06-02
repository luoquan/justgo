package dataobj

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Dao struct {
}

func GetDb() *sql.DB {
	db, err := sql.Open("mysql", "root:luoquan@tcp(localhost:3306)/studygolang?charset=utf8")
	if err != nil {
		fmt.Println(err)
		panic("can not connect to mysql")
	}
	return db
}

func (dao *Dao) Delete(table string, paraMap map[string]interface{}) (re int) {
	db := GetDb()
	defer db.Close()

	var paras []interface{}
	sql := "delete from " + table
	if paraMap != nil && len(paraMap) > 0 {
		sql += " where "
		count := 0
		for key, val := range paraMap {
			count += 1
			sql += (key + "=?")
			if count < len(paraMap) {
				sql += " and "
			}
			paras = append(paras, val)
		}
	}

	if result, err := db.Exec(sql, paras...); err != nil {
		fmt.Println("execute sql err:"+sql, err)
	} else {
		temp, _:=result.RowsAffected()
		return int(temp)
	}
	return 0
}

func (dao *Dao) Insert(table string, paraMap map[string]interface{}) (re int) {
	if paraMap == nil || len(paraMap) <= 0 {
		return 0
	}

	db := GetDb()
	defer db.Close()

	var paras []interface{}
	sql := "insert into " + table + "("
	count := 0
	for key, val := range paraMap {
		count += 1
		sql += key
		if count < len(paraMap) {
			sql += ","
		}
		paras = append(paras, val)
	}

	sql += ") values ("
	count = 0
	for _, _ = range paraMap {
		count += 1
		sql += "?"
		if count < len(paraMap) {
			sql += ","
		}
	}
	sql += ")"

	if result, err := db.Exec(sql, paras...); err != nil {
		fmt.Println("execute sql err:"+sql, err)
	} else {
		temp, _:=result.RowsAffected()
		return int(temp)
	}
	return 0
}

func (dao *Dao) Update(table string, setValue map[string]interface{}, queryData map[string]interface{}) (re int) {
	if setValue == nil || len(setValue) <= 0 {
		return 0
	}

	db := GetDb()
	defer db.Close()

	var paras []interface{}
	sql := "update " + table + " set "
	count := 0
	for key, val := range setValue {
		count += 1
		sql += (key+"=?")
		if count < len(setValue) {
			sql += ","
		}
		paras = append(paras, val)
	}

	if setValue != nil && len(setValue) > 0 {
		sql += " where "
		count = 0
		for key, val := range queryData {
			count += 1
			sql += (key + "=?")
			if count < len(queryData) {
				sql += " and "
			}
			paras = append(paras, val)
		}
	}


	if result, err := db.Exec(sql, paras...); err != nil {
		fmt.Println("execute sql err:"+sql, err)
	} else {
		temp, _:=result.RowsAffected()
		return int(temp)
	}
	return 0
}

func (dao *Dao) QueryAndPaging(table string, paraMap map[string]interface{}, page, pageSize int) *sql.Rows {
	db := GetDb()
	defer db.Close()

	var paras []interface{}
	sql := "select * from " + table
	if paraMap != nil && len(paraMap) > 0 {
		queryPara := " where "
		count := 0
		for key, val := range paraMap {
			count += 1
			queryPara += " " + key + "=? "
			paras = append(paras, val)
			if count < len(paraMap) {
				queryPara += " and "
			}
		}
		sql += queryPara
	}
	if page > 0 {
		sql += " limit ?,?"
		paras = append(paras, (page-1)*pageSize, pageSize)
	}

	fmt.Println("executing query sql:" + sql)

	if stmt, err := db.Prepare(sql); err != nil {
		fmt.Println("can not prepare sql:"+sql, err)
	} else {
		if rows, err2 := stmt.Query(paras...); err2 != nil {
			fmt.Println("query error", err2)
		} else {
			return rows
		}
	}

	return nil
}

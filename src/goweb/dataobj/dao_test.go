package dataobj

import (
	"database/sql"
	"testing"
	"time"
)

func TestUpdate(t *testing.T) {
	dao := &Dao{}
	queryData := make(map[string]interface{})
	queryData["tid"] = 14


	setValue := make(map[string]interface{})
	setValue["content"] = "-------------------luoquan-------------------------"


	var rows int = dao.Update("topics", setValue, queryData)
	t.Log(rows)
}

func TestDelete(t *testing.T) {
	dao := &Dao{}
	paraMap := make(map[string]interface{})
	paraMap["tid"] = 13


	var rows int = dao.Delete("topics", paraMap)
	t.Log(rows)
}

func TestInsert(t *testing.T) {
	dao := &Dao{}
	paraMap := make(map[string]interface{})
	//paraMap["tid"] = 3
	paraMap["title"] = "luoquan"
	paraMap["content"] = "hahaluoquan hahaluoquan"
	paraMap["nid"] = 2
	paraMap["uid"] = 1
	paraMap["lastreplyuid"] = 1
	paraMap["lastreplytime"] = time.Now()
	paraMap["flag"] = 0
	paraMap["ctime"] = time.Now()
	paraMap["mtime"] = time.Now()

	var rows int = dao.Insert("topics", paraMap)
	t.Log(rows)
}

func TestQueryAndPaging(t *testing.T) {
	dao := &Dao{}
	paraMap := make(map[string]interface{})
	paraMap["tid"] = 2
	var rows *sql.Rows = dao.QueryAndPaging("topics", paraMap, 1, 10)
	t.Log("haha....", rows)
	for rows.Next() {
		topic := new(Topic)
		rows.Scan(&topic.Tid, &topic.Title, &topic.Content, &topic.Nid, &topic.Uid, &topic.Lastreplyuid,
			&topic.Lastreplytime, &topic.Flag, &topic.Ctime, &topic.Mtime)
	}
}

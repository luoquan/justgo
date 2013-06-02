package dataobj

import (
	"database/sql"
)

type Topic struct {
	Tid           int    `json:"tid"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	Nid           int    `json:"nid"`
	Uid           int    `json:"uid"`
	Lastreplyuid  int    `json:"lastreplyuid"`
	Lastreplytime string `json:"lastreplytime"`
	Flag          int    `json:"flag"`
	Ctime         string `json:"ctime"`
	Mtime         string `json:"mtime"`

	*Dao
}

const (
	TOPIC = "topics"
)

func RowsToTopic(rows *sql.Rows) (re []*Topic) {
	for rows.Next() {
		topic := new(Topic)
		rows.Scan(&topic.Tid, &topic.Title, &topic.Content, &topic.Nid, &topic.Uid, &topic.Lastreplyuid,
			&topic.Lastreplytime, &topic.Flag, &topic.Ctime, &topic.Mtime)
		re = append(re, topic)

	}
	return
}

func GetAllTops() (re []*Topic) {
	dao := &Dao{}
	re = RowsToTopic(dao.QueryAndPaging(TOPIC, nil, -1, -1))
	return
}

func GetTopicById(id int) (re *Topic) {
	dao := &Dao{}
	params := map[string] interface {} {"tid":id}
	queryRe := RowsToTopic(dao.QueryAndPaging(TOPIC, params, -1, -1))
	if queryRe == nil || len(queryRe) > 0 {
		re = queryRe[0]
	}
	return
}

func DeleteTopicById(id int) (re int) {
	dao := &Dao{}
	params := map[string] interface {} {"tid":id}
	re = dao.Delete(TOPIC, params)
	return
}

func SaveTopic(topic *Topic) (re int) {
	dao := &Dao{}
	paraMap := make(map[string]interface{})

	paraMap["title"] = topic.Title
	paraMap["content"] = topic.Content
	paraMap["nid"] = topic.Nid
	paraMap["uid"] = topic.Uid
	paraMap["lastreplyuid"] = topic.Lastreplyuid
	paraMap["lastreplytime"] = topic.Lastreplytime
	paraMap["flag"] = topic.Flag
	paraMap["ctime"] = topic.Ctime
	paraMap["mtime"] = topic.Mtime

	if topic.Tid <= 0 {
		re = dao.Insert(TOPIC, paraMap)
	} else {
		re = dao.Update(TOPIC, paraMap, map[string]interface {}{"tid": topic.Tid})
	}
	return
}

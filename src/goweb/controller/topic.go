package controller

import (
	"net/http"
	"time"
	"goweb/dataobj"
	"strconv"
)

func NewTopic(rw http.ResponseWriter, r *http.Request) {
	RedirectUtil(rw, "/template/editTopic.html", map[string]interface {}{"action":"/topics/update", "message":"点击提交保存"})
}

func DeleteTopic(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idStr := r.FormValue("id")
	if id, err := strconv.Atoi(idStr); err == nil {
		rows := dataobj.DeleteTopicById(id)
		var message string
		if rows > 0 {
			message = "删除成功"
		}   else {
			message = "删除失败"
		}
		topics := dataobj.GetAllTops()
		params := map[string]interface {}{"topics":topics, "message":message}
		RedirectUtil(rw, "/template/index.html", params)
	}

}
func ViewTopic(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idStr := r.FormValue("id")
	if id, err := strconv.Atoi(idStr); err == nil {
		topic := dataobj.GetTopicById(id)

		params := map[string]interface {}{"topic":topic, "action":"/topics/update", "message":"点击提交更新"}
		RedirectUtil(rw, "/template/editTopic.html", params)
	}

}
func UpdateTopic(rw http.ResponseWriter, r *http.Request) {
	var message string
	var id int = 0
	r.ParseForm()
	idStr := r.FormValue("Tid")
	if idre, err := strconv.Atoi(idStr); err == nil {
		message = "更新成功！！！！！！！！！！！！！！！"
		id = idre
	} else {
		message = "保存成功！！！！！！！！！！！！！！！"
	}

	TitleStr := r.FormValue("Title")
	ContentStr := r.FormValue("Content")
	NidStr := r.FormValue("Nid")
	UidStr := r.FormValue("Uid")
	LastreplyuidStr := r.FormValue("Lastreplyuid")
	LastreplytimeStr := r.FormValue("Lastreplytime")
	FlagStr := r.FormValue("Flag")
	CtimeStr := r.FormValue("Ctime")
	MtimeStr := r.FormValue("Mtime")

	topic := &dataobj.Topic{}
	topic.Tid = id
	topic.Title = TitleStr
	topic.Content = ContentStr
	topic.Nid = parseInt(NidStr)
	topic.Uid = parseInt(UidStr)
	topic.Lastreplyuid = parseInt(LastreplyuidStr)
	topic.Lastreplytime = LastreplytimeStr
	topic.Flag = parseInt(FlagStr)
	topic.Ctime = CtimeStr
	topic.Mtime = MtimeStr

	if dataobj.SaveTopic(topic) <= 0 {
		message = "操作失败"
	}


	topics := dataobj.GetAllTops()
	params := map[string]interface {}{"topics":topics, "message":message}
	RedirectUtil(rw, "/template/index.html", params)
}

func parseInt(str string) (re int) {
	re = 0
	if result, err := strconv.Atoi(str); err == nil {
		re = result
	}
	return
}

func parseTime(str string) (re time.Time) {
	re = time.Now()
	if result, err := time.Parse("yyyy-MM-dd hh:mm:ss", str); err == nil {
		re = result
	}
	return
}




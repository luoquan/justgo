/**
 * Created with IntelliJ IDEA.
 * User: luoquan
 * Date: 13-5-20
 * Time: 下午11:52
 * To change this template use File | Settings | File Templates.
 */
package dao

import (
	"fmt"
	"time"
)

//select id,version,type,title,organization_full_name,status,has_opinion,create_time
// from common_workflow_task where id<10;

var queue chan Wf = make(chan Wf, 1000)

type Wf struct {
	Id         int64
	Version    int
	Type       int
	Title      string
	Org_name   string
	Status     string
	HasOpinion bool
	CreateTime time.Time
}

func ReadRecords(start, offset int) (result []Wf) {
	sql := `select id,version,type,title,organization_full_name,status,has_opinion,create_time
	from common_workflow_task order by id asc limit ?,?`

	db := GetRDb()

	if stat, err := db.Prepare(sql); err != nil {
		if rows, err2 := stat.Query(start, offset); err2 != nil {
			for rows.Next() {
				wf := &Wf{}
				rows.Scan(wf.Id, wf.Version, wf.Type, wf.Title, wf.Org_name, wf.HasOpinion, wf.CreateTime)
				result = append(result, *wf)
			}
		} else {
			fmt.Println("error occured when query")
		}
	} else {
		fmt.Println("error occured when prepare")
	}
	return result
}

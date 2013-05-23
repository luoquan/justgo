package dao

import (
	"fmt"
	"database/sql"
	"strconv"
)

func WriteRecords(db *sql.DB, arr []*Wf) {
	if arr == nil || len(arr) == 0 {
		return
	}

	sql := "insert into wf values(?,?,?,?,?,?)"

	if stat, err := db.Prepare(sql); err == nil {
		for _, val := range arr {
			//fmt.Println(val.Title, val.Org_name, val.Status)
			if _, err2 := stat.Exec(val.Id, val.Version, val.Type, val.Title, val.Org_name, val.Status); err2 == nil {
//				rowsAffected, _ := re.RowsAffected()
//				lastInsertId, _ := re.LastInsertId()
				//fmt.Println("rows_affected:", rowsAffected , "last_insert_id:", lastInsertId)
			}else {
				fmt.Println("error occured when insert", err2)
			}
		}
		defer stat.Close();
	}   else {
		fmt.Println("error occured when prepare", err)
	}
}

func ReadRecords(db *sql.DB, start, offset int) (result []*Wf) {
	sql := `select id,version,type,title,organization_full_name,status
	from common_workflow_task order by id asc limit ?,?`

	fmt.Println("query data start:", strconv.Itoa(start), " offset:", strconv.Itoa(offset))
	if stat, err := db.Prepare(sql); err == nil {
		if rows, err2 := stat.Query(start, offset); err2 == nil {
			for rows.Next() {
				var id, version, ty int
				var title, org, status string
				rows.Scan(&id, &version, &ty, &title, &org, &status)
				//fmt.Println(id, version, ty, title, org, status)

				wf := new(Wf)
				wf.Id = id
				wf.Version = version
				wf.Type = ty
				wf.Title = title
				wf.Org_name = org
				wf.Status = status
				result = append(result, wf)
			}
			defer rows.Close()
		} else {
			fmt.Println("error occured when query", err2)
		}
		defer stat.Close()
	} else {
		fmt.Println("error occured when prepare", err)
	}
	return result
}

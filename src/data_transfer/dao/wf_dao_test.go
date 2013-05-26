package dao

import (
	"testing"
	"strconv"
)

func TestRead(t *testing.T) {
	var count int = 1
	arr := ReadRecords(GetRDb(), 0, count)
	if len(arr) != count {
		t.Errorf("result size not right:",strconv.Itoa(len(arr)))
	}
}

func TestWrite(t *testing.T) {
	wf := new(Wf)
	wf.Version = 2
	wf.Type = 3
	wf.Title = "HR"
	wf.Org_name = "北京"
	wf.Status = "complete"
	var arr []*Wf
	arr = append(arr, wf)

	WriteRecords(GetWDb(), arr)
}

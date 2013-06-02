package dataobj

import (
	"testing"
	"fmt"
)

func TestGetAllTopics(t *testing.T) {
	 arr := GetAllTops()
	t.Error(len(arr))
	for _,val := range arr{
		fmt.Println(val.Content)
	}

	t.Error("end")
}


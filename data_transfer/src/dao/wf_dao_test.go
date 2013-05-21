/**
 * Created with IntelliJ IDEA.
 * User: luoquan
 * Date: 13-5-21
 * Time: 下午8:14
 * To change this template use File | Settings | File Templates.
 */
package dao

import (
	"testing"
	"strconv"
	"fmt"
)

func TestRead(t *testing.T) {
	var count int = 10
	arr := ReadRecords(0, count)
	t.Errorf(strconv.Itoa(len(arr)))
	if false {
		t.Errorf("result size not right")
	} else {
		for _, val := range arr {
			fmt.Println(val.Id, val.Title)
		}
	}
}

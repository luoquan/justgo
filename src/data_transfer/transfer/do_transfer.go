package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"dao"
	"runtime"
)

const (
	NUM_OF_READER int = 1
	NUM_OF_WRITER int = 10
	QUEUE_SIZE  int   = 1000
)

var POISON *dao.Wf = &dao.Wf{Id:-1}

func write(queue chan *dao.Wf, finishFlag chan int) {
	db := dao.GetWDb()
	defer db.Close()
	buffer := 10
	finish := false
	for !finish {
		var arr []*dao.Wf
		for {
			wf := <-queue
			if (wf == POISON) {
				queue <- POISON
				finish = true
				break
			}
			arr = append(arr, wf)
			if len(arr) == buffer {
				break
			}
		}
		if len(arr) > 0 {
			dao.WriteRecords(db, arr)
		}
	}
	finishFlag <- 1
}

func read(queue chan *dao.Wf, finishFlag chan int) {
	db := dao.GetRDb()
	defer db.Close()
	start := 0
	for {
		wfArr := dao.ReadRecords(db, start, QUEUE_SIZE)
		if wfArr == nil {
			queue<-POISON
			break
		}        else {
			for _, wf := range wfArr {
				queue <-  wf
			}
			start += QUEUE_SIZE
		}
	}
	finishFlag <-1
}


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	queue := make(chan *dao.Wf, QUEUE_SIZE)
	finishFlag := make(chan int, NUM_OF_WRITER + 1)

	fmt.Println("starting transfer...")
	go read(queue, finishFlag)
	for i := 0; i < NUM_OF_WRITER; i++ {
		go write(queue, finishFlag)
	}
	fmt.Println("transfer finished...")

	for i := 0; i < NUM_OF_WRITER + 1; i++ {
		<-finishFlag
	}
}

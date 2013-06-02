package main

import (
	"net/http"
	"fmt"
	"goweb/controller"
)

type MyHandler struct {
}

func (h *MyHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	http.Redirect(rw, r, "/goweb/index.htm", 200)
}

func main() {
//		var h MyHandler
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/topics/new", controller.NewTopic)
	http.HandleFunc("/topics/view", controller.ViewTopic)
	http.HandleFunc("/topics/delete", controller.DeleteTopic)
	http.HandleFunc("/topics/update", controller.UpdateTopic)
		if err := http.ListenAndServe("localhost:8080", nil); err != nil {
			fmt.Println("Start server error...")
		}
}

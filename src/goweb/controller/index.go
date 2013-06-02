package controller

import (
	"net/http"
	"fmt"
	"goweb/dataobj"
)

func Index(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("executing index...", r.RequestURI)
	topics := dataobj.GetAllTops()
	RedirectUtil(rw, "/template/index.html", map[string]interface {}{"topics":topics})
}


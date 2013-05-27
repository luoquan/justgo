/*
this is commet
 */
package main

import (
	"net/http"
	"fmt"
	"math/rand"
)

type String string

func (h String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is String handler:" + h)
}

type Struct struct {
	Who string
}

/* handler of type Struct
//handle path:/struct
*/
func (h Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is Struct handler" + h.Who)
}

func main() {
	var sh String = "luoquan's String"
	http.Handle("/string", sh)
	var sth Struct = Struct{"this is 罗全"}
	http.Handle("/struct", sth)

	re:=0
	for  re<1{

	p:=rand.Perm(10)
	fmt.Println(p)
		re+=1
	}

	//http.ListenAndServe("localhost:8080", nil)
}


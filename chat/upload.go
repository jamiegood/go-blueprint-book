package main

import (
	"fmt"
	"io"
	"net/http"
)

// uploaderHandler ...
func uploaderHandler(w http.ResponseWriter, r *http.Request) {

	userid := r.FormValue("userid")

	fmt.Println(userid)
	io.WriteString(w, "hello")

}

package api

import "net/http"

func HomeIndex(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(">>>>>>"))
}


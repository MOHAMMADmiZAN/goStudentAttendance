package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	route := httprouter.New()
	route.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	})

}

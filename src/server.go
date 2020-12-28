package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(":3001", router))
}

func Index(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	_, _ = writer.Write([]byte("Hello"))
}

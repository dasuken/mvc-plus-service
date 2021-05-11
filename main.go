package main

import (
	"github.com/noguchidaisuke/mvc-plus-service/controller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controller.Handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

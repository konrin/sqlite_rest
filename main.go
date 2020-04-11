package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	if isRequestVersion() {
		fmt.Printf("v%s", Version)
		return
	}

	conf := GetConfig()

	db, err := NewDB(conf.DB.Connection)
	if err != nil {
		log.Fatal(err)
	}

	service := NewService(db)
	rest := NewRest(service)

	http.HandleFunc("/sql/exec", rest.ExecHandler())
	http.HandleFunc("/sql/query", rest.QueryHandler)

	http.ListenAndServe(":1212", nil)
}
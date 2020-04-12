package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	if isRequestVersion() {
		fmt.Printf("v%s\n", Version)
		return
	}

	conf := GetConfig()

	db, err := NewDB(conf.DB.Connection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	service := NewService(db)
	rest := NewRest(service)

	http.HandleFunc("/sql/exec", rest.ExecHandler())
	http.HandleFunc("/sql/query", rest.QueryHandler)

	http.ListenAndServe(fmt.Sprintf("%s:%d", conf.HTTP.Host, conf.HTTP.Port), nil)
}
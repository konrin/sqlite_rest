package main

import (
	"fmt"
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
		fmt.Printf("DB init error: %s\n", err.Error())
		return
	}
	defer db.Close()

	service := NewService(db)
	rest := NewRest(service)

	http.HandleFunc("/sql/exec", rest.ExecHandler())
	http.HandleFunc("/sql/query", rest.QueryHandler)

	if err = http.ListenAndServe(fmt.Sprintf("%s:%d", conf.HTTP.Host, conf.HTTP.Port), nil); err != nil {
		fmt.Printf("Server error: %s\n", err.Error())
	}
}

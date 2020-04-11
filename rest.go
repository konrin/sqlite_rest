package main

import (
	"encoding/json"
	"net/http"
)

type Rest struct {
	service *Service
}

func NewRest(service *Service) *Rest {
	return &Rest{service}
}

func (r *Rest) ExecHandler() func(http.ResponseWriter, *http.Request) {
	type bodyData struct {
		Sql    string
		Params []interface{}
	}

	return func(w http.ResponseWriter, req *http.Request) {
		var data bodyData

		if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
			r.sendError(w, 500, "QUERY_ERR", err.Error())
			return
		}

		lastInsertId, rowsAffected, err := r.service.ExecRaw(data.Sql, data.Params...)
		if err != nil {
			r.sendError(w, 500, "QUERY_ERR", err.Error())
			return
		}

		r.sendOK(w, map[string]interface{}{
			"lastInsertId": lastInsertId,
			"rowsAffected": rowsAffected,
		})
	}
}

func (r *Rest) QueryHandler(w http.ResponseWriter, req *http.Request) {
	sql := req.URL.Query().Get("sql")
	paramsRaw := req.URL.Query().Get("params")
	params := make([]interface{}, 0)

	if len(paramsRaw) > 0 {
		if err := json.Unmarshal([]byte(paramsRaw), &params); err != nil {
			r.sendError(w, 500, "QUERY_ERR", err.Error())
			return
		}
	}

	rows, err := r.service.QueryRaw(sql, params...)
	if err != nil {
		r.sendError(w, 500, "QUERY_ERR", err.Error())
		return
	}

	r.sendOK(w, r.service.RowsToMap(rows))
}

func (r *Rest) sendError(w http.ResponseWriter, httpCode int, code string, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    code,
		"message": message,
	})
}

func (r *Rest) sendOK(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(data)
}

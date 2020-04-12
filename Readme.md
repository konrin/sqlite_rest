# SQLite rest server

###### REST API

`POST /sql/exec`  
Body JSON params: sql, params  
Return JSON object with _lastInsertId_ and _rowsAffected_ 

`GET /sql/query`  
Query params: sql, params  
Return array query result list

###### Sample usage

Build  
`go build -ldflags "-s -w"`

Run  
`./sqlite_rest -c :memory:`
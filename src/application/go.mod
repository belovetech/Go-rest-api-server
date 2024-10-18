module github.com/belovetech/Go-rest-api-server.git/application

go 1.20

replace rest-api.com/backend => ../backend

require rest-api.com/backend v0.0.0-00010101000000-000000000000

require (
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
)

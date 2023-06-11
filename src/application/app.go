package main

import (
	"rest-api.com/backend"
)

func main() {
	app := backend.App{}
	app.Port = ":3003"
	app.Initialize()
	app.Run()
}

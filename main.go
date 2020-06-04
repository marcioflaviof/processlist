package main

import (
	"processlist/database"
	"processlist/server"
)

func main() {

	database.CreateClient()

	server.Start()
}

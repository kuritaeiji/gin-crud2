package main

import (
	"gin-crud2/db"
	"gin-crud2/server"
)

func main() {
	db.Init()
	defer db.CloseDB()

	server.Init()
}

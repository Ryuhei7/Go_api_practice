package main

import (
	"ryuhei/ex_api/db"
	"ryuhei/ex_api/server"
)

func main() {

	db.Init()
	server.Init()

}

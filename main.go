package main

import (
	"github.com/xavimoreno548/BookStore/db"
	"github.com/xavimoreno548/BookStore/router"
)

func main() {

	var port = "9000"

	db.InitDB()
	router.InitRouter(port)
}
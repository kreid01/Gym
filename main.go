package main

import (
	"github.com/kreid01/gym/db"
	"github.com/kreid01/gym/router"
)

func main() {
	db.InitPostgresDB()
	router.InitRouter().Run()
}

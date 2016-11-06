package main

import (
	"log"
	"os"

	"github.com/ironecally/domino-api/apiHandler"
	"github.com/ironecally/domino-api/database"
)

var dbModule *database.DB

func init() {

	db, err := database.InitDB()
	if err != nil {
		os.Exit(1)
		return
	}
	dbModule = db

	apiHandler.InitAPI()

	return
}

func main() {
	log.Println(dbModule)
}

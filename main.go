package main

import (
	"os"

	"github.com/ironecally/domino-api/database"
)

var dbModule database.DB

func init() {
	dbModule, err := database.InitDB()
	if err != nil {
		os.Exit(1)
		return
	}
	dbModule.DBConn.Ping()

	return
}

func main() {

}

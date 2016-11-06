package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// MysqlConfig will be filled by the config's data
type MysqlConfig struct {
	Username  string
	Password  string
	DBName    string
	ParseTime string
	Loc       string
}

// DB will be returned as result of this package
type DB struct {
	DBConn *sql.DB
}

// InitDB should be called first to initialize the db module
func InitDB() (*DB, error) {
	dbModule := new(DB)

	// for now hardcode the config
	config := MysqlConfig{
		Username:  "root",
		Password:  "",
		DBName:    "domino",
		ParseTime: "true",
		Loc:       "Local",
	}
	dsn := fmt.Sprintf(`%s:%s@/%s?parseTime=%s&loc=%s`, config.Username, config.Password, config.DBName, config.ParseTime, config.Loc)
	dbConnection, err := sql.Open("mysql", dsn)
	log.Println(dbConnection)
	if err != nil {
		log.Println("DB Error", err.Error())
		log.Println("Failed to Open Connection:", dsn)
		return dbModule, err
	}

	err = dbConnection.Ping()
	if err != nil {
		log.Println("Failed to Ping DB", err.Error())
		return dbModule, err
	}

	dbModule.DBConn = dbConnection

	return dbModule, nil
}

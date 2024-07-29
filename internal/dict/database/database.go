package dict_db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func New() *sql.DB {

	dictConfig := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "dictdb",
		AllowNativePasswords: true,
	}
	dictDB, err := sql.Open("mysql", dictConfig.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return dictDB

}

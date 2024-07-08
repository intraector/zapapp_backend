package zap_db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func ZapDB() *sql.DB {

	zapConfig := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "zapdb",
		AllowNativePasswords: true,
	}
	zapDB, err := sql.Open("mysql", zapConfig.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return zapDB

}

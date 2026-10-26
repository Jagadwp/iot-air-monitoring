package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Koneksi database yang dipakai di semua handler
var DB *sql.DB

func ConnectDB() {
	// Format: user:password@tcp(host:port)/nama_database
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/iot_monitoring")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DB = db
}

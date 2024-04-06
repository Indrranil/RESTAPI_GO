package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

// implementation for my sql
type MySQLStorage struct {
	db *sql.DB
}

// constructor
// (cfg mysql.Config) -> it will receive a config from mysql.config
// and it return a MySQLStorage
func NewMYSQLStorage(cfg mysql.Config) *MySQLStorage {

	// here initilize the database and make connection

	db, err := sql.Open("mysql", cfg.FormatDSN())
	
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MySQL")

	return &MySQLStorage{db:db}

}

// initilize the database tables and return them 
func (s *MySQLStorage) Init() (*sql.DB, error){
	return s.db, nil 
}

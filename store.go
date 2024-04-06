package main

import "database/sql"

type Store interface {
	// Users
	createUser() error
}

// methods to communicate with the database
type Storage struct {
	// here it will receive the dependency of the communication tool (sql)
	db *sql.DB
}

// constructor
func NewStore(db *sql.DB) *Storage { // sql.db dependencies will be stored in the *So
	return &Storage{
		db: db,
	}
}

// Describes the createUser() here
// method for store, this method will implement the store interface above
func (s *Storage) createUser() error {
	return nil
}

package main

import "database/sql"

// Message holds messages
type Message struct {
	Body string
}

// Database holds the store
type Database struct {
	db sql.DB
}

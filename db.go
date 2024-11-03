// Package main demonstrates basic SQLite database operations in Go
// including database creation, table initialization, and proper resource management.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

// Database configuration constants
const (
    // DbPath specifie the directory where the database file will b estored
    DbPath = "priv"
    // DbFile specifies the name of the SQLite database file
    DbFile = "sqlite.db"
    // CreateTableSQL defines the SQL statement for creatint eh attendees table
    CreateTableSQL = '
        CREATE TABLE IF NOT EXIST attendees (
            id INTEGER PRIMARY KEY,
            name VARCHAR NOT NULL,
            age INTEGER CHECK (age >= 0)
        )' 
)

// Attendee represents a person in the attendees table
type Attendee struct {
    ID int64
    Name string
    Age string
}

// DatabaseManager handles database operations and connections
type DatabaseManager struct {
    db *sql.DB 
    path string
}

// New

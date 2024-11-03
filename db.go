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

// NewDatabaseManager creates and initializes a new database manager
func NewDatabaseManager(dbPath string) (*DatabaseManager, error) {
	// Ensure database directory exists
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}

	// Open database connection
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		db.Close() // Clean up on error
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &DatabaseManager{
		db:   db,
		path: dbPath,
	}, nil
}

// Close closes the database connection
func (dm *DatabaseManager) Close() error {
	if dm.db != nil {
		return dm.db.Close()
	}
	return nil
}

// InitializeTable creates the attendees table if it doesn't exist
func (dm *DatabaseManager) InitializeTable() error {
    // Prepare the create table statement
    stmt, err := dm.db.Prepare(CreateTableSQL)
    if err != nil {
        return fmt.Errorf("failed to prepare create table statement: %w", err)
    }
    defer stmt.Close()

    	// Execute the create table statement
	if _, err := stmt.Exec(); err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}

// ResetDatabase removes the existing database file and creates a new one
func ResetDatabase(dbPath string) error {
	// Remove existing database file
	if err := os.Remove(dbPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove existing database: %w", err)
	}
	return nil
}

func main() {
	// Construct database path
	dbPath := filepath.Join(DbPath, DbFile)

	// Reset database (remove existing file)
	if err := ResetDatabase(dbPath); err != nil {
		log.Fatalf("Failed to reset database: %v", err)
	}

	// Initialize database manager
	dbManager, err := NewDatabaseManager(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database manager: %v", err)
	}
	defer dbManager.Close() // Ensure database connection is closed

	// Create attendees table
	if err := dbManager.InitializeTable(); err != nil {
		log.Fatalf("Failed to initialize table: %v", err)
	}

	log.Println("Successfully created attendees table")
}
//go:build ignore
// +build ignore

package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect without database name first to create it
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/?multiStatements=true")
	if err != nil {
		log.Fatal("Connection error:", err)
	}
	defer db.Close()

	// Create database
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS trangleagent CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;")
	if err != nil {
		log.Fatal("Create DB error:", err)
	}

	// Connect to the specific database
	db, err = sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/trangleagent?multiStatements=true")
	if err != nil {
		log.Fatal("Connection to DB error:", err)
	}
	defer db.Close()

	// Execute 000_init_schema.sql
	executeFile(db, "manifest/sql/000_init_schema.sql")
	// Execute 004_refactor.sql
	executeFile(db, "manifest/sql/004_refactor.sql")

	fmt.Println("Database initialized successfully!")
}

func executeFile(db *sql.DB, path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Read file %s error: %v", path, err)
	}

	requests := strings.Split(string(content), ";")
	for _, request := range requests {
		request = strings.TrimSpace(request)
		if request == "" || strings.HasPrefix(request, "--") {
			continue
		}
		_, err := db.Exec(request)
		if err != nil {
			log.Printf("Execute SQL error in %s: %v\nQuery: %s", path, err, request)
		}
	}
}

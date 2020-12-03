package main

import (
	"fmt"
	"log"
	"time"

	"database/sql" // This package is for querying all sorts of SQL databases

	_ "github.com/go-sql-driver/mysql" // The database driver is specifically for MySQL
)

func main() {
	// Configure the database connection (always check errors)
	db, err := sql.Open("mysql", "root:password@(127.0.0.1:8080)/godb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	{
		// mysql commands
		// Ex. creating a table
		query := `
        CREATE TABLE users (
            id INT AUTO_INCREMENT,
            username TEXT NOT NULL,
            password TEXT NOT NULL,
            created_at DATETIME,
            PRIMARY KEY (id)
        );`
		// Executes the SQL query in our database. Check err to ensure there was no error.
		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{
		// by default, Go uses prepared statements for inserting dynamic data into SQL queries
		// Ex: INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)
		username := "johndoe"
		password := "secret"
		createdAt := time.Now()

		// Inserts our data into the users table and returns with the result and a possible error.
		// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		// I can grab the newly created id from
		userID, err := result.LastInsertId()
		fmt.Println(userID)
	}

	{
		// For querying a single row of data, we first need to declare variables to store the data
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)
		// Query the database and scan the values into out variables. Don't forget to check for errors.
		query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)
	}

	{
		// For querying multiple rows
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}
		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`) // check err
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user
		for rows.Next() {
			var u user
			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt) // check err
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%#v", users)
	}

	{
		// To delete data from the tables
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1) // check err
		if err != nil {
			log.Fatal(err)
		}
	}

}

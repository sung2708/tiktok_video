package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/root?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	//	{
	//		query := `
	//			CREATE TABLE users (
	//				id INT AUTO_INCREMENT,
	//				username TEXT NOT NULL,
	//				password TEXT NOT NULL,
	//				created_at DATETIME,
	//				PRIMARY KEY (id)
	//			);`
	//
	//		if _, err := db.Exec(query); err != nil {
	//			log.Fatal(err)
	//		}
	//	}

	{
		username := "Phan Duc Sung"
		password := "Sungp2708@"
		created_at := time.Now()

		result, err := db.Exec(
			`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`,
			username,
			password,
			created_at,
		)

		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()

		fmt.Println(id)
	}
	{
		var (
			id         int
			username   string
			password   string
			created_at time.Time
		)

		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"

		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &created_at); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, created_at)
	}

	{
		type User struct {
			id         int
			username   string
			password   string
			created_at time.Time
		}

		row, err := db.Query(`SELECT id, username, password, created_at FROM users`)

		if err != nil {
			log.Fatal(err)
		}
		defer row.Close()

		var users []User

		for row.Next() {
			var u User

			err := row.Scan(&u.id, &u.username, &u.password, &u.created_at)

			if err != nil {
				log.Fatal(err)
			}

			users = append(users, u)
		}
		if err := row.Err(); err != nil {
			log.Fatal(err)
		}
	}
	{
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}

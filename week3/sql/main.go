package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	ID           int64
	Name         string
	Email        string
	Password     string
	RegisteredAt time.Time
}

func main() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres sslmode=disable password=goLANGninja")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("CONNECTED")

	rows, err := db.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisteredAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
}

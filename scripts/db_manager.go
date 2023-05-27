package main

import (
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

/// DB MANAGER
// A script for databases creation and fill

const ordersPath string = "internal/orders/database/orders.db"
const usersPath string = "internal/users/database/users.db"

func ordersManage() error {
	// Remove DataBase
	if err := os.Remove(ordersPath); err != nil {
        return err
    }

	// Create DataBase
	file, err := os.Create(ordersPath)
	if err != nil {
		return err
	}
	file.Close()
	log.Println("Orders DataBase created successfully!")

	// Connect to DataBase
	db, err := sql.Open("sqlite3", ordersPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create Tables
	table := `CREATE TABLE orders (
				orderID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				userID INTEGER NOT NULL,
				status INTEGER,
				date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				coffees TEXT,
				total REAL
			);`
	if _, err := db.Exec(table); err != nil {
		return err
	}
	log.Println("Table created successfully!")

	// Fill tables
	records := `INSERT INTO orders (
					userID,
					status,
					coffees,
					total
				) VALUES (?, ?, ?, ?);`
	res, err := db.Exec(records, 1, 0, `[{"Type": "Espresso", "Sugar": 2}, {"Type": "Americano", "Sugar": 1}]`, 5.5)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	log.Printf("Table filled successfully with %v rows!\n", rows)
	return nil
}

func usersManage() error {
	// Remove DataBase
	if err := os.Remove(usersPath); err != nil {
        return err
    }

	// Create DataBase
	file, err := os.Create(usersPath)
	if err != nil {
		return err
	}
	file.Close()
	log.Println("User DataBase created successfully!")

	// Connect to DataBase
	db, err := sql.Open("sqlite3", usersPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create Tables
	table := `CREATE TABLE users (
				userID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				username TEXT NOT NULL,
				address TEXT NOT NULL,
				passwordHash TEXT NOT NULL,
				regDate TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
			);`
	if _, err := db.Exec(table); err != nil {
		return err
	}
	log.Println("Table created successfully!")

	// Fill tables
	records := `INSERT INTO users (
					username,
					address,
					passwordHash
				) VALUES (?, ?, ?);`
	hasher := sha1.New()
	hasher.Write([]byte("testPassword"))
	res, err := db.Exec(records, "Test Name", "Test City, 95 st.", hex.EncodeToString(hasher.Sum(nil)))
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	log.Printf("Table filled successfully with %v rows!\n", rows)
	return nil
}

func main() {
	if err := ordersManage(); err != nil {
		log.Printf("%v", err)
	}
	if err := usersManage(); err != nil {
		log.Printf("%v", err)
	}
}
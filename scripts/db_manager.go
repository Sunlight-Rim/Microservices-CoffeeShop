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
// A script for creating and filling a databases

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
	createTables :=
		`CREATE TABLE coffee (
			coffeeID INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(60) UNIQUE NOT NULL,
			price REAL
		);
		CREATE TABLE topping (
			toppingID INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(60) UNIQUE NOT NULL,
			price REAL
		);
		CREATE TABLE order_ (
			orderID INTEGER PRIMARY KEY AUTOINCREMENT,
			userID INTEGER NOT NULL,
			coffeeID INTEGER NOT NULL,
			toppingID INTEGER NOT NULL,
			sugar INTEGER DEFAULT 1,
			status INTEGER DEFAULT 0,
			date TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
			FOREIGN KEY (coffeeID) REFERENCES coffee(coffeeID),
			FOREIGN KEY (toppingID) REFERENCES topping(toppingID)
		);`
	if _, err := db.Exec(createTables); err != nil {
		return err
	}
	log.Println("Orders tables created successfully!")

	// Fill tables
	insertRecords :=
		`INSERT INTO coffee(name, price) VALUES ('Espresso', 3.0);
		INSERT INTO coffee(name, price) VALUES ('Americano', 3.5);
		INSERT INTO topping(name, price) VALUES ('', 0.0);
		INSERT INTO topping(name, price) VALUES ('Banana', 0.5);
		INSERT INTO topping(name, price) VALUES ('Strawberry', 0.5);
		
		INSERT INTO order_(userID, coffeeID, toppingID, sugar, status) VALUES (1, 1, 1, 2, 0);
		INSERT INTO order_(userID, coffeeID, toppingID) VALUES (1, 2, 1);`
	res, err := db.Exec(insertRecords)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()

	log.Printf("Orders tables filled successfully with %v rows!\n", rows)
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
	log.Println("User database created successfully!")

	// Connect to DataBase
	db, err := sql.Open("sqlite3", usersPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create Tables
	createTables :=
		`CREATE TABLE user (
			userID INTEGER PRIMARY KEY AUTOINCREMENT,
			username VARCHAR(30) NOT NULL,
			address VARCHAR(150) NOT NULL,
			passwordHash NVARCHAR(100) NOT NULL,
			date TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
		);`
	if _, err := db.Exec(createTables); err != nil {
		return err
	}
	log.Println("User tables created successfully!")

	// Fill tables
	insertRecords :=
		`INSERT INTO user (
			username,
			address,
			passwordHash
		) VALUES (?, ?, ?);`
	hasher := sha1.New()
	hasher.Write([]byte("testpass"))
	res, err := db.Exec(insertRecords, "testname", "Test City, Test st.", hex.EncodeToString(hasher.Sum(nil)))
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	log.Printf("User tables filled successfully with %v rows!\n", rows)
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
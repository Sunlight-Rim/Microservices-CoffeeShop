package main

import (
	configuration "coffeeshop/config"
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

/// DB MANAGER
// A script for initial creation and fill a databases

// ORDERS DB

func ordersManage(dbPath string) error {
	// Remove db
	if err := os.Remove(dbPath); err != nil {
		log.Println("Orders DB already no exists")
	}

	// Create db
	file, err := os.Create(dbPath)
	if err != nil {
		return err
	}
	file.Close()
	log.Println("Orders DB created successfully!")

	// Connect to db
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create tables
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

// USERS DB

func usersManage(dbPath string) error {
	// Remove db
	if err := os.Remove(dbPath); err != nil {
		log.Println("Users DB already no exists")
	}

	// Create db
	file, err := os.Create(dbPath)
	if err != nil {
		return err
	}
	file.Close()
	log.Println("Users DB created successfully!")

	// Connect to db
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create tables
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
	log.Println("Users tables created successfully!")

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
	log.Printf("Users tables filled successfully with %v rows!\n", rows)
	return nil
}

const configPath = "config/config.yaml"

func main() {
	config := configuration.New(configPath)

	if err := ordersManage(config.Services["orders"].DB); err != nil {
		log.Println(err)
	}
	if err := usersManage(config.Services["users"].DB); err != nil {
		log.Println(err)
	}
}

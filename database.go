package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/**
 * Connect to MySQL
 */
func DBConnect() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		config.DB.User,
		config.DB.Pass,
		config.DB.Host,
		config.DB.Schema,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	database = db
}

/**
 * Test function to see if database connection
 * actually works
 */
func DBGetTime() {
	query, err := database.Query("SELECT NOW()")
	if err != nil {
		log.Println(err)
		return
	}
	defer query.Close()

	var time string
	for query.Next() {
		err := query.Scan(&time)
		if err != nil {
			return
		}

		log.Printf("%s", time)
	}
}

/**
 * Get a Unix timestamp compatible with inserting
 * into the db
 */
func DBNow() int {
	return int(time.Now().Unix())
}

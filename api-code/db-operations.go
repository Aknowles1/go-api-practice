package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() *sql.DB {
	fmt.Printf("Connecting DB...\n")
	db, err := sql.Open("mysql", "root:secure@tcp(mysql.mysql.svc.cluster.local:3306)/")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InitialiseDB() {

	db := connectDB()
	fmt.Printf("initialising DB...\n")
	insert, err := db.Query("CREATE DATABASE IF NOT EXISTS HITCOUNT")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("USE HITCOUNT")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS CountAPIHits (counter VARCHAR(255), hits int)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO CountAPIHits(counter, hits) SELECT * FROM (SELECT 'counter1' as counter, 0 as hits) AS new_value WHERE NOT EXISTS (SELECT counter FROM CountAPIHits WHERE counter = 'counter1') LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}

	insert.Close()
	db.Close()
}

func (a *application) IncrementHitCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		message := fmt.Sprintf("unsupported method %v", r.Method)
		http.Error(w, message, http.StatusMethodNotAllowed)
		//exit function if the method (i.e r.Method, wasnt "GET")
		return
	}

	fmt.Println("Incrementing Record.........")
	db := connectDB()

	_, err := db.Exec("USE HITCOUNT")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("UPDATE CountAPIHits SET hits = hits + 1 WHERE counter = 'counter1'")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("SELECT hits FROM CountAPIHits where counter = 'counter1'")
	if err != nil {
		log.Fatal(err)
	}

	result := db.QueryRow("SELECT *  FROM CountAPIHits;")

	var count string
	var value int
	if err := result.Scan(&count, &value); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("counter: %d\n", value)
	if err := result.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%d", value)
	db.Close()
}

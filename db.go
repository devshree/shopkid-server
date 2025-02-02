package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func initDB() *sql.DB {
	log.Printf("Connecting to database with settings:")
	log.Printf("DB_HOST: %s", os.Getenv("DB_HOST"))
	log.Printf("DB_PORT: %s", os.Getenv("DB_PORT")) 
	log.Printf("DB_USER: %s", os.Getenv("DB_USER"))
	log.Printf("DB_NAME: %s", os.Getenv("DB_NAME"))
 

    dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    log.Printf("DB URL: %s", dbURL)
    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("DB connected")
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("DB pinged")
    return db
} 

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Wolf-111/Identity-Server/messageHandler"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Establish connection to PostgreSQL server
func establishDBConnection() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		host, _     = os.LookupEnv("HOST")
		port        = 5432
		user, _     = os.LookupEnv("USERNAME")
		password, _ = os.LookupEnv("PASSWORD")
		dbname, _   = os.LookupEnv("DBNAME")
	)

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// twilioErr := messageHandler.SendSMS("+", "+", "Helloooo!")
	// if twilioErr != nil {
	// 	fmt.Println(err)
	// }
	http.HandleFunc("/webhook", messageHandler.WebhookHandler)
	webhookErr := http.ListenAndServe(":8080", nil)

	if webhookErr != nil {
		fmt.Println("ListenAndServe error:", webhookErr)
	}

	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

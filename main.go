package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi/v5"
)

var db *sql.DB

func init() {
	dbHost := os.Getenv("DB_HOST_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PWD")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s slmode=verify-full sslrootcert=./certs/ca.crt", dbHost, dbPort, dbUser, dbPwd, dbName)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("DB Connection error : %s", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to ping DB: %s", err)
	}

	log.Println("Successfully connected to DB!")
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	appListenPort := os.Getenv("APP_LISTEN_PORT")

	defer db.Close()

	r := chi.NewRouter()
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		// ctx := context.WithValue(r.Context(), "x-request-id", uuid.New().String())
		// err := doDBOperation(ctx)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		w.Write([]byte(""))
	})
	http.ListenAndServe(":"+appListenPort, r)
}

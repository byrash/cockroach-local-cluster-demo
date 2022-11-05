package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/google/uuid"
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

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func doDBOperation(ctx context.Context) error {
	log.SetPrefix(fmt.Sprintf("%s ", ctx.Value("x-request-id")))
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	parentID := 0
	createParent := fmt.Sprintf("INSERT INTO parent(name) VALUES ('%s') RETURNING id", randSeq(10))
	err = db.QueryRowContext(ctx, createParent).Scan(&parentID)
	if err != nil {
		log.Println(err.Error())
		err = tx.Rollback()
		if err != nil {
			log.Println(err.Error())
		}
		return err
	}
	childID := 0
	createChild := fmt.Sprintf("INSERT INTO child(name, parent_id) VALUES ('%s', %v) RETURNING id", randSeq(10), parentID)
	err = db.QueryRowContext(ctx, createChild).Scan(&childID)
	if err != nil {
		log.Println(err.Error())
		err = tx.Rollback()
		if err != nil {
			log.Println(err.Error())
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Printf("Inserted Parent [%d] and Child [%d]\n", parentID, childID)
	return nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	appListenPort := os.Getenv("APP_LISTEN_PORT")
	appName := os.Getenv("APP_NAME")

	defer db.Close()

	r := chi.NewRouter()
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "x-request-id", uuid.New().String())
		err := doDBOperation(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Hello, Boss!!! from " + appName))
	})
	http.ListenAndServe(":"+appListenPort, r)
}

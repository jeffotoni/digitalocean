package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/jeffotoni/gconcat"
)

var (
	once    sync.Once
	err     error
	dbLocal *sql.DB

	database = os.Getenv("DB_NAME")
	host     = os.Getenv("DB_HOST")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	port     = os.Getenv("DB_PORT")
	ssl      = "require"
	source   = "postgres"

	httpPort = ":8080"
)

func DbConnect() *sql.DB {
	once.Do(func() {
		if dbLocal != nil {
			return
		}
		connStr := gconcat.Build("host=", host, " port=", port,
			" user=", user, " password=", password, " dbname=", database,
			" sslmode=", ssl)
		if dbLocal, err = sql.Open("postgres", connStr); err != nil {
			if err != nil {
				log.Println(err)
			}
			dbLocal = nil
			return
		}
	})
	return dbLocal
}

func main() {

	http.HandleFunc("/api/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello man!!!"))
	})

	println("Run Server", httpPort)
	http.ListenAndServe(httpPort, nil)

}

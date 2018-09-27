package main

import (
	"database/sql"
	"log"
	"os"

	"bitbucket.org/boomstarternetwork/minerserver/handler"
	"bitbucket.org/boomstarternetwork/minerserver/store"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

// Environment variable like
// `"postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"`
// or
// `user=pqgotest password=password dbname=pqgotest sslmode=verify-full`
const connStrEnv = "POSTGRES_CONNECTION_STRING"

func main() {
	e := echo.New()

	connStr, exists := os.LookupEnv(connStrEnv)
	if !exists {
		e.Logger.Fatalf("%s environment variable haven't set", connStrEnv)
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ps := store.NewDBProjectsStore(db)

	h := handler.NewHandler(ps)

	e.GET("/projects/list", h.ProjectsList)

	e.Logger.Fatal(e.Start(":80"))
}

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// tursoUrl := os.Getenv("TURSO_DATABASE_URL")
	// tursoKey := os.Getenv("TURSO_AUTH_TOKEN")

	// now do something with s3 or whatever
	url := "libsql://[DATABASE].turso.io?authToken=[TOKEN]"

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
	defer db.Close()
}

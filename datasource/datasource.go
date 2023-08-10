package datasource

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	GLOBAL_DB *sql.DB
)

func init() {
	dbConn := os.Getenv("DB_CONN")
	if len(dbConn) == 0 {
		dbConn = "root:12345678@tcp(127.0.0.1:3306)/test_db"
	}

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Fatal(fmt.Errorf("database open error: %w", err))
	}
	GLOBAL_DB = db
}

// TODO: DB Connection, close ... with repository

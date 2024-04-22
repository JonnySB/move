package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "move"
)

func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	// dynamic insert
	insertDynStmt := `insert into "exercises"("exercise", "reps", "sets", "weight") values($1, $2, $3, $4)`
	_, e := db.Exec(insertDynStmt, "Bench", 8, 8, 52.5)
	CheckError(e)

	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

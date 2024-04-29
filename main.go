package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "move"
)

func retrieveRecords(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rowRs, err := db.Query("SELECT * FROM exercises")

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer rowsRs.Close()
}

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

	// query
	rows, err := db.Query(`SELECT * FROM exercises`)
	CheckError(err)

	type exercise struct {
		id       int
		exercise string
		reps     int
		sets     int
		weight   float64
	}

	var exercises []exercise

	defer rows.Close()
	for rows.Next() {
		exercise := exercise{}

		err = rows.Scan(&exercise.id, &exercise.exercise, &exercise.reps, &exercise.sets, &exercise.weight)
		CheckError(err)

		fmt.Print(exercise)
		exercises = append(exercises, exercise)
	}
	CheckError(err)

	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

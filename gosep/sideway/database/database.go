package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := sql.Open("postgres", "postgres://gosctihb:CqOz6dVYlooEBPY4quY9KHvySa2OmADZ@arjuna.db.elephantsql.com:5432/gosctihb")

	if err != nil {
		log.Fatal("connect to database error", err)
	}

	defer db.Close()

	createTb := `CREATE TABLE IF NOT EXISTS pairs (
      DEVICE_ID INTEGER NOT NULL,
       USER_ID INTEGER NOT NULL
   );`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table todos", err)
	}
	fmt.Println("create table success.")
}

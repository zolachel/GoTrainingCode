package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var err error
var db *sql.DB

func init() {
	db, err = sql.Open("postgres", "postgres://gosctihb:CqOz6dVYlooEBPY4quY9KHvySa2OmADZ@arjuna.db.elephantsql.com:5432/gosctihb")
	if err != nil {
		log.Fatal(err)
	}
}

func createTable() {
	createTb := `
	CREATE TABLE IF NOT EXISTS todos (
	id SERIAL PRIMARY KEY,
	title TEXT,
	status TEXT
	);
	`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table", err)
	}
}

func insertTodo(title string, status string) {
	row := db.QueryRow("INSERT INTO todos (title, status) values ($1, $2) RETURNING id", title, status)
	var id int
	err = row.Scan(&id)
	if err != nil {
		fmt.Println("can't scan id", err)
		return
	}
	fmt.Println("insert todo success id : ", id)
}

func queryRow(id int) {
	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		log.Fatal("can'tprepare query one row statment", err)
	}
	row := stmt.QueryRow(id)
	var title, status string
	err = row.Scan(&id, &title, &status)
	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}
	fmt.Println("query result: ", id, title, status)
}

func updateStatus(id int, status string) {
	stmt, err := db.Prepare("UPDATE todos SET status=$2 WHERE id=$1;")
	if err != nil {
		log.Fatal("can't prepare statment update", err)
	}
	if _, err := stmt.Exec(id, status); err != nil {
		log.Fatal("error execute update ", err)
	}
	fmt.Println("update success")
}

func queryAll() {
	stmt, err := db.Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		log.Fatal("can't prepare query all todos statment", err)
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("can't query all todos", err)
	}
	for rows.Next() {
		var id int
		var title, status string
		err := rows.Scan(&id, &title, &status)
		if err != nil {
			log.Fatal("can't Scan row into variable", err)
		}
		fmt.Println(id, title, status)
	}
	fmt.Println("query all todos success")
}

func main() {
	createTable()
	insertTodo("buy benz", "active")
	queryRow(1)
	updateStatus(1, "inactive")
	queryAll()

	defer db.Close()
}

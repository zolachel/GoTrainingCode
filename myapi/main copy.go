package maincopy

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://gosctihb:CqOz6dVYlooEBPY4quY9KHvySa2OmADZ@arjuna.db.elephantsql.com:5432/gosctihb")
	if err != nil {
		log.Fatal(err)
	}
}

//Todo ...
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = map[int]*Todo{
	1: &Todo{ID: 1, Title: "pay phone bills", Status: "active"},
}

func helloHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

func getTodoByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) //convert id from string to int - strconv is string convert
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	/*
		t, ok := todos[id]
		if !ok {
			c.JSON(http.StatusOK, gin.H{}) //don't return null on REST API - just return empty
			return
		}
	*/

	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		log.Fatal("can't prepare query one row statment", err)
	}
	row := stmt.QueryRow(id)
	var title, status string
	err = row.Scan(&id, &title, &status)
	if err != nil {
		c.JSON(http.StatusOK, "No data found!")
	} else {
		c.JSON(http.StatusOK, Todo{ID: id, Title: title, Status: status})
	}
}

func getTodosHandler(context *gin.Context) {
	/*
		items := []*Todo{}
		for _, item := range todos {
			items = append(items, item)
		}
		context.JSON(http.StatusOK, items)
	*/
	stmt, err := db.Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		log.Fatal("can't prepare query all todos statment", err)
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("can't query all todos", err)
	}

	items := []*Todo{}

	for rows.Next() {
		var id int
		var title, status string
		err := rows.Scan(&id, &title, &status)
		if err != nil {
			log.Fatal("can't Scan row into variable", err)
		} else {
			items = append(items, &Todo{ID: id, Title: title, Status: status})
		}
	}

	context.JSON(http.StatusOK, items)
}

func createTodosHandler(c *gin.Context) {
	t := Todo{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	/*
		id := len(todos)
		id++
		t.ID = id
		todos[t.ID] = &t
	*/

	row := db.QueryRow("INSERT INTO todos (title, status) values ($1, $2) RETURNING id", t.Title, t.Status)
	var id int
	err = row.Scan(&id)
	if err != nil {
		fmt.Println("can't scan id", err)
		return
	} else {
		c.JSON(http.StatusCreated, id)
	}
}

func updateTodosHandler(c *gin.Context) {
	t := Todo{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err := db.Prepare("UPDATE todos SET title=$2, status=$3 WHERE id=$1;")
	if err != nil {
		log.Fatal("can't prepare statment update", err)
	}
	if _, err := stmt.Exec(t.ID, t.Title, t.Status); err != nil {
		log.Fatal("error execute update ", err)
	} else {
		c.JSON(http.StatusCreated, t)
	}
}

func authMiddleware(c *gin.Context) {
	log.Println("start middleware")
	authKey := c.GetHeader("Authorization")
	if authKey != "Bearer token123" {
		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		c.Abort()
		return
	}
	c.Next()
	log.Println("end middleware")
}

func setupRouter() *gin.Engine {
	route := gin.Default()
	route.Use(authMiddleware)

	route.GET("/hello", func(c *gin.Context) {
		c.String(200, "hi.")
	})

	//route.GET("/hello", helloHandler)

	route.GET("/todos", getTodosHandler)
	route.GET("/todos/:id", getTodoByIDHandler)
	route.POST("/todos", createTodosHandler)
	route.PUT("/todos", updateTodosHandler)

	return route
}

func main() {
	route := setupRouter()

	route.Run(":1234") // listen and serve on 127.0.0.0:8080
}

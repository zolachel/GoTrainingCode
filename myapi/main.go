package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zolachel/myapi/middleware"
	"github.com/zolachel/myapi/task"

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

func setupRouter() *gin.Engine {
	route := gin.Default()

	apiV1 := route.Group("api/v1")

	apiV1.Use(middleware.Auth)

	handler := task.Handler{DB: db}

	apiV1.GET("/todos", handler.GetTodosHandler)
	apiV1.GET("/todos/:id", handler.GetTodoByIDHandler)
	apiV1.POST("/todos", handler.CreateTodosHandler)
	apiV1.PUT("/todos", handler.UpdateTodosHandler)
	apiV1.DELETE("/todos/:id", handler.DeleteTodosHandler)

	return route
}

func main() {
	route := setupRouter()

	route.Run(":1234") // listen and serve on 127.0.0.0:8080

	defer db.Close()
}

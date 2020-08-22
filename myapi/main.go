package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
	t, ok := todos[id]
	if !ok {
		c.JSON(http.StatusOK, gin.H{}) //don't return null on REST API - just return empty
		return
	}
	c.JSON(http.StatusOK, t)
}

func getTodosHandler(context *gin.Context) {
	items := []*Todo{}
	for _, item := range todos {
		items = append(items, item)
	}
	context.JSON(http.StatusOK, items)
}

func createTodosHandler(c *gin.Context) {
	t := Todo{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := len(todos)
	id++
	t.ID = id
	todos[t.ID] = &t
	c.JSON(http.StatusCreated, "created todo.")
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
	//route.GET("/hello", helloHandler)
	//route.GET("/todos", getTodosHandler)
	//route.GET("/todos/:id", getTodoByIDHandler)
	//route.POST("/todos", createTodosHandler)

	route.Use(authMiddleware)

	route.GET("/hello", func(c *gin.Context) {
		c.String(200, "hi.")
	})
	return route
}

func main() {
	route := setupRouter()

	route.Run(":1234") // listen and serve on 127.0.0.0:8080
}

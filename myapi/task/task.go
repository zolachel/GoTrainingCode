package task

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var err error

//Todo ...
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

//Handler ...
type Handler struct {
	DB *sql.DB
}

//GetTodoByIDHandler ...
func (h *Handler) GetTodoByIDHandler(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id")) //convert id from string to int - strconv is string convert
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	stmt, err := h.DB.Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	row := stmt.QueryRow(id)
	var title, status string
	err = row.Scan(&id, &title, &status)
	if err != nil {
		context.JSON(http.StatusOK, "No data found!")
	} else {
		context.JSON(http.StatusOK, Todo{ID: id, Title: title, Status: status})
	}
}

//GetTodosHandler ...
func (h *Handler) GetTodosHandler(context *gin.Context) {
	stmt, err := h.DB.Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	rows, err := stmt.Query()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	} else {
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
}

//CreateTodosHandler ...
func (h *Handler) CreateTodosHandler(context *gin.Context) {
	t := Todo{}
	if err := context.ShouldBindJSON(&t); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	row := h.DB.QueryRow("INSERT INTO todos (title, status) values ($1, $2) RETURNING id", t.Title, t.Status)
	var id int
	err = row.Scan(&id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	} else {
		context.JSON(http.StatusCreated, id)
	}
}

//UpdateTodosHandler ...
func (h *Handler) UpdateTodosHandler(context *gin.Context) {
	t := Todo{}
	if err := context.ShouldBindJSON(&t); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err := h.DB.Prepare("UPDATE todos SET title=$2, status=$3 WHERE id=$1;")
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	if _, err := stmt.Exec(t.ID, t.Title, t.Status); err != nil {
		context.JSON(http.StatusInternalServerError, err)
	} else {
		context.JSON(http.StatusOK, t)
	}
}

//DeleteTodosHandler ...
func (h *Handler) DeleteTodosHandler(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id")) //convert id from string to int - strconv is string convert
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	stmt, err := h.DB.Prepare("DELETE FROM todos WHERE id=$1;")
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	if _, err := stmt.Exec(id); err != nil {
		context.JSON(http.StatusInternalServerError, err)
	} else {
		context.JSON(http.StatusOK, fmt.Sprintf("id = %d was deleted", id))
	}
}

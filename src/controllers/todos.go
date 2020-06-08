package controllers

import (
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	orm "github.com/go-pg/pg/orm"
	"github.com/go-pg/pg"
	guuid "github.com/google/uuid"
)

type Todo struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Completed string `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
// Create Todo Table

func CreateTodoTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&Todo{}, opts)
	if createError != nil {
		log.Printf("Error while creating todo table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Todo table created")
	return nil
}

// Initialise database
var dbConnect *pg.DB
func InitiateDB(db *pg.DB) {
	dbConnect = db
}

func GetAllTodos(c *gin.Context) {
	var todos []Todo
	err := dbConnect.Model(&todos).Select()
	if err != nil {
		log.Printf("Error while getting all todos. Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Something on Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
			"message": "All todo list",
			"data": todos,
	})
	return
}
func CreateTodo(c *gin.Context){
	var todo Todo
	c.BindJSON(&todo)

	title := todo.Title
	body := todo.Body
	completed := todo.Completed
	id := guuid.New().String()
	insertError := dbConnect.Insert(&Todo{
		ID: id,
		Title: title,
		Completed: completed,
		Body: body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	log.Printf("error: ",insertError!=nil)
	if insertError!=nil {
		log.Printf("Error while inserting new todo, reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Todo created Successfully",
	})
}
func GetSingleTodo(c *gin.Context){
	todoId := c.Param("todoId")
	todo := &Todo{ID: todoId}
	err := dbConnect.Select(todo)
	if err != nil {
		log.Printf("Error while getting all todos. Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"message": "Not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "A Single Todo",
		"data": todo,
	})
	return
}
func EditTodo(c *gin.Context){
	todoId := c.Param("todoId")
	var todo Todo
	c.BindJSON(&todo)
	// newValue := Todo{
	// 	Title: todo
 //    Body: todo.Body,  // this column will not be updated
 //    Completed: todo.Completed,
	// }
// _, err := db.Model(&book).Set("title = ?title").WherePK().Returning("*").Update()
// 	body := todo.Body
// 	completed := todo.Completed
	_, err := dbConnect.Model(&todo).Where("id = ?", todoId).Update()
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Update failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Update Successfully",
		"data": todo,
	})
	return
}
func DeleteSingleTodo(c *gin.Context) {
	todoId := c.Param("todoId")
	todo := &Todo{ID: todoId}
	err :=dbConnect.Delete(todo)
	if err != nil {
		log.Printf("Error while deleting a single todo. Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Todo deleted successfully",
	})
	return
}

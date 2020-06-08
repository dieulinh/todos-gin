package routes
import "net/http"
import "github.com/gin-gonic/gin"
import "controllers"


func Routes(router *gin.Engine) {
	router.GET("/", Welcome)
	router.NoRoute(notFound)
	router.GET("/todos", controllers.GetAllTodos)
	router.POST("/todos", controllers.CreateTodo)
	router.GET("/todos/:todoId", controllers.GetSingleTodo)
	router.PUT("/todos/:todoId", controllers.EditTodo)
	router.DELETE("/todos/:todoId", controllers.DeleteSingleTodo)
}


func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Welcome to do API",
	})
}
func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"message": "Route not found",
	})
}

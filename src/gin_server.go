package main
import "github.com/gin-gonic/gin"
import "log"
import (
  "config"
  "controllers"
  "routes"
)


func main() {
  db := config.Connect()
  controllers.CreateTodoTable(db)
  controllers.InitiateDB(db)
  router := gin.Default()
  routes.Routes(router)

  log.Println("Main log...")
  log.Fatal(router.Run(":8000"))
}

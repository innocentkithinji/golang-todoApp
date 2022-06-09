package main

import (
	"github.com/spf13/viper"
	"todoAPI/config"
	"todoAPI/controller"
	"todoAPI/repository"
	"todoAPI/server"
	"todoAPI/service"
)

var (
	firestoreRepo  = repository.NewFirestoreRepo()
	todoService    = service.NewTodoService(firestoreRepo)
	todoController = controller.NewTodoController(todoService)
	todoServer     = server.NewServer()
)

func main() {
	config.InitializeConfig()

	tGroup := todoServer.AddGroup("/todo")

	tGroup.POST("/", todoController.CreateTodo)
	tGroup.GET("/:id", todoController.RetrieveTodo)
	tGroup.GET("/", todoController.RetrieveAll)
	tGroup.PUT("/:id", todoController.UpdateTodo)
	tGroup.DELETE("/:id", todoController.DeleteTodo)

	port := viper.Get("port").(string)
	todoServer.Serve(port)

}

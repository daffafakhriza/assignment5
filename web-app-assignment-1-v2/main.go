package main

import (
	"a21hc3NpZ25tZW50/api"
	"a21hc3NpZ25tZW50/db/filebased"
	repo "a21hc3NpZ25tZW50/repository"
	"a21hc3NpZ25tZW50/service"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// APIHandler holds the API handlers for categories and tasks
type APIHandler struct {
	CategoryAPIHandler api.CategoryAPI
	TaskAPIHandler     api.TaskAPI
}

func main() {
	gin.SetMode(gin.ReleaseMode) // Set Gin to release mode

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] \"%s %s %s\"\n",
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery()) // Recover from panics

	filebasedDb, err := filebased.InitDB() // Initialize the database
	if err != nil {
		panic(err)
	}

	router = RunServer(router, filebasedDb) // Set up the server and routes

	fmt.Println("Server is running on port 8080")
	err = router.Run(":8080") // Start the server
	if err != nil {
		panic(err)
	}
}

// RunServer sets up the routes and initializes API handlers
func RunServer(gin *gin.Engine, filebasedDb *filebased.Data) *gin.Engine {
	categoryRepo := repo.NewCategoryRepo(filebasedDb)
	taskRepo := repo.NewTaskRepo(filebasedDb)

	categoryService := service.NewCategoryService(categoryRepo)
	taskService := service.NewTaskService(taskRepo)

	categoryAPIHandler := api.NewCategoryAPI(categoryService)
	taskAPIHandler := api.NewTaskAPI(taskService)

	apiHandler := APIHandler{
		CategoryAPIHandler: categoryAPIHandler,
		TaskAPIHandler:     taskAPIHandler,
	}

	// Task routes
	task := gin.Group("/task")
	{
		task.POST("/add", apiHandler.TaskAPIHandler.AddTask)
		task.GET("/get/:id", apiHandler.TaskAPIHandler.GetTaskByID)
		task.PUT("/update/:id", apiHandler.TaskAPIHandler.UpdateTask)
		task.DELETE("/delete/:id", apiHandler.TaskAPIHandler.DeleteTask)
		task.GET("/list", apiHandler.TaskAPIHandler.GetTaskList)
		task.GET("/category/:id", apiHandler.TaskAPIHandler.GetTaskListByCategory)
	}

	// Category routes
	category := gin.Group("/category")
	{
		category.POST("/add", apiHandler.CategoryAPIHandler.AddCategory)
		category.GET("/get/:id", apiHandler.CategoryAPIHandler.GetCategoryByID)
		category.PUT("/update/:id", apiHandler.CategoryAPIHandler.UpdateCategory)
		category.DELETE("/delete/:id", apiHandler.CategoryAPIHandler.DeleteCategory)
		category.GET("/list", apiHandler.CategoryAPIHandler.GetCategoryList)
	}

	return gin
}
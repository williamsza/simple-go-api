package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	//"github.com/pborman/uuid"
	//routers "simple-go-api/routers/v1"
)

func main() {
	r := gin.Default()

	appRoutes(r)

	r.Run("localhost:3001")
}

func appRoutes(router *gin.Engine) *gin.RouterGroup {
	taskController := newTaskController()
	v1 := router.Group("/v1")
	{
		v1.GET("/tasks", taskController.getAllTasks)
		v1.POST("/tasks", taskController.createTask)
		v1.DELETE("/tasks/:id", taskController.deleteTask)

	}
	return v1

}

type taskController struct {
	tasks []Task
}

func newTaskController() *taskController {
	return &taskController{
		tasks: []Task{},
	}
}
func (tc *taskController) getAllTasks(cxt *gin.Context) {
	cxt.JSON(http.StatusOK, tc.tasks)

}

func (tc *taskController) createTask(cxt *gin.Context) {
	task := newTask()
	if err := cxt.BindJSON(&task); err != nil {
		return
	}
	tc.tasks = append(tc.tasks, *task)
	cxt.JSON(http.StatusOK, tc.tasks)
}
func (tc *taskController) deleteTask(cxt *gin.Context) {
	id := cxt.Param("id")

	for idx, task := range tc.tasks {
		if task.ID == id {
			tc.tasks = append(tc.tasks[0:idx], tc.tasks[idx+1:]...)
			cxt.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
			return
		}
	}
	cxt.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

type Task struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func newTask() *Task {
	task := Task{
		ID: uuid.New().String(),
	}
	return &task
}

package routes


import (
	"task-manager/handlers"
	"github.com/gin-gonic/gin"
)


func RegisterRoutes(r *gin.Engine){
	r.POST("/tasks",handlers.CreateTask)
	r.GET("/tasks",handlers.GetAllTask)
	r.GET("/tasks/:id",handlers.GetTaskByID)
	r.PUT("/tasks/:id",handlers.UpdateTask)
	r.DELETE("/tasks/:id",handlers.DeleteTask)
}
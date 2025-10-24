package handlers

import (
	"net/http"
	"strconv"
	"task-manager/models"
	"task-manager/repository"
	"github.com/gin-gonic/gin"
)


func CreateTask(c *gin.Context){
	var t models.Task
	if err:=c.BindJSON(&t); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}
	created := repository.CreateTask(t)
	c.JSON(http.StatusCreated,created)
}

func GetAllTask(c *gin.Context){
	c.JSON(http.StatusOK,repository.GetAllTasks())
}

func GetTaskByID(c *gin.Context){
	id,err:= strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": "Invalid ID"})
		return
	}
	task := repository.GetTaskByID(id)
	if task==nil{
		c.JSON(http.StatusNotFound,gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK,task)
}

func UpdateTask(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": "invalid id"})
		return
	}
	var updated models.Task
	if err:=c.BindJSON(&updated); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
	}
	task :=repository.UpdateTask(id,updated)
	if task == nil{
		c.JSON(http.StatusNotFound,gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK,task)
}

func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	ok := repository.DeleteTask(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}

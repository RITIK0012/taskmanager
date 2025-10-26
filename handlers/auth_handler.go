package handlers

import (
	"net/http"
	"task-manager/models"
	"task-manager/repository"
	"task-manager/utils"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context){
	var user models.User
	if err:=c.BindJSON(&user); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":  err.Error()})
		return 
	}
	if !repository.CreateUser(user.Username, user.Password) {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Signup successful"})
}


func Login(c *gin.Context){
	var user models.User
	if err:=c.BindJSON(&user); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}

	if!repository.ValidateUser(user.Username,user.Password){
		c.JSON(http.StatusUnauthorized,gin.H{"error": "Invalid credentials"})
		return
	}

	token,err:=utils.GenerateToken(user.Username)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"token":token})
}
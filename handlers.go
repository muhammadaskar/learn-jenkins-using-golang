package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}

func GetUser(c *gin.Context) {
	var user User
	userID := c.Param("id")

	if err := DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}

func UpdateUser(c *gin.Context) {
	var user User
	userID := c.Param("id")

	if err := DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
		return
	}

	c.BindJSON(&user)
	DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}

func DeleteUser(c *gin.Context) {
	var user User
	userID := c.Param("id")

	if err := DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
		return
	}

	DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User deleted successfully"})
}

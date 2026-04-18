package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

// path parameter
func path_parameter(c *gin.Context){
	path := c.Param("name")
	c.JSON(http.StatusOK, gin.H{"message": "Path Value: " +  path})
}

func query_parameter(c *gin.Context){
	query := c.Query("id")
	c.JSON(http.StatusOK, gin.H{"message":"Query value: " + query})
}

type Login struct {
  User     string `json:"user" binding:"required"`
  Password string `json:"password" binding:"required"`
}

func payload_parameter(c *gin.Context){
	var payload Login
	if err := c.ShouldBindJSON(&payload); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
	c.JSON(http.StatusCreated, gin.H{"data": payload.User + " " + payload.Password})
	return
}


func main() {
  router := gin.Default()
  
  router.GET("/api/:name", path_parameter)
  router.GET("/api", query_parameter)
  router.POST("/login", payload_parameter)
  router.Run() // listens on 0.0.0.0:8080 by default
}
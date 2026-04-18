package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)



func getting(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"message": "You also started GO series"})
}

func main() {
  router := gin.Default()
  
  router.GET("/ping", getting)
  router.Run() // listens on 0.0.0.0:8080 by default
}
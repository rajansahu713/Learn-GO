package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreatesItem struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name" binding:"required"`
}

type UpdateItem struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

var items = []Item{
  {ID: "1", Name: "Item 1"},
  {ID: "2", Name: "Item 2"},
}

func get_items(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"data": items})
		return
	}
	
	for _, item := range items {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
	return
}

func create_item(c *gin.Context) {
	var new_item CreatesItem
	if err := c.ShouldBindJSON(&new_item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Additional validation using validator
	validate := validator.New()
	if err := validate.Struct(new_item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name field is required"})
		return
	}

	// Generate ID if not provided
	if new_item.ID == "" {
		new_item.ID = fmt.Sprintf("%d", len(items)+1)
	}

	// Convert to Item struct for storage
	item := Item{
		ID:   new_item.ID,
		Name: new_item.Name,
	}

	items = append(items, item)
	c.JSON(http.StatusCreated, gin.H{"message": "Item created successfully", "item": item})
}

func update_item(c *gin.Context) {
	id := c.Param("id")
	for i, item := range items {
		if item.ID == id {
			var updated_item Item
			if err := c.BindJSON(&updated_item); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			items[i] = updated_item
			c.JSON(http.StatusOK, gin.H{"method": "PUT"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}


func delete_item(c *gin.Context) {
	id := c.Param("id")
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"method": "DELETE"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})		
}

func main() {
  // Creates a gin router with default middleware:
  // logger and recovery (crash-free) middleware
  router := gin.Default()

  router.GET("/items", get_items)
  router.POST("/items", create_item)
  router.PUT("/items/:id", update_item)
  router.DELETE("/items/:id", delete_item)

  // By default it serves on :8080 unless a
  // PORT environment variable was defined.
  router.Run()
  // router.Run(":3000") for a hard coded port
}
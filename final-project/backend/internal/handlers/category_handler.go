package handlers

import (
	"example/backend/internal/models"
	"example/backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	categories := services.FetchCategories()
	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func PostCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	result := services.PostCategory(category)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	category, err := services.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var updatedCategory models.Category

	if err := c.ShouldBindJSON(&updatedCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	err := services.UpdateCategory(id, updatedCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

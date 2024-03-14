package controllers

import (
	"assignment2-golang/database"
	"assignment2-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.GetDB().Create(&item)
	c.JSON(http.StatusCreated, gin.H{
		"data": item,
	})
}

func GetItemById(c *gin.Context) {
	var item models.Item
	id := c.Param("item_id")
	if err := database.GetDB().Where("item_id = ?", id).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Item not found!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": item,
	})
}
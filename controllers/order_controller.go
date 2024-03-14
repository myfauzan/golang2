package controllers

import (
	"assignment2-golang/database"
	"assignment2-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder (c *gin.Context) {
	var order = models.Order{
		OrderID: c.GetUint("OrderID"),
		CustomerName: c.PostForm("CustomerName"),
		Items: []models.Item{
			{
				ItemCode: c.PostForm("ItemCode"),
				Description: c.PostForm("Description"),
				Quantity: c.GetInt("Quantity"),
			},
		},
		OrderedAt: c.GetTime("OrderedAt"),
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := database.GetDB().Create(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Order not created!",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": order,
	})
}

func GetOrders(c *gin.Context) {
	var orders []models.Order
	err := database.GetDB().Preload("Items").Find(&orders).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Order not found!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": orders,
	})
}

func UpdateOrder(c *gin.Context) {
	var order models.Order
	id := c.Param("id")
	if err := database.GetDB().Where("order_id= ?", id).Preload("Items").First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Order not found!",
		})
		return
	}

	var updateOrder models.Order
	if err := c.ShouldBindJSON(&updateOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Ensure OrderID is correct
	updateOrder.OrderID = order.OrderID
	
	// Update order details
	if err := database.GetDB().Model(&order).Updates(&updateOrder).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Update failed!",
		})
		return
	}

	// Update items
	for i, item := range updateOrder.Items {
		if err := database.GetDB().Model(&order.Items[i]).Updates(&item).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to update item!",
			})
			return
		}
	}

	// Fetch updated order from the database
	if err := database.GetDB().Where("order_id = ?", id).Preload("Items").First(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch updated order!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}



func DeleteOrder(c *gin.Context) {
	var order models.Order
	id := c.Param("id")
	if err := database.GetDB().Where("order_id = ?", id).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Order not found!",
		})
		return
	}

	err :=  database.GetDB().Select("Items").Delete(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Order not deleted",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order has been deleted successfully",
	})
}
package handlers

import (
	"net/http"
	"strconv"

	"blood-donor-api/config"
	"blood-donor-api/models"

	"github.com/gin-gonic/gin"
)

// GET /api/donors?blood_group=A%2B&area=Dhaka
func GetDonors(c *gin.Context) {
	var donors []models.Donor
	q := config.DB.Order("created_at desc")

	if bg := c.Query("blood_group"); bg != "" {
		q = q.Where("blood_group = ?", bg)
	}
	if area := c.Query("area"); area != "" {
		q = q.Where("area LIKE ?", "%"+area+"%")
	}
	if available := c.Query("available"); available == "true" {
		q = q.Where("available = ?", true)
	}

	if err := q.Find(&donors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": len(donors), "donors": donors})
}

func GetDonor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var donor models.Donor
	if err := config.DB.First(&donor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Donor not found"})
		return
	}
	c.JSON(http.StatusOK, donor)
}

func CreateDonor(c *gin.Context) {
	var donor models.Donor
	if err := c.ShouldBindJSON(&donor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&donor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, donor)
}

func UpdateDonor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var donor models.Donor
	if err := config.DB.First(&donor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Donor not found"})
		return
	}
	if err := c.ShouldBindJSON(&donor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&donor)
	c.JSON(http.StatusOK, donor)
}

func DeleteDonor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Delete(&models.Donor{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Donor deleted"})
}

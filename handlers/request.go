package handlers

import (
	"net/http"
	"strconv"

	"blood-donor-api/config"
	"blood-donor-api/models"

	"github.com/gin-gonic/gin"
)

func GetRequests(c *gin.Context) {
	var reqs []models.BloodRequest
	q := config.DB.Order("created_at desc")

	if bg := c.Query("blood_group"); bg != "" {
		q = q.Where("blood_group = ?", bg)
	}
	if status := c.Query("status"); status != "" {
		q = q.Where("status = ?", status)
	} else {
		q = q.Where("status = ?", "open")
	}

	if err := q.Find(&reqs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": len(reqs), "requests": reqs})
}

func CreateRequest(c *gin.Context) {
	var req models.BloodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, req)
}

func UpdateRequestStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req models.BloodRequest
	if err := config.DB.First(&req, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
		return
	}
	var body struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Status = body.Status
	config.DB.Save(&req)
	c.JSON(http.StatusOK, req)
}

func DeleteRequest(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.Delete(&models.BloodRequest{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Request deleted"})
}

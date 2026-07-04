package handlers

import (
	"net/http"

	"blood-donor-api/config"
	"blood-donor-api/models"

	"github.com/gin-gonic/gin"
)

// GET /api/stats — dashboard data
func GetStats(c *gin.Context) {
	type GroupCount struct {
		BloodGroup string `json:"blood_group"`
		Count      int    `json:"count"`
	}

	var totalDonors int64
	var totalRequests int64
	var openRequests int64
	var fulfilledRequests int64

	config.DB.Model(&models.Donor{}).Count(&totalDonors)
	config.DB.Model(&models.BloodRequest{}).Count(&totalRequests)
	config.DB.Model(&models.BloodRequest{}).Where("status = ?", "open").Count(&openRequests)
	config.DB.Model(&models.BloodRequest{}).Where("status = ?", "fulfilled").Count(&fulfilledRequests)

	var groups []GroupCount
	config.DB.Model(&models.Donor{}).
		Select("blood_group, COUNT(*) as count").
		Group("blood_group").
		Order("count desc").
		Scan(&groups)

	c.JSON(http.StatusOK, gin.H{
		"total_donors":        totalDonors,
		"total_requests":      totalRequests,
		"open_requests":       openRequests,
		"fulfilled_requests":  fulfilledRequests,
		"donors_by_blood":     groups,
	})
}

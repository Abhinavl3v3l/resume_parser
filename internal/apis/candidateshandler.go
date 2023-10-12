package apis

import (
	"net/http"
	"strconv"

	"github.com/automagic-tools/go-coding-challenge/SeeCV/internal/db/ops"
	"github.com/automagic-tools/go-coding-challenge/SeeCV/utils/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetCandidateHandler fetches details of a specific candidate based on their ID
func GetCandidateHandler(c *gin.Context) {
	candidateIDStr := c.Param("id")
	candidateID, err := strconv.Atoi(candidateIDStr)
	if err != nil {
		logger.Error("Invalid candidate ID format", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid candidate ID format",
		})
		return
	}

	candidate, err := ops.GetCandidate(int32(candidateID))
	if err != nil {
		logger.Error("Error fetching candidate details", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching candidate details",
		})
		return
	}

	c.JSON(http.StatusOK, candidate)
}

// GetSkillsForCandidateHandler fetches all skills associated with a given candidate
func GetSkillsForCandidateHandler(c *gin.Context) {
	candidateIDStr := c.Param("id")
	candidateID, err := strconv.Atoi(candidateIDStr)
	if err != nil {
		logger.Error("Invalid candidate ID format", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid candidate ID format",
		})
		return
	}

	skills, err := ops.GetSkillsForCandidate(int32(candidateID))
	if err != nil {
		logger.Error("Error fetching skills for candidate", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching skills for candidate",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"skills": skills,
	})
}

func GetCandidateByEmailHandler(c *gin.Context) {
	email := c.Query("email")

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email query parameter is required"})
		return
	}

	candidate, err := ops.GetCandidateByEmail(email)
	if err != nil {
		logger.Error("Failed to fetch candidate by email", email, zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch candidate"})
		return
	}

	if candidate == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Candidate not found"})
		return
	}

	c.JSON(http.StatusOK, candidate)
}

// DeleteCandidateHandler deletes a specific candidate based on their ID
func DeleteCandidateHandler(c *gin.Context) {
	candidateIDStr := c.Param("id")
	candidateID, err := strconv.Atoi(candidateIDStr)
	if err != nil {
		logger.Error("Invalid candidate ID format", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid candidate ID format",
		})
		return
	}

	err = ops.DeleteCandidate(int32(candidateID))
	if err != nil {
		logger.Error("Error deleting candidate", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting candidate",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Candidate successfully deleted",
	})
}

package apis

import (
	"encoding/json"
	"fmt"
	"github.com/automagic-tools/go-coding-challenge/SeeCV/internal/db/ops"
	"github.com/automagic-tools/go-coding-challenge/SeeCV/internal/services"
	"github.com/automagic-tools/go-coding-challenge/SeeCV/internal/types"
	"github.com/automagic-tools/go-coding-challenge/SeeCV/utils/logger"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// extractCandidateInfo Extract Candidates information from resposnse object from ChatGPT.
func extractCandidateInfo(r io.Reader) (types.CandidateInfo, error) {
	var candidateInfo types.CandidateInfo

	candidateData, err := services.InsightStream(r)
	if err != nil {
		return candidateInfo, fmt.Errorf("insight into document failed: %w", err)
	}

	if err := json.Unmarshal([]byte(candidateData), &candidateInfo); err != nil {
		return candidateInfo, fmt.Errorf("failed to unmarshal candidate: %w", err)
	}

	return candidateInfo, nil
}

// UploadResumeHandler Handler for request upload-resume, Handles only PDF documents
func UploadResumeHandler(c *gin.Context) {
	fileOperation := services.NewFileOperations(c)
	reader, err := fileOperation.ConvertPDFToStream()
	if err != nil {
		logger.Error("Failed to convert PDF to stream", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to convert PDF to stream"})
		return
	}

	candidateInfo, err := extractCandidateInfo(reader)
	if err != nil {
		logger.Error("Failed to extract candidate info", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to extract candidate info"})
		return
	}

	logger.Info("Candidate Information", candidateInfo)

	if err := ops.AddCandidates(candidateInfo); err != nil {
		logger.Error("Failed to add candidates to DB", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add candidates to DB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded and content extracted"})
}

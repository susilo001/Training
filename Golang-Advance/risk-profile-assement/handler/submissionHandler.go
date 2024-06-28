package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/susilo001/golang-advance/risk-profile-assement/entity"
	"github.com/susilo001/golang-advance/risk-profile-assement/service"
	"github.com/susilo001/golang-advance/risk-profile-assement/utility"
)

// ISubmissionHandler defines the interface for submission handlers
type ISubmissionHandler interface {
	CreateSubmission(c *gin.Context)
	UpdateSubmission(c *gin.Context)
	DeleteSubmission(c *gin.Context)
	GetAllSubmissions(c *gin.Context)
	GetSubmissionByID(c *gin.Context)
}

// SubmissionHandler is the implementation of ISubmissionHandler
type SubmissionHandler struct {
	submissionService service.ISubmissionService
}

// NewSubmissionHandler creates a new instance of SubmissionHandler
func NewSubmissionHandler(submissionService service.ISubmissionService) ISubmissionHandler {
	return &SubmissionHandler{
		submissionService: submissionService,
	}
}

// CreateSubmission handles the request to create a new submission
func (h *SubmissionHandler) CreateSubmission(c *gin.Context) {
	var submission entity.Submission
	if err := c.ShouldBindJSON(&submission); err != nil {
		errMsg := err.Error()
		errMsg = utility.ConvertMandatoryFieldErrorMessage(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	createdSubmission, err := h.submissionService.CreateSubmission(c.Request.Context(), &submission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdSubmission)
}

// UpdateSubmission handles the request to update a submission
func (h *SubmissionHandler) UpdateSubmission(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var submission entity.Submission
	if err := c.ShouldBindJSON(&submission); err != nil {
		errMsg := err.Error()
		errMsg = utility.ConvertMandatoryFieldErrorMessage(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}
	submission.ID = id

	updatedSubmission, err := h.submissionService.UpdateSubmission(c.Request.Context(), id, &submission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedSubmission)
}

// DeleteSubmission handles the request to delete a submission
func (h *SubmissionHandler) DeleteSubmission(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.submissionService.DeleteSubmission(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Submission deleted"})
}

// GetAllSubmissions handles the request to get all submissions
func (h *SubmissionHandler) GetAllSubmissions(c *gin.Context) {
	submissions, err := h.submissionService.GetAllSubmissions(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, submissions)
}

// GetSubmissionByID handles the request to get a submission by ID
func (h *SubmissionHandler) GetSubmissionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	submission, err := h.submissionService.GetSubmissionByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Submission not found"})
		return
	}

	c.JSON(http.StatusOK, submission)
}

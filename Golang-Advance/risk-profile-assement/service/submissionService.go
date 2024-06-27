package service

import (
	"context"
	"fmt"

	"github.com/susilo001/golang-advance/risk-profile-assement/entity"
)

// ISubmissionService defines the interface for submission services
type ISubmissionService interface {
	CreateSubmission(ctx context.Context, submission *entity.Submission) (entity.Submission, error)
	UpdateSubmission(ctx context.Context, id int, submission *entity.Submission) (entity.Submission, error)
	DeleteSubmission(ctx context.Context, id int) error
	GetAllSubmissions(ctx context.Context) ([]entity.Submission, error)
	GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error)
}

// ISubmissionRepository defines the interface for submission repository
type ISubmissionRepository interface {
	CreateSubmission(ctx context.Context, submission *entity.Submission) (entity.Submission, error)
	UpdateSubmission(ctx context.Context, id int, submission *entity.Submission) (entity.Submission, error)
	DeleteSubmission(ctx context.Context, id int) error
	GetAllSubmissions(ctx context.Context) ([]entity.Submission, error)
	GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error)
}

// submissionService is the implementation of ISubmissionService using ISubmissionRepository
type submissionService struct {
	submissionRepo ISubmissionRepository
}

// NewSubmissionService creates a new instance of submissionService
func NewSubmissionService(submissionRepo ISubmissionRepository) ISubmissionService {
	return &submissionService{submissionRepo: submissionRepo}
}

// CreateSubmission creates a new submission
func (s *submissionService) CreateSubmission(ctx context.Context, submission *entity.Submission) (entity.Submission, error) {
	createdSubmission, err := s.submissionRepo.CreateSubmission(ctx, submission)
	if err != nil {
		return entity.Submission{}, fmt.Errorf("failed to create submission: %v", err)
	}
	return createdSubmission, nil
}

// UpdateSubmission updates an existing submission
func (s *submissionService) UpdateSubmission(ctx context.Context, id int, submission *entity.Submission) (entity.Submission, error) {
	updatedSubmission, err := s.submissionRepo.UpdateSubmission(ctx, id, submission)
	if err != nil {
		return entity.Submission{}, fmt.Errorf("failed to update submission: %v", err)
	}
	return updatedSubmission, nil
}

// DeleteSubmission deletes a submission by ID
func (s *submissionService) DeleteSubmission(ctx context.Context, id int) error {
	err := s.submissionRepo.DeleteSubmission(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete submission: %v", err)
	}
	return nil
}

// GetAllSubmissions retrieves all submissions
func (s *submissionService) GetAllSubmissions(ctx context.Context) ([]entity.Submission, error) {
	submissions, err := s.submissionRepo.GetAllSubmissions(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all submissions: %v", err)
	}
	return submissions, nil
}

// GetSubmissionByID retrieves a submission by ID
func (s *submissionService) GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error) {
	submission, err := s.submissionRepo.GetSubmissionByID(ctx, id)
	if err != nil {
		return entity.Submission{}, fmt.Errorf("failed to get submission by ID: %v", err)
	}
	return submission, nil
}

func CheckUserSubmittedAnswer() {

}

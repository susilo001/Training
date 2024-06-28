package service

import (
	"context"
	"fmt"

	"github.com/susilo001/golang-advance/risk-profile-assement/entity"
)

type ISubmissionService interface {
	CreateSubmission(ctx context.Context, submission *entity.Submission) (entity.Submission, error)
	UpdateSubmission(ctx context.Context, id int, submission *entity.Submission) (entity.Submission, error)
	DeleteSubmission(ctx context.Context, id int) error
	GetAllSubmissions(ctx context.Context) ([]entity.Submission, error)
	GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error)
}

type ISubmissionRepository interface {
	CreateSubmission(ctx context.Context, submission *entity.Submission) (entity.Submission, error)
	UpdateSubmission(ctx context.Context, id int, submission *entity.Submission) (entity.Submission, error)
	DeleteSubmission(ctx context.Context, id int) error
	GetAllSubmissions(ctx context.Context) ([]entity.Submission, error)
	GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error)
}

type submissionService struct {
	submissionRepo ISubmissionRepository
}

func NewSubmissionService(submissionRepo ISubmissionRepository) ISubmissionService {
	return &submissionService{submissionRepo: submissionRepo}
}

func (s *submissionService) CreateSubmission(ctx context.Context, submission *entity.Submission) (entity.Submission, error) {
	createdSubmission, err := s.submissionRepo.CreateSubmission(ctx, submission)
	if err != nil {
		return entity.Submission{}, fmt.Errorf("failed to create submission: %v", err)
	}
	return createdSubmission, nil
}

func (s *submissionService) UpdateSubmission(ctx context.Context, id int, submission *entity.Submission) (entity.Submission, error) {
	updatedSubmission, err := s.submissionRepo.UpdateSubmission(ctx, id, submission)
	if err != nil {
		return entity.Submission{}, fmt.Errorf("failed to update submission: %v", err)
	}
	return updatedSubmission, nil
}

func (s *submissionService) DeleteSubmission(ctx context.Context, id int) error {
	err := s.submissionRepo.DeleteSubmission(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete submission: %v", err)
	}
	return nil
}

func (s *submissionService) GetAllSubmissions(ctx context.Context) ([]entity.Submission, error) {
	submissions, err := s.submissionRepo.GetAllSubmissions(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all submissions: %v", err)
	}
	return submissions, nil
}

func (s *submissionService) GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error) {
	submission, err := s.submissionRepo.GetSubmissionByID(ctx, id)
	if err != nil {
		return entity.Submission{}, fmt.Errorf("failed to get submission by ID: %v", err)
	}
	return submission, nil
}

func CheckUserSubmittedAnswer() {

}

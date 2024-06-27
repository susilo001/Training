package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/susilo001/golang-advance/risk-profile-assement/entity"
	mock_service "github.com/susilo001/golang-advance/risk-profile-assement/mock/service"
	"github.com/susilo001/golang-advance/risk-profile-assement/service"
)

func TestSubmissionService_CreateSubmission(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_service.NewMockISubmissionRepository(ctrl)
	submissionService := service.NewSubmissionService(mockRepo)

	ctx := context.Background()
	submission := &entity.Submission{
		UserID:       1,
		RiskCategory: "john@example.com",
		RiskScore:    99,
	}

	t.Run("PositiveCase", func(t *testing.T) {
		mockRepo.EXPECT().CreateSubmission(ctx, submission).Return(*submission, nil)

		createdSubmission, err := submissionService.CreateSubmission(ctx, submission)
		assert.NoError(t, err)
		assert.Equal(t, *submission, createdSubmission)
	})

	t.Run("NegativeCase", func(t *testing.T) {
		mockRepo.EXPECT().CreateSubmission(ctx, submission).Return(entity.Submission{}, errors.New("failed to create submission"))

		createdSubmission, err := submissionService.CreateSubmission(ctx, submission)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create submission")
		assert.Equal(t, entity.Submission{}, createdSubmission)
	})
}

func TestSubmissionService_GetSubmissionByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_service.NewMockISubmissionRepository(ctrl)
	submissionService := service.NewSubmissionService(mockRepo)

	ctx := context.Background()
	submissionID := 1
	submission := entity.Submission{
		ID:           submissionID,
		UserID:       1,
		RiskCategory: "john@example.com",
		RiskScore:    99,
	}

	t.Run("PositiveCase", func(t *testing.T) {
		mockRepo.EXPECT().GetSubmissionByID(ctx, submissionID).Return(submission, nil)

		foundSubmission, err := submissionService.GetSubmissionByID(ctx, submissionID)
		assert.NoError(t, err)
		assert.Equal(t, submission, foundSubmission)
	})

	t.Run("NegativeCase", func(t *testing.T) {
		mockRepo.EXPECT().GetSubmissionByID(ctx, submissionID).Return(entity.Submission{}, errors.New("submission not found"))

		foundSubmission, err := submissionService.GetSubmissionByID(ctx, submissionID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "submission not found")
		assert.Equal(t, entity.Submission{}, foundSubmission)
	})
}

func TestSubmissionService_UpdateSubmission(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_service.NewMockISubmissionRepository(ctrl)
	submissionService := service.NewSubmissionService(mockRepo)

	ctx := context.Background()
	submissionID := 1
	submission := entity.Submission{
		ID:           submissionID,
		UserID:       1,
		RiskCategory: "john@example.com",
		RiskScore:    99,
	}

	t.Run("PositiveCase", func(t *testing.T) {
		mockRepo.EXPECT().UpdateSubmission(ctx, submissionID, submission).Return(submission, nil)

		updatedSubmission, err := submissionService.UpdateSubmission(ctx, submissionID, &submission)
		assert.NoError(t, err)
		assert.Equal(t, submission, updatedSubmission)
	})

	t.Run("NegativeCase", func(t *testing.T) {
		mockRepo.EXPECT().UpdateSubmission(ctx, submissionID, submission).Return(entity.Submission{}, errors.New("submission not found"))

		updatedSubmission, err := submissionService.UpdateSubmission(ctx, submissionID, &submission)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "submission not found")
		assert.Equal(t, entity.Submission{}, updatedSubmission)
	})
}

func TestSubmissionService_DeleteSubmission(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_service.NewMockISubmissionRepository(ctrl)
	submissionService := service.NewSubmissionService(mockRepo)

	ctx := context.Background()
	submissionID := 1

	t.Run("PositiveCase", func(t *testing.T) {
		mockRepo.EXPECT().DeleteSubmission(ctx, submissionID).Return(nil)

		err := submissionService.DeleteSubmission(ctx, submissionID)
		assert.NoError(t, err)
	})

	t.Run("NegativeCase", func(t *testing.T) {
		mockRepo.EXPECT().DeleteSubmission(ctx, submissionID).Return(errors.New("submission not found"))

		err := submissionService.DeleteSubmission(ctx, submissionID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "submission not found")
	})
}

func TestSubmissionService_GetAllSubmissions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_service.NewMockISubmissionRepository(ctrl)
	submissionService := service.NewSubmissionService(mockRepo)

	ctx := context.Background()
	submissions := []entity.Submission{
		{
			ID:           1,
			UserID:       1,
			RiskCategory: "john@example.com",
			RiskScore:    99,
			
		},
		{
			ID:           2,
			UserID:       2,
			RiskCategory: "jane@example.com",
			RiskScore:    99,
			
		},
	}

	t.Run("PositiveCase", func(t *testing.T) {
		mockRepo.EXPECT().GetAllSubmissions(ctx).Return(submissions, nil)

		retrievedSubmissions, err := submissionService.GetAllSubmissions(ctx)
		assert.NoError(t, err)
		assert.Equal(t, submissions, retrievedSubmissions)
	})

	t.Run("NegativeCase", func(t *testing.T) {
		mockRepo.EXPECT().GetAllSubmissions(ctx).Return([]entity.Submission{}, errors.New("no submissions found"))

		retrievedSubmissions, err := submissionService.GetAllSubmissions(ctx)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "no submissions found")
		assert.Empty(t, retrievedSubmissions)
	})
}

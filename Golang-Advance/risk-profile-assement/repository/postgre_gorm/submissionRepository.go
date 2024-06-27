package postgres_gorm

import (
	"context"

	"github.com/susilo001/golang-advance/risk-profile-assement/entity"
)

type submissionRepository struct {
	db GormDBIface
}

func NewSubmissionRepository(db GormDBIface) *submissionRepository {
	return &submissionRepository{db}
}

func (rep *submissionRepository) CreateSubmission(ctx context.Context, submission *entity.Submission) (entity.Submission, error) {
	if err := rep.db.WithContext(ctx).Create(submission).Error; err != nil {
		return entity.Submission{}, err
	}
	return *submission, nil
}

func (rep *submissionRepository) UpdateSubmission(ctx context.Context, id int, submission *entity.Submission) (entity.Submission, error) {
	var existingSubmission entity.Submission
	if err := rep.db.WithContext(ctx).Select("id", "user_id", "answers", "risk_score", "risk_category", "created_at", "updated_at").First(&existingSubmission, id).Error; err != nil {
		return entity.Submission{}, err
	}

	existingSubmission.Answers = submission.Answers
	existingSubmission.RiskCategory = submission.RiskCategory
	existingSubmission.RiskScore = submission.RiskScore
	if err := rep.db.WithContext(ctx).Save(&existingSubmission).Error; err != nil {
		return entity.Submission{}, err
	}
	return *submission, nil
}

func (rep *submissionRepository) DeleteSubmission(ctx context.Context, id int) error {
	if err := rep.db.WithContext(ctx).Delete(&entity.Submission{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (rep *submissionRepository) GetAllSubmissions(ctx context.Context) ([]entity.Submission, error) {
	var submissions []entity.Submission
	if err := rep.db.WithContext(ctx).Find(&submissions).Error; err != nil {
		return nil, err
	}
	return submissions, nil
}

func (rep *submissionRepository) GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error) {
	var submission entity.Submission
	if err := rep.db.WithContext(ctx).First(&submission, id).Error; err != nil {
		return entity.Submission{}, err
	}
	return submission, nil
}

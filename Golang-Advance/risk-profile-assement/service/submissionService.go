package service

import (
	"context"
	"encoding/json"
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

	var answers []entity.Answer
	if err := json.Unmarshal(submission.Answers, &answers); err != nil {
		return entity.Submission{}, fmt.Errorf("failed to unmarshal json: %v", err)
	}

	riskScore := CalculateRiskScore(answers)

	submission.RiskScore = riskScore

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

type ProfileRiskCategory string

const (
	ProfileRiskCategoryConservative ProfileRiskCategory = "Conservative"
	ProfileRiskCategoryModerate     ProfileRiskCategory = "Moderate"
	ProfileRiskCategoryBalanced     ProfileRiskCategory = "Balanced"
	ProfileRiskCategoryGrowth       ProfileRiskCategory = "Growth"
	ProfileRiskCategoryAggresive    ProfileRiskCategory = "Aggresive"
)

type ProfileRisk struct {
	MinScore   int
	MaxScore   int
	Category   ProfileRiskCategory
	Definition string
}

var RiskMapping = []ProfileRisk{
	{
		MinScore: 0,
		MaxScore: 11,
		Category: ProfileRiskCategoryConservative,
		Definition: "Tujuan utama Anda adalah untuk melindungi modal/dana yang ditempatkan dan Anda tidak memiliki toleransi " +
			"sama sekali terhadap perubahan harga/nilai dari dana investasinya tersebut. " +
			"Anda memiliki pengalaman yang sangat terbatas atau tidak memiliki pengalaman sama sekali mengenai produk investasi.",
	},
	{
		MinScore:   12,
		MaxScore:   19,
		Category:   ProfileRiskCategoryModerate,
		Definition: "Anda memiliki toleransi yang rendah dengan perubahan harga/nilai dari dana investasi dan risiko investasi.",
	},
	{
		MinScore: 20,
		MaxScore: 28,
		Category: ProfileRiskCategoryBalanced,
		Definition: "Anda memiliki toleransi yang cukup terhadap produk investasi dan dapat menerima perubahan yang besar dari " +
			"harga/nilai dari harga yang diinvestasikan.",
	},
	{
		MinScore: 29,
		MaxScore: 35,
		Category: ProfileRiskCategoryGrowth,
		Definition: "Anda memiliki toleransi yang cukup tinggi dan dapat menerima perubahan yang besar dari harga/nilai portfolio" +
			"pada produk investasi yang diinvestasikan." +
			"Pada umumnya Anda sudah pernah atau berpengalaman dalam berinvestasi di produk investasi.",
	},
	{
		MinScore: 36,
		MaxScore: 40,
		Category: ProfileRiskCategoryAggresive,
		Definition: "Anda sangat berpengalaman terhadap produk investasi dan memiliki toleransi yang sangat tinggi atas" +
			"produk-produk investasi. Anda bahkan dapat menerima perubahan signifikan pada modal/nilai investasi." +
			"Pada umumnya portfolio Anda sebagian besar dialokasikan pada produk investasi.",
	},
}

func CalculateRiskScore(answers []entity.Answer) (score int) {
	questions := make(map[int]entity.Question)
	for _, q := range entity.Questions {
		questions[int(q.ID)] = q
	}

	riskScore := 0
	for _, a := range answers {
		q, exists := questions[a.QuestionID]
		if !exists {
			return 0
		}

		optionFound := false
		for _, o := range q.Options {
			if a.Answer == o.Answer {
				riskScore += o.Weight
				optionFound = true
				break
			}
		}

		if !optionFound {
			return 0
		}
	}
	return score
}

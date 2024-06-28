package entity

import (
	"encoding/json"
	"time"
)

type Submission struct {
	ID           int             `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       int             `gorm:"index;not null" json:"user_id"`
	Answers      json.RawMessage `gorm:"type:jsonb" json:"answers"`
	RiskScore    int           `gorm:"not null" json:"risk_score"`
	RiskCategory string          `gorm:"type:varchar(255);not null" json:"risk_category"`
	CreatedAt    *time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    *time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}


type Answer struct {
	QuestionID int
	Answer string
}

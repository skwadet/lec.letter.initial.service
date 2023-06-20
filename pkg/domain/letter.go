package domain

import (
	"github.com/google/uuid"
	"time"
)

type Letter struct {
	Id             uuid.UUID
	SubmissionId   uuid.UUID
	CreatedAt      time.Time
	AdditionalInfo *string

	// Not db mapped!
	LetterResearchInfo *LetterResearchInfo
	Documents          []*Document
	Objectives         []*Objective
	Purposes           []*Purpose
}

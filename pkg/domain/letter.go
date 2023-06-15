package domain

import (
	"fmt"
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
}

func (l *Letter) GetAdditionalInfo() *string {
	if l.AdditionalInfo == nil {
		return nil
	}

	return l.AdditionalInfo
}

func (l *Letter) GetDirection() *string {
	if l.LetterResearchInfo == nil {
		return nil
	}

	if l.LetterResearchInfo.CommitteeChairmanTitleTo == nil {
		return nil
	}

	if l.LetterResearchInfo.MainResearcher == nil {
		return nil
	}

	welcomePart := fmt.Sprintf("%s\n Главный исследователь\n %s\n %s",
		*l.LetterResearchInfo.CommitteeChairmanTitleTo,
		*l.LetterResearchInfo.MainResearcher,
		l.CreatedAt.String(),
	)

	return &welcomePart
}

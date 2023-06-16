package utils

import (
	"github.com/ecosafety/lec.letter.initial.service/pkg/domain"
	"github.com/google/uuid"
)

func SortPurposes(purposes []*domain.Purpose) []*domain.Purpose {
	return quickSortLetterAdditions(purposes)
}

func ChangePurposeOrderNumber(purposes []*domain.Purpose, newOrderNumber int, id uuid.UUID) []*domain.Purpose {
	index := ResolveLetterAdditionIndex(purposes, newOrderNumber)

	purposes = SortPurposes(purposes)

	purposes, poppedPurpose := PopLetterAddition(purposes, id)

	purposes = InsertLetterAddition(purposes, index, poppedPurpose)

	return purposes
}

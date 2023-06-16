package utils

import (
	"github.com/ecosafety/lec.letter.initial.service/pkg/domain"
	"github.com/google/uuid"
)

func SortObjectives(objectives []*domain.Objective) []*domain.Objective {
	return quickSortLetterAdditions(objectives)
}

func ChangeObjectiveOrderNumber(objectives []*domain.Objective, newOrderNumber int, id uuid.UUID) []*domain.Objective {
	index := ResolveLetterAdditionIndex(objectives, newOrderNumber)

	objectives = SortObjectives(objectives)

	objectives, poppedObjective := PopLetterAddition(objectives, id)

	objectives = InsertLetterAddition(objectives, index, poppedObjective)

	return objectives
}

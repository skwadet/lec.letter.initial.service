package utils

import (
	"github.com/ecosafety/lec.letter.initial.service/pkg/domain"
	"github.com/google/uuid"
)

// SortObjectives Ожидается малое кол-во элементов, смысла писать быструю сортировку нет, пользуемся пузырьковой
func SortObjectives(objectives []*domain.Objective) []*domain.Objective {
	return bubbleSortObjectives(objectives)
}

func bubbleSortObjectives(objectives []*domain.Objective) []*domain.Objective {
	sortedObjectives := make([]*domain.Objective, len(objectives))
	copy(sortedObjectives, objectives)

	// Всё просто, проходим по массиву, меняя местами соседние элементы, если L < R, до тех пор, пока всё не отсортируется
	for i := 0; i < len(sortedObjectives)-1; i++ {
		for j := 0; j < len(sortedObjectives)-i-1; j++ {
			if sortedObjectives[j].OrderNumber > sortedObjectives[j+1].OrderNumber {
				sortedObjectives[j], sortedObjectives[j+1] = sortedObjectives[j+1], sortedObjectives[j]
			}
		}
	}

	return sortedObjectives
}

func ChangeObjectiveOrderNumber(objectives []*domain.Objective, newOrderNumber int, id uuid.UUID) []*domain.Objective {
	var objectiveToChange *domain.Objective
	index := newOrderNumber - 1

	objectives = SortObjectives(objectives)

	for _, elem := range objectives {
		if elem.Id == id {
			objectiveToChange = elem
			objectives = append(objectives[:objectiveToChange.OrderNumber-1], objectives[objectiveToChange.OrderNumber:]...)
		}
	}

	objectives = append(objectives[:index+1], objectives[index:]...)
	objectives[index] = objectiveToChange

	return objectives
}

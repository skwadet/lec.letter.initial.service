package utils

import (
	"github.com/ecosafety/lec.letter.initial.service/pkg/domain"
)

// SortPurposes Ожидается малое кол-во элементов, смысла писать быструю сортировку нет, пользуемся пузырьковой
func SortPurposes(purposes []*domain.Purpose) []*domain.Purpose {
	return bubbleSortPurposes(purposes)
}

func bubbleSortPurposes(purposes []*domain.Purpose) []*domain.Purpose {
	sortedPurposes := make([]*domain.Purpose, len(purposes))
	copy(sortedPurposes, purposes)

	// Всё просто, проходим по массиву, меняя местами соседние элементы, если L < R, до тех пор, пока всё не отсортируется
	for i := 0; i < len(sortedPurposes)-1; i++ {
		for j := 0; j < len(sortedPurposes)-i-1; j++ {
			if sortedPurposes[j].OrderNumber > sortedPurposes[j+1].OrderNumber {
				sortedPurposes[j], sortedPurposes[j+1] = sortedPurposes[j+1], sortedPurposes[j]
			}
		}
	}

	return sortedPurposes
}

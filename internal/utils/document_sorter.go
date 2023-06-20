package utils

import (
	"github.com/ecosafety/lec.letter.initial.service/pkg/domain"
)

// SortDocuments ожидается большое кол-во элементов, задействуем быструю сортировку
func SortDocuments(docs []*domain.Document) []*domain.Document {
	return quickSortLetterAdditions(docs)
}

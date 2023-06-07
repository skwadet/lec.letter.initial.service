package utils

import (
	"github.com/ecosafety/lec.letter.initial.service/pkg/domain"

	"math/rand"
)

// SortDocuments ожидается большое кол-во элементов, задействуем быструю сортировку
func SortDocuments(docs []*domain.Document) []*domain.Document {
	return quickSortDocuments(docs)
}

func quickSortDocuments(docs []*domain.Document) []*domain.Document {
	if len(docs) < 2 {
		return docs
	}

	// Определяем левый и правый элемент
	left, endOfSlice := 0, len(docs)-1

	// Определяем рандомно заданный элемент, который будет разделителем (пивот)
	borderElem := rand.Int() % len(docs)

	// Меняем местами правый элемент и пивот
	docs[borderElem], docs[endOfSlice] = docs[endOfSlice], docs[borderElem]

	// Итерируемся по всему массиву документов
	for i := range docs {
		// Если элемент, на котором указатель в данный момент меньше, чем бывший пивот и ныне правые элемент, то
		if docs[i].OrderNumber < docs[endOfSlice].OrderNumber {

			// 1. Меняем местами элемент, на котором остановилась итерация и left
			docs[left], docs[i] = docs[i], docs[left]

			// 2. Сдвигаем указатель left на одну единицу вправо
			left++

			// Итерируемся дальше...
		}
	}

	// По итогу у нас остаётся единственный элемент, который больше правого, меняем местами,
	// пивот встаёт на левую позицию и получается слева элементы меньше пивота, справа - больше
	docs[left], docs[endOfSlice] = docs[endOfSlice], docs[left]

	// Рекурсией запускаем такие же сортировки для массива слева, включая пивот и для массива справа пивота
	quickSortDocuments(docs[:left])
	quickSortDocuments(docs[left+1:])

	return docs
}

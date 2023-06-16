package utils

import (
	"github.com/ecosafety/lec.letter.initial.service/pkg/domain"
	"github.com/google/uuid"
	"math/rand"
)

func ResolveLetterAdditionIndex[T domain.LetterAddition](list []T, newOrderNumber int) int {
	index := newOrderNumber - 1
	if newOrderNumber >= len(list) {
		index = len(list) - 1
	}

	return index
}

func PopLetterAddition[T domain.LetterAddition](list []T, idToPop uuid.UUID) ([]T, T) {
	var res T
	for _, elem := range list {
		if elem.GetId() == idToPop {
			res = elem
			list = append(list[:res.GetOrderNumber()-1], list[res.GetOrderNumber():]...)
		}
	}

	return list, res
}

func InsertLetterAddition[T domain.LetterAddition](list []T, index int, additionToInsert T) []T {
	list = append(list[:index+1], list[index:]...)
	list[index] = additionToInsert

	return list
}

func quickSortLetterAdditions[T domain.LetterAddition](list []T) []T {
	if len(list) < 2 {
		return list
	}

	// Определяем левый и правый элемент
	left, endOfSlice := 0, len(list)-1

	// Определяем рандомно заданный элемент, который будет разделителем (пивот)
	borderElem := rand.Int() % len(list)

	// Меняем местами правый элемент и пивот
	list[borderElem], list[endOfSlice] = list[endOfSlice], list[borderElem]

	// Итерируемся по всему массиву
	for i := range list {
		// Если элемент, на котором указатель в данный момент меньше, чем бывший пивот и ныне правые элемент, то
		if list[i].GetOrderNumber() < list[endOfSlice].GetOrderNumber() {

			// 1. Меняем местами элемент, на котором остановилась итерация и left
			list[left], list[i] = list[i], list[left]

			// 2. Сдвигаем указатель left на одну единицу вправо
			left++

			// Итерируемся дальше...
		}
	}

	// По итогу у нас остаётся единственный элемент, который больше правого, меняем местами,
	// пивот встаёт на левую позицию и получается слева элементы меньше пивота, справа - больше
	list[left], list[endOfSlice] = list[endOfSlice], list[left]

	// Рекурсией запускаем такие же сортировки для массива слева, включая пивот и для массива справа пивота
	quickSortLetterAdditions(list[:left])
	quickSortLetterAdditions(list[left+1:])

	return list
}

func bubbleSortLetterAdditions[T domain.LetterAddition](list []T) []T {
	sortedList := make([]T, len(list))
	copy(sortedList, list)

	// Всё просто, проходим по массиву, меняя местами соседние элементы, если L < R, до тех пор, пока всё не отсортируется
	for i := 0; i < len(sortedList)-1; i++ {
		for j := 0; j < len(sortedList)-i-1; j++ {
			if sortedList[j].GetOrderNumber() > sortedList[j+1].GetOrderNumber() {
				sortedList[j], sortedList[j+1] = sortedList[j+1], sortedList[j]
			}
		}
	}

	return sortedList
}

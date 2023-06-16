package domain

import "github.com/google/uuid"

type LetterAddition interface {
	GetId() uuid.UUID
	GetOrderNumber() int
}

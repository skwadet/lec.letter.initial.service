package domain

import (
	"github.com/google/uuid"
	"time"
)

type Purpose struct {
	Id          uuid.UUID
	OrderNumber int
	LetterId    uuid.UUID
	Title       string
	CreatedAt   time.Time
}

func (p *Purpose) GetId() uuid.UUID {
	return p.Id
}

func (p *Purpose) GetOrderNumber() int {
	return p.OrderNumber
}

package domain

import "github.com/google/uuid"

type Document struct {
	Id          uuid.UUID
	OrderNumber int
	Title       string
}

func (d *Document) GetId() uuid.UUID {
	return d.Id
}

func (d *Document) GetOrderNumber() int {
	return d.OrderNumber
}

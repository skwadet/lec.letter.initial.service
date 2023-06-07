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

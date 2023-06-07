package domain

import (
	"github.com/google/uuid"
	"time"
)

type Objective struct {
	Id          uuid.UUID
	OrderNumber int
	LetterId    uuid.UUID
	Title       string
	CreatedAt   time.Time
}

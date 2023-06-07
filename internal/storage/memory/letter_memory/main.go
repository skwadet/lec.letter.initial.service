package letter_memory

import (
	"context"
	"errors"

	"github.com/ecosafety/lec.letter.initial.service/pkg/domain"

	"github.com/google/uuid"

	"sync"
)

type Cache struct {
	sync.RWMutex
	letterMap map[uuid.UUID]*domain.Letter
}

func NewCache() *Cache {
	letterMap := make(map[uuid.UUID]*domain.Letter)

	cache := Cache{
		letterMap: letterMap,
	}

	return &cache
}

func (c *Cache) SetData(letters []*domain.Letter) {
	c.Lock()

	defer c.Unlock()

	for _, elem := range letters {
		c.letterMap[elem.SubmissionId] = elem
	}
}

func (c *Cache) Get(_ context.Context, submissionId uuid.UUID) (*domain.Letter, error) {
	letter := c.letterMap[submissionId]

	if letter != nil {
		return letter, nil
	}

	return nil, errors.New("letter not found")
}

func (c *Cache) Add(_ context.Context, letter *domain.Letter) error {
	c.Lock()

	defer c.Unlock()

	c.letterMap[letter.Id] = letter

	return nil
}

func (c *Cache) Rename(ctx context.Context, submissionId uuid.UUID, newInfo string) error {
	c.Lock()

	defer c.Unlock()

	var letterToUpdate, err = c.Get(ctx, submissionId)
	if err != nil {
		return err
	}

	c.letterMap[submissionId] = &domain.Letter{
		Id:             letterToUpdate.Id,
		SubmissionId:   letterToUpdate.SubmissionId,
		CreatedAt:      letterToUpdate.CreatedAt,
		AdditionalInfo: &newInfo,
	}

	return nil
}

func (c *Cache) Delete(_ context.Context, submissionId uuid.UUID) error {
	c.Lock()

	defer c.Unlock()

	delete(c.letterMap, submissionId)

	return nil
}

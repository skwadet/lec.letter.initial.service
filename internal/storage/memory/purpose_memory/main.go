package purpose_memory

import (
	"context"
	"errors"

	"github.com/ecosafety/lec.letter.initial.service/pkg/domain"

	"github.com/google/uuid"

	"sync"
)

type Cache struct {
	sync.RWMutex
	purposeMap  map[uuid.UUID]*domain.Purpose
	relationMap map[uuid.UUID][]uuid.UUID
}

func NewCache() *Cache {
	purposeMap := make(map[uuid.UUID]*domain.Purpose)

	cache := Cache{
		purposeMap: purposeMap,
	}

	return &cache
}

func (c *Cache) SetData(purposes []*domain.Purpose) {
	relations := make([]uuid.UUID, 0)
	c.Lock()

	defer c.Unlock()

	for _, elem := range purposes {
		relations = c.relationMap[elem.LetterId]

		c.relationMap[elem.LetterId] = append(relations, elem.Id)
		c.purposeMap[elem.Id] = elem
	}
}

func (c *Cache) Get(_ context.Context, id uuid.UUID) (*domain.Purpose, error) {
	purpose := c.purposeMap[id]

	if purpose != nil {
		return purpose, nil
	}

	return nil, errors.New("purpose not found")
}

func (c *Cache) GetListByLetterId(ctx context.Context, letterId uuid.UUID) ([]*domain.Purpose, error) {
	var purposes []*domain.Purpose
	ids := c.relationMap[letterId]

	if ids != nil {
		return nil, errors.New("purpose ids not found")
	}

	for _, id := range ids {
		purpose, err := c.Get(ctx, id)
		if err != nil {
			return nil, errors.New("purpose not found")
		}

		purposes = append(purposes, purpose)
	}

	return purposes, nil
}

func (c *Cache) Add(_ context.Context, purpose *domain.Purpose) error {
	c.Lock()

	defer c.Unlock()

	c.purposeMap[purpose.Id] = purpose
	c.relationMap[purpose.LetterId] = append(c.relationMap[purpose.LetterId], purpose.Id)

	return nil
}

func (c *Cache) Rename(ctx context.Context, id uuid.UUID, newTitle string) error {
	c.Lock()

	defer c.Unlock()

	var purposeToUpdate, err = c.Get(ctx, id)
	if err != nil {
		return err
	}

	c.purposeMap[id] = &domain.Purpose{
		Id:          purposeToUpdate.Id,
		OrderNumber: purposeToUpdate.OrderNumber,
		LetterId:    purposeToUpdate.LetterId,
		Title:       newTitle,
		CreatedAt:   purposeToUpdate.CreatedAt,
	}

	return nil
}

func (c *Cache) ChangeOrderNumber(ctx context.Context, id uuid.UUID, newOrderNumber int) error {
	c.Lock()

	defer c.Unlock()

	var purposeToUpdate, err = c.Get(ctx, id)
	if err != nil {
		return err
	}

	c.purposeMap[id] = &domain.Purpose{
		Id:          purposeToUpdate.Id,
		OrderNumber: newOrderNumber,
		LetterId:    purposeToUpdate.LetterId,
		Title:       purposeToUpdate.Title,
		CreatedAt:   purposeToUpdate.CreatedAt,
	}

	return nil
}

func (c *Cache) Delete(ctx context.Context, id uuid.UUID) error {
	c.Lock()

	defer c.Unlock()

	var purposeToDelete, err = c.Get(ctx, id)
	if err != nil {
		return err
	}

	idToDelete := purposeToDelete.LetterId

	purposeIds := c.relationMap[idToDelete]

	delete(c.purposeMap, idToDelete)

	for index, objectiveId := range purposeIds {
		if objectiveId == idToDelete {
			c.relationMap[idToDelete] = append(purposeIds[:index], purposeIds[index+1:]...)
			return nil
		}
	}

	return nil
}

package objective_memory

import (
	"context"
	"errors"

	"github.com/ecosafety/lec.letter.initial.service/pkg/domain"

	"github.com/google/uuid"

	"sync"
)

type Cache struct {
	sync.RWMutex
	objectiveMap map[uuid.UUID]*domain.Objective
	relationMap  map[uuid.UUID][]uuid.UUID
}

func NewCache() *Cache {
	objectiveMap := make(map[uuid.UUID]*domain.Objective)
	relationMap := make(map[uuid.UUID][]uuid.UUID)

	cache := Cache{
		objectiveMap: objectiveMap,
		relationMap:  relationMap,
	}

	return &cache
}

func (c *Cache) SetData(objectives []*domain.Objective) {
	relations := make([]uuid.UUID, 0)
	c.Lock()

	defer c.Unlock()

	for _, elem := range objectives {
		relations = c.relationMap[elem.LetterId]

		c.relationMap[elem.LetterId] = append(relations, elem.Id)
		c.objectiveMap[elem.Id] = elem
	}
}

func (c *Cache) Get(_ context.Context, id uuid.UUID) (*domain.Objective, error) {
	objective := c.objectiveMap[id]

	if objective != nil {
		return objective, nil
	}

	return nil, errors.New("objective not found")
}

func (c *Cache) GetListByLetterId(ctx context.Context, letterId uuid.UUID) ([]*domain.Objective, error) {
	var objectives []*domain.Objective
	ids := c.relationMap[letterId]

	if ids != nil {
		return nil, errors.New("objective ids not found")
	}

	for _, id := range ids {
		objective, err := c.Get(ctx, id)
		if err != nil {
			return nil, errors.New("objective not found")
		}

		objectives = append(objectives, objective)
	}

	return objectives, nil
}

func (c *Cache) Add(_ context.Context, objective *domain.Objective) error {
	c.Lock()

	defer c.Unlock()

	c.objectiveMap[objective.Id] = objective
	c.relationMap[objective.LetterId] = append(c.relationMap[objective.LetterId], objective.Id)

	return nil
}

func (c *Cache) Rename(ctx context.Context, id uuid.UUID, newTitle string) error {
	c.Lock()

	defer c.Unlock()

	var objectiveToUpdate, err = c.Get(ctx, id)
	if err != nil {
		return err
	}

	c.objectiveMap[id] = &domain.Objective{
		Id:          objectiveToUpdate.Id,
		OrderNumber: objectiveToUpdate.OrderNumber,
		LetterId:    objectiveToUpdate.LetterId,
		Title:       newTitle,
		CreatedAt:   objectiveToUpdate.CreatedAt,
	}

	return nil
}

func (c *Cache) ChangeOrder(ctx context.Context, id uuid.UUID, newOrderNumber int) error {
	c.Lock()

	defer c.Unlock()

	var objectiveToUpdate, err = c.Get(ctx, id)
	if err != nil {
		return err
	}

	c.objectiveMap[id] = &domain.Objective{
		Id:          objectiveToUpdate.Id,
		OrderNumber: newOrderNumber,
		LetterId:    objectiveToUpdate.LetterId,
		Title:       objectiveToUpdate.Title,
		CreatedAt:   objectiveToUpdate.CreatedAt,
	}

	return nil
}

func (c *Cache) Delete(ctx context.Context, id uuid.UUID) error {
	c.Lock()

	defer c.Unlock()

	var objectiveToDelete, err = c.Get(ctx, id)
	if err != nil {
		return err
	}

	idToDelete := objectiveToDelete.LetterId

	objectiveIds := c.relationMap[idToDelete]

	delete(c.objectiveMap, idToDelete)

	for index, objectiveId := range objectiveIds {
		if objectiveId == idToDelete {
			c.relationMap[idToDelete] = append(objectiveIds[:index], objectiveIds[index+1:]...)
		}
	}

	return nil
}

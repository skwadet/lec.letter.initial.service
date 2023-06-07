package pg_objective

import (
	"context"

	"github.com/ecosafety/lec.letter.initial.service/pkg/domain"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	getListQ           = "SELECT id, order_number, letter_id, title, created_at FROM objectives"
	addQ               = "INSERT INTO objectives (id, order_number, letter_id, title, created_at) VALUES($1, $2, $3, $4)"
	renameQ            = "UPDATE objectives SET title = $1 WHERE id = $2"
	changeOrderNumberQ = "UPDATE objectives SET order_number = $1 WHERE id = $2"
	deleteQ            = "DELETE FROM objectives WHERE id = $1"
)

type Storage struct {
	dbPool *pgxpool.Pool
}

func NewStorage(dbPool *pgxpool.Pool) *Storage {
	return &Storage{dbPool: dbPool}
}

func (s *Storage) GetList(ctx context.Context) ([]*domain.Objective, error) {
	var outputList []*domain.Objective

	if err := pgxscan.Select(ctx, s.dbPool, &outputList, getListQ); err != nil {
		return nil, err
	}

	return outputList, nil
}

func (s *Storage) Add(ctx context.Context, req *domain.Objective) error {
	if _, err := s.dbPool.Exec(ctx, addQ, req.Id, req.OrderNumber, req.LetterId, req.Title, req.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (s *Storage) ChangeOrderNumber(ctx context.Context, id uuid.UUID, newOrderNumber int) error {
	if _, err := s.dbPool.Exec(ctx, changeOrderNumberQ, newOrderNumber, id); err != nil {
		return err
	}

	return nil
}

func (s *Storage) Rename(ctx context.Context, id uuid.UUID, newTitle string) error {
	if _, err := s.dbPool.Exec(ctx, renameQ, newTitle, id); err != nil {
		return err
	}

	return nil
}

func (s *Storage) Delete(ctx context.Context, id uuid.UUID) error {
	if _, err := s.dbPool.Exec(ctx, deleteQ, id); err != nil {
		return err
	}

	return nil
}

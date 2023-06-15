package pg_letter

import (
	"context"

	"github.com/ecosafety/lec.letter.initial.service/pkg/domain"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	getListQ = "SELECT id, submission_id, created_at, additional_info FROM letters"
	addQ     = "INSERT INTO letters (id, submission_id, created_at, additional_info) VALUES($1, $2, $3, $4)"
	renameQ  = "UPDATE letters SET additional_info = $1 WHERE id = $2"
	deleteQ  = "DELETE FROM letters WHERE id = $1"
)

type Storage struct {
	dbPool *pgxpool.Pool
}

func NewStorage(dbPool *pgxpool.Pool) *Storage {
	return &Storage{dbPool: dbPool}
}

func (s *Storage) GetList(ctx context.Context) ([]*domain.Letter, error) {
	var outputList []*domain.Letter

	if err := pgxscan.Select(ctx, s.dbPool, &outputList, getListQ); err != nil {
		return nil, err
	}

	return outputList, nil
}

func (s *Storage) Add(ctx context.Context, req *domain.Letter) error {
	if _, err := s.dbPool.Exec(ctx, addQ, req.Id, req.SubmissionId, req.CreatedAt, req.AdditionalInfo); err != nil {
		return err
	}

	return nil
}

func (s *Storage) Rename(ctx context.Context, id uuid.UUID, newAdditionalInfo string) error {
	if _, err := s.dbPool.Exec(ctx, renameQ, newAdditionalInfo, id); err != nil {
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

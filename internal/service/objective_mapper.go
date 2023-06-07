package service

import (
	"context"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"

	"github.com/ecosafety/lec.letter.initial.service/pkg/domain"
	"github.com/google/uuid"
)

func mapAddObjectiveRequestFromProto(_ context.Context, req *pb.AddObjectiveRequest) (*domain.Objective, error) {
	lId, err := uuid.Parse(req.GetLetterId())
	if err != nil {
		return nil, err
	}

	return &domain.Objective{
		LetterId: lId,
		Title:    req.GetTitle(),
	}, nil
}

func mapObjectiveListToProto(ctx context.Context, objectives []*domain.Objective) ([]*pb.ObjectiveData, error) {
	var mappedObjectives []*pb.ObjectiveData

	for _, objective := range objectives {
		mappedObjective, mErr := mapObjectiveToProto(ctx, objective)
		if mErr != nil {
			return nil, mErr
		}

		mappedObjectives = append(mappedObjectives, mappedObjective)
	}

	return mappedObjectives, nil
}

func mapObjectiveToProto(_ context.Context, letter *domain.Objective) (*pb.ObjectiveData, error) {
	return &pb.ObjectiveData{
		Id:          letter.Id.String(),
		OrderNumber: int32(letter.OrderNumber),
		LetterId:    letter.LetterId.String(),
		Title:       letter.Title,
	}, nil
}

package service

import (
	"context"
	"github.com/ecosafety/lec.letter.initial.service/internal/utils"

	gateway "github.com/ecosafety/lec.letter.initial.service/internal/pb/document_generator_api"
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

func mapObjectiveToProto(_ context.Context, objective *domain.Objective) (*pb.ObjectiveData, error) {
	return &pb.ObjectiveData{
		Id:          objective.Id.String(),
		OrderNumber: int32(objective.OrderNumber),
		LetterId:    objective.LetterId.String(),
		Title:       objective.Title,
		CreatedAt:   utils.InnerFormatTime(objective.CreatedAt),
	}, nil
}

func mapObjectiveListToReq(ctx context.Context, objectives []*domain.Objective) ([]*gateway.ObjectiveData, error) {
	var mappedObjectives []*gateway.ObjectiveData

	for _, objective := range objectives {
		mappedObjective, mErr := mapObjectiveToReq(ctx, objective)
		if mErr != nil {
			return nil, mErr
		}

		mappedObjectives = append(mappedObjectives, mappedObjective)
	}

	return mappedObjectives, nil
}

func mapObjectiveToReq(_ context.Context, objective *domain.Objective) (*gateway.ObjectiveData, error) {
	return &gateway.ObjectiveData{
		Id:          objective.Id.String(),
		OrderNumber: int32(objective.OrderNumber),
		LetterId:    objective.LetterId.String(),
		Title:       objective.Title,
		CreatedAt:   utils.InnerFormatTime(objective.CreatedAt),
	}, nil
}

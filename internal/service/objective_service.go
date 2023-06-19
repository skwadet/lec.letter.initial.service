package service

import (
	"context"
	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
	"github.com/ecosafety/lec.letter.initial.service/internal/utils"

	"github.com/google/uuid"
	"time"
)

func (s *Service) AddObjective(ctx context.Context, req *pb.AddObjectiveRequest) (*pb.AddObjectiveResponse, error) {
	objective, mErr := mapAddObjectiveRequestFromProto(ctx, req)
	if mErr != nil {
		return nil, mErr
	}

	objectiveId := uuid.New()

	objective.Id = objectiveId
	objective.CreatedAt = time.Now()

	objectives, gMemErr := s.objectiveMemCache.GetListByLetterId(ctx, objective.LetterId)
	if gMemErr != nil {
		return nil, gMemErr
	}

	objective.OrderNumber = len(objectives) + 1

	pgErr := s.objectiveStorage.Add(ctx, objective)
	if pgErr != nil {
		return nil, pgErr
	}

	memErr := s.objectiveMemCache.Add(ctx, objective)
	if memErr != nil {
		return nil, memErr
	}

	return &pb.AddObjectiveResponse{
		Message: "Задача успешно добавлена!",
		Data:    objectiveId.String(),
	}, nil

}

func (s *Service) GetObjective(ctx context.Context, req *pb.GetObjectiveRequest) (*pb.GetObjectiveResponse, error) {
	id, pErr := uuid.Parse(req.GetId())
	if pErr != nil {
		return nil, pErr
	}

	objective, gErr := s.objectiveMemCache.Get(ctx, id)
	if gErr != nil {
		return nil, gErr
	}

	mappedObjective, mErr := mapObjectiveToProto(ctx, objective)
	if mErr != nil {
		return nil, mErr
	}

	return &pb.GetObjectiveResponse{
		Message: "",
		Data:    mappedObjective,
	}, nil
}

func (s *Service) GetObjectiveList(ctx context.Context,
	req *pb.GetObjectiveListRequest) (*pb.GetObjectiveListResponse, error) {
	letterId, pErr := uuid.Parse(req.GetLetterId())
	if pErr != nil {
		return nil, pErr
	}

	objectives, gErr := s.objectiveMemCache.GetListByLetterId(ctx, letterId)
	if gErr != nil {
		return nil, gErr
	}
	objectives = utils.SortObjectives(objectives)

	mappedObjective, mErr := mapObjectiveListToProto(ctx, objectives)
	if mErr != nil {
		return nil, mErr
	}

	return &pb.GetObjectiveListResponse{
		Message: "",
		Data:    mappedObjective,
	}, nil

}

func (s *Service) RenameObjective(ctx context.Context,
	req *pb.RenameObjectiveRequest) (*pb.RenameObjectiveResponse, error) {
	id, pErr := uuid.Parse(req.GetId())

	if pErr != nil {
		return nil, pErr
	}

	newTitle := req.GetNewTitle()

	pgErr := s.objectiveStorage.Rename(ctx, id, newTitle)
	if pgErr != nil {
		return nil, pgErr
	}

	memErr := s.objectiveMemCache.Rename(ctx, id, newTitle)
	if memErr != nil {
		return nil, memErr
	}

	objective, gErr := s.objectiveMemCache.Get(ctx, id)
	if gErr != nil {
		return nil, gErr
	}

	return &pb.RenameObjectiveResponse{
		Message: "Задача успешно обновлена!",
		Data:    objective.Id.String(),
	}, nil
}

func (s *Service) ChangeObjectiveOrderNumber(ctx context.Context,
	req *pb.ChangeObjectiveOrderNumberRequest) (*pb.ChangeObjectiveOrderNumberResponse, error) {
	id, pErr := uuid.Parse(req.GetId())
	if pErr != nil {
		return nil, pErr
	}

	objective, gErr := s.objectiveMemCache.Get(ctx, id)
	if gErr != nil {
		return nil, gErr
	}

	objectives, glErr := s.objectiveMemCache.GetListByLetterId(ctx, objective.LetterId)
	if glErr != nil {
		return nil, glErr
	}

	newOrderNumber := int(req.GetNewOrderNumber())

	objectives = utils.ChangeObjectiveOrderNumber(objectives, newOrderNumber, id)

	for index, elem := range objectives {
		eId := elem.Id

		pgErr := s.objectiveStorage.ChangeOrderNumber(ctx, eId, index+1)
		if pgErr != nil {
			return nil, pgErr
		}

		memErr := s.objectiveMemCache.ChangeOrderNumber(ctx, eId, index+1)
		if memErr != nil {
			return nil, memErr
		}
	}

	return &pb.ChangeObjectiveOrderNumberResponse{
		Message: "Задача успешно обновлена!",
		Data:    objective.Id.String(),
	}, nil
}

func (s *Service) DeleteObjective(ctx context.Context,
	req *pb.DeleteObjectiveRequest) (*pb.DeleteObjectiveResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}

	pgErr := s.objectiveStorage.Delete(ctx, id)
	if pgErr != nil {
		return nil, pgErr
	}

	memErr := s.objectiveMemCache.Delete(ctx, id)
	if memErr != nil {
		return nil, memErr
	}

	return &pb.DeleteObjectiveResponse{
		Message: "Задача успешно удалена!",
	}, nil

}

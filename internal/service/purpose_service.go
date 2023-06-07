package service

import (
	"context"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
	"github.com/ecosafety/lec.letter.initial.service/internal/utils"

	"github.com/google/uuid"
	"time"
)

func (s *Service) AddPurpose(ctx context.Context, req *pb.AddPurposeRequest) (*pb.AddPurposeResponse, error) {
	var maxOrderNumber int

	purpose, mErr := mapAddPurposeRequestFromProto(ctx, req)
	if mErr != nil {
		return nil, mErr
	}

	purposeId := uuid.New()

	purpose.Id = purposeId
	purpose.CreatedAt = time.Now()

	purposes, gMemErr := s.purposeMemCache.GetListByLetterId(ctx, purpose.LetterId)
	if gMemErr != nil {
		return nil, gMemErr
	}

	for _, elem := range purposes {
		if elem.OrderNumber < maxOrderNumber {
			maxOrderNumber = elem.OrderNumber
		}
	}
	purpose.OrderNumber = maxOrderNumber

	pgErr := s.purposeStorage.Add(ctx, purpose)
	if pgErr != nil {
		return nil, pgErr
	}

	memErr := s.purposeMemCache.Add(ctx, purpose)
	if memErr != nil {
		return nil, memErr
	}

	return &pb.AddPurposeResponse{
		Message: "Цель успешно добавлена!",
		Data:    purposeId.String(),
	}, nil

}

func (s *Service) GetPurpose(ctx context.Context, req *pb.GetPurposeRequest) (*pb.GetPurposeResponse, error) {
	id, pErr := uuid.Parse(req.GetId())
	if pErr != nil {
		return nil, pErr
	}

	purpose, gErr := s.purposeMemCache.Get(ctx, id)
	if gErr != nil {
		return nil, gErr
	}

	mappedPurpose, mErr := mapPurposeToProto(ctx, purpose)
	if mErr != nil {
		return nil, mErr
	}

	return &pb.GetPurposeResponse{
		Message: "",
		Data:    mappedPurpose,
	}, nil
}

func (s *Service) GetPurposeList(ctx context.Context,
	req *pb.GetPurposeListRequest) (*pb.GetPurposeListResponse, error) {
	letterId, pErr := uuid.Parse(req.GetLetterId())
	if pErr != nil {
		return nil, pErr
	}

	purposes, gErr := s.purposeMemCache.GetListByLetterId(ctx, letterId)
	if gErr != nil {
		return nil, gErr
	}
	purposes = utils.SortPurposes(purposes)

	mappedPurpose, mErr := mapPurposeListToProto(ctx, purposes)
	if mErr != nil {
		return nil, mErr
	}

	return &pb.GetPurposeListResponse{
		Message: "",
		Data:    mappedPurpose,
	}, nil

}

func (s *Service) RenamePurpose(ctx context.Context,
	req *pb.RenamePurposeRequest) (*pb.RenamePurposeResponse, error) {
	id, pErr := uuid.Parse(req.GetId())

	if pErr != nil {
		return nil, pErr
	}

	newTitle := req.GetNewTitle()

	pgErr := s.purposeStorage.Rename(ctx, id, newTitle)
	if pgErr != nil {
		return nil, pgErr
	}

	memErr := s.purposeMemCache.Rename(ctx, id, newTitle)
	if memErr != nil {
		return nil, memErr
	}

	purpose, gErr := s.purposeMemCache.Get(ctx, id)
	if gErr != nil {
		return nil, gErr
	}

	return &pb.RenamePurposeResponse{
		Message: "Цель успешно переименована!",
		Data:    purpose.Id.String(),
	}, nil
}

func (s *Service) ChangePurposeOrderNumber(ctx context.Context,
	req *pb.ChangePurposeOrderNumberRequest) (*pb.ChangePurposeOrderNumberResponse, error) {
	id, pErr := uuid.Parse(req.GetId())
	if pErr != nil {
		return nil, pErr
	}

	purpose, gErr := s.purposeMemCache.Get(ctx, id)
	if gErr != nil {
		return nil, gErr
	}

	purposes, glErr := s.purposeMemCache.GetListByLetterId(ctx, purpose.LetterId)
	if glErr != nil {
		return nil, glErr
	}

	for _, elem := range purposes {
		if elem.Id == id {
			purpose = elem
			elem.OrderNumber = int(req.GetNewOrderNumber())
		}
	}

	purposes = utils.SortPurposes(purposes)

	for index, elem := range purposes {
		eId := elem.Id

		pgErr := s.purposeStorage.ChangeOrderNumber(ctx, eId, index+1)
		if pgErr != nil {
			return nil, pgErr
		}

		memErr := s.purposeMemCache.ChangeOrderNumber(ctx, eId, index+1)
		if memErr != nil {
			return nil, memErr
		}
	}

	return &pb.ChangePurposeOrderNumberResponse{
		Message: "",
		Data:    purpose.Id.String(),
	}, nil
}

func (s *Service) DeletePurpose(ctx context.Context,
	req *pb.DeletePurposeRequest) (*pb.DeletePurposeResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}

	pgErr := s.purposeStorage.Delete(ctx, id)
	if pgErr != nil {
		return nil, pgErr
	}

	memErr := s.purposeMemCache.Delete(ctx, id)
	if memErr != nil {
		return nil, memErr
	}

	return &pb.DeletePurposeResponse{
		Message: "Цель успешно удалена!",
	}, nil

}

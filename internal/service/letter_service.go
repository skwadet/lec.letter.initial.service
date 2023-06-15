package service

import (
	"context"
	"github.com/google/uuid"
	"time"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func (s *Service) AddLetter(ctx context.Context, req *pb.AddLetterRequest) (*pb.AddLetterResponse, error) {
	letter, mErr := mapAddLetterRequestFromProto(ctx, req)
	if mErr != nil {
		return nil, mErr
	}

	letterId := uuid.New()

	letter.Id = letterId
	letter.CreatedAt = time.Now()

	pgErr := s.letterStorage.Add(ctx, letter)
	if pgErr != nil {
		return nil, pgErr
	}

	memErr := s.letterMemCache.Add(ctx, letter)
	if memErr != nil {
		return nil, memErr
	}

	return &pb.AddLetterResponse{
		Message: "Письмо успешно добавлено!",
		Data:    letter.SubmissionId.String(),
	}, nil
}

func (s *Service) GetLetter(ctx context.Context, req *pb.GetLetterRequest) (*pb.GetLetterResponse, error) {
	id, err := uuid.Parse(req.GetSubmissionId())
	if err != nil {
		return nil, err
	}

	letter, gErr := s.letterMemCache.Get(ctx, id)
	if gErr != nil {
		return nil, gErr
	}

	mappedLetter, mErr := mapLetterToProto(ctx, letter)
	if mErr != nil {
		return nil, mErr
	}

	return &pb.GetLetterResponse{
		Message: "",
		Data:    mappedLetter,
	}, nil
}

func (s *Service) GetFullLetter(ctx context.Context, req *pb.GetFullLetterRequest) (*pb.GetFullLetterResponse, error) {
	id, err := uuid.Parse(req.GetSubmissionId())
	if err != nil {
		return nil, err
	}

	letter, gErr := s.letterMemCache.Get(ctx, id)
	if gErr != nil {
		return nil, gErr
	}

	objectives, oErr := s.objectiveMemCache.GetListByLetterId(ctx, letter.Id)
	if oErr != nil {
		return nil, oErr
	}

	purposes, pErr := s.purposeMemCache.GetListByLetterId(ctx, letter.Id)
	if pErr != nil {
		return nil, pErr
	}

	mappedLetter, mlErr := mapLetterToProto(ctx, letter)
	if mlErr != nil {
		return nil, mlErr
	}

	mappedObjectives, moErr := mapObjectiveListToProto(ctx, objectives)
	if moErr != nil {
		return nil, moErr
	}

	mappedPurposes, mpErr := mapPurposeListToProto(ctx, purposes)
	if mpErr != nil {
		return nil, mpErr
	}

	mappedResData, mrErr := mapLetterResponseDataToProto(ctx, mappedLetter, mappedObjectives, mappedPurposes)
	if mrErr != nil {
		return nil, mrErr
	}

	return &pb.GetFullLetterResponse{
		Message: "",
		Data:    mappedResData,
	}, nil
}

func (s *Service) RenameLetter(ctx context.Context, req *pb.RenameLetterRequest) (*pb.RenameLetterResponse, error) {
	sId, err := uuid.Parse(req.GetSubmissionId())
	if err != nil {
		return nil, err
	}

	letter, gErr := s.letterMemCache.Get(ctx, sId)
	if gErr != nil {
		return nil, gErr
	}

	letterId := letter.Id

	pgErr := s.letterStorage.Rename(ctx, letterId, req.GetNewInfo())
	if pgErr != nil {
		return nil, pgErr
	}

	memErr := s.letterMemCache.Rename(ctx, letter.SubmissionId, req.GetNewInfo())
	if memErr != nil {
		return nil, memErr
	}

	return &pb.RenameLetterResponse{
		Message: "Письмо успешно обновлено!",
		Data:    letter.SubmissionId.String(),
	}, nil
}

func (s *Service) DeleteLetter(ctx context.Context, req *pb.DeleteLetterRequest) (*pb.DeleteLetterResponse, error) {
	id, err := uuid.Parse(req.GetSubmissionId())
	if err != nil {
		return nil, err
	}

	pgErr := s.letterStorage.Delete(ctx, id)
	if pgErr != nil {
		return nil, pgErr
	}

	memErr := s.letterMemCache.Delete(ctx, id)
	if memErr != nil {
		return nil, memErr
	}

	return &pb.DeleteLetterResponse{
		Message: "Письмо успешно удалено!",
	}, nil
}

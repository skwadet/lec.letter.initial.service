package service

import (
	"context"
	"github.com/google/uuid"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func (s *Service) GenerateLetter(ctx context.Context, req *pb.GenerateLetterRequest) (*pb.GenerateLetterResponse, error) {
	sId, sIdErr := uuid.Parse(req.GetSubmissionId())
	if sIdErr != nil {
		return nil, sIdErr
	}

	letter, lErr := s.letterMemCache.Get(ctx, sId)
	if lErr != nil {
		return nil, lErr
	}

	docs, dErr := mapDocumentListToDomain(ctx, req.GetDocumentList())
	if dErr != nil {
		return nil, dErr
	}

	letterResearchInfo, lrErr := mapResearchInfoToDomain(ctx, req.GetResearchInfo())
	if lrErr != nil {
		return nil, lrErr
	}

	objectives, oErr := s.objectiveMemCache.GetListByLetterId(ctx, letter.Id)
	if oErr != nil {
		return nil, oErr
	}

	purposes, pErr := s.purposeMemCache.GetListByLetterId(ctx, letter.Id)
	if pErr != nil {
		return nil, pErr
	}

	letter.Documents = docs
	letter.Objectives = objectives
	letter.Purposes = purposes
	letter.LetterResearchInfo = letterResearchInfo

	mappedLetter, mlErr := mapFullLetterToReq(ctx, letter)
	if mlErr != nil {
		return nil, mlErr
	}

	gReq, grErr := mapGenerateRequest(ctx, req, mappedLetter)
	if grErr != nil {
		return nil, grErr
	}

	res, rErr := s.generatorC.Generate(ctx, gReq)
	if rErr != nil {
		return nil, rErr
	}

	return &pb.GenerateLetterResponse{
		Message: "Письмо успешно сгенерировано!",
		Data:    res.GetFile(),
	}, nil
}

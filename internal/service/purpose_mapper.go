package service

import (
	"context"
	"github.com/ecosafety/lec.letter.initial.service/internal/utils"

	gateway "github.com/ecosafety/lec.letter.initial.service/internal/pb/document_generator_api"
	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"

	"github.com/ecosafety/lec.letter.initial.service/pkg/domain"
	"github.com/google/uuid"
)

func mapAddPurposeRequestFromProto(_ context.Context, req *pb.AddPurposeRequest) (*domain.Purpose, error) {
	lId, err := uuid.Parse(req.GetLetterId())
	if err != nil {
		return nil, err
	}

	return &domain.Purpose{
		LetterId: lId,
		Title:    req.GetTitle(),
	}, nil
}

func mapPurposeListToProto(ctx context.Context, letters []*domain.Purpose) ([]*pb.PurposeData, error) {
	var mappedPurposes []*pb.PurposeData

	for _, letter := range letters {
		mappedPurpose, mErr := mapPurposeToProto(ctx, letter)
		if mErr != nil {
			return nil, mErr
		}

		mappedPurposes = append(mappedPurposes, mappedPurpose)
	}

	return mappedPurposes, nil
}

func mapPurposeToProto(_ context.Context, purpose *domain.Purpose) (*pb.PurposeData, error) {
	return &pb.PurposeData{
		Id:          purpose.Id.String(),
		OrderNumber: int32(purpose.OrderNumber),
		LetterId:    purpose.LetterId.String(),
		Title:       purpose.Title,
		CreatedAt:   utils.InnerFormatTime(purpose.CreatedAt),
	}, nil
}

func mapPurposeListToReq(ctx context.Context, letters []*domain.Purpose) ([]*gateway.PurposeData, error) {
	var mappedPurposes []*gateway.PurposeData

	for _, letter := range letters {
		mappedPurpose, mErr := mapPurposeToReq(ctx, letter)
		if mErr != nil {
			return nil, mErr
		}

		mappedPurposes = append(mappedPurposes, mappedPurpose)
	}

	return mappedPurposes, nil
}

func mapPurposeToReq(_ context.Context, purpose *domain.Purpose) (*gateway.PurposeData, error) {
	return &gateway.PurposeData{
		Id:          purpose.Id.String(),
		OrderNumber: int32(purpose.OrderNumber),
		LetterId:    purpose.LetterId.String(),
		Title:       purpose.Title,
		CreatedAt:   utils.InnerFormatTime(purpose.CreatedAt),
	}, nil
}

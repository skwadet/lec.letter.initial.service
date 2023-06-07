package service

import (
	"context"

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

func mapPurposeToProto(_ context.Context, letter *domain.Purpose) (*pb.PurposeData, error) {
	return &pb.PurposeData{
		Id:          letter.Id.String(),
		OrderNumber: int32(letter.OrderNumber),
		LetterId:    letter.LetterId.String(),
		Title:       letter.Title,
	}, nil
}

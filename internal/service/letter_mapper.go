package service

import (
	"context"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"

	"github.com/ecosafety/lec.letter.initial.service/internal/utils"
	"github.com/ecosafety/lec.letter.initial.service/pkg/domain"

	"github.com/google/uuid"
)

func mapAddLetterRequestFromProto(_ context.Context, req *pb.AddLetterRequest) (*domain.Letter, error) {
	sId, err := uuid.Parse(req.GetSubmissionId())
	if err != nil {
		return nil, err
	}

	return &domain.Letter{
		SubmissionId:   sId,
		AdditionalInfo: nil,
	}, nil
}

func mapLetterListToProto(ctx context.Context, letters []*domain.Letter) ([]*pb.LetterData, error) {
	var mappedLetters []*pb.LetterData

	for _, letter := range letters {
		mappedLetter, mErr := mapLetterToProto(ctx, letter)
		if mErr != nil {
			return nil, mErr
		}

		mappedLetters = append(mappedLetters, mappedLetter)
	}

	return mappedLetters, nil
}

func mapLetterToProto(_ context.Context, letter *domain.Letter) (*pb.LetterData, error) {
	return &pb.LetterData{
		Id:             letter.Id.String(),
		SubmissionId:   letter.SubmissionId.String(),
		CreatedAt:      utils.InnerFormatTime(letter.CreatedAt),
		AdditionalInfo: letter.AdditionalInfo,
	}, nil
}

func mapLetterResponseDataToProto(_ context.Context,
	letter *pb.LetterData, objectives []*pb.ObjectiveData, purposes []*pb.PurposeData) (*pb.GetFullLetterResponseData, error) {
	return &pb.GetFullLetterResponseData{
		Letter:     letter,
		Objectives: objectives,
	}, nil
}

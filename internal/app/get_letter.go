package app

import (
	"context"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func (i *Implementation) GetLetter(ctx context.Context, req *pb.GetLetterRequest) (*pb.GetLetterResponse, error) {
	return i.service.GetLetter(ctx, req)
}

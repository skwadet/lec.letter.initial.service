package app

import (
	"context"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func (i *Implementation) GetFullLetter(ctx context.Context,
	req *pb.GetFullLetterRequest) (*pb.GetFullLetterResponse, error) {
	return i.service.GetFullLetter(ctx, req)
}

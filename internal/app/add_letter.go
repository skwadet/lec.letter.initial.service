package app

import (
	"context"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func (i *Implementation) AddLetter(ctx context.Context, req *pb.AddLetterRequest) (*pb.AddLetterResponse, error) {
	return i.service.AddLetter(ctx, req)
}

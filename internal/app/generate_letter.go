package app

import (
	"context"
	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func (i *Implementation) GenerateLetter(ctx context.Context,
	req *pb.GenerateLetterRequest) (*pb.GenerateLetterResponse, error) {
	return i.service.GenerateLetter(ctx, req)
}

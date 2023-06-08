package app

import (
	"context"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func (i *Implementation) DeleteLetter(ctx context.Context,
	req *pb.DeleteLetterRequest) (*pb.DeleteLetterResponse, error) {
	return i.service.DeleteLetter(ctx, req)
}

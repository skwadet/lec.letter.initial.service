package app

import (
	"context"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func (i *Implementation) GetPurpose(ctx context.Context, req *pb.GetPurposeRequest) (*pb.GetPurposeResponse, error) {
	return i.service.GetPurpose(ctx, req)
}

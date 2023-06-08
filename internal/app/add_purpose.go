package app

import (
	"context"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func (i *Implementation) AddPurpose(ctx context.Context, req *pb.AddPurposeRequest) (*pb.AddPurposeResponse, error) {
	return i.service.AddPurpose(ctx, req)
}

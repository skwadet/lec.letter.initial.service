package app

import (
	"context"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func (i *Implementation) RenamePurpose(ctx context.Context,
	req *pb.RenamePurposeRequest) (*pb.RenamePurposeResponse, error) {
	return i.service.RenamePurpose(ctx, req)
}

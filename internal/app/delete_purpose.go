package app

import (
	"context"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func (i *Implementation) DeletePurpose(ctx context.Context,
	req *pb.DeletePurposeRequest) (*pb.DeletePurposeResponse, error) {
	return i.service.DeletePurpose(ctx, req)
}

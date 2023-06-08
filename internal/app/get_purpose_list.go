package app

import (
	"context"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func (i *Implementation) GetPurposeList(ctx context.Context,
	req *pb.GetPurposeListRequest) (*pb.GetPurposeListResponse, error) {
	return i.service.GetPurposeList(ctx, req)
}

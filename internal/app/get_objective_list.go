package app

import (
	"context"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func (i *Implementation) GetObjectiveList(ctx context.Context,
	req *pb.GetObjectiveListRequest) (*pb.GetObjectiveListResponse, error) {
	return i.service.GetObjectiveList(ctx, req)
}

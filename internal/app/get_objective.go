package app

import (
	"context"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func (i *Implementation) GetObjective(ctx context.Context,
	req *pb.GetObjectiveRequest) (*pb.GetObjectiveResponse, error) {
	return i.service.GetObjective(ctx, req)
}

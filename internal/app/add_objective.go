package app

import (
	"context"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func (i *Implementation) AddObjective(ctx context.Context,
	req *pb.AddObjectiveRequest) (*pb.AddObjectiveResponse, error) {
	return i.service.AddObjective(ctx, req)
}

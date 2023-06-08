package app

import (
	"context"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func (i *Implementation) DeleteObjective(ctx context.Context,
	req *pb.DeleteObjectiveRequest) (*pb.DeleteObjectiveResponse, error) {
	return i.service.DeleteObjective(ctx, req)
}

package app

import (
	"context"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func (i *Implementation) RenameObjective(ctx context.Context,
	req *pb.RenameObjectiveRequest) (*pb.RenameObjectiveResponse, error) {
	return i.service.RenameObjective(ctx, req)
}

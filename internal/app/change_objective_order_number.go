package app

import (
	"context"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func (i *Implementation) ChangeObjectiveOrderNumber(ctx context.Context,
	req *pb.ChangeObjectiveOrderNumberRequest) (*pb.ChangeObjectiveOrderNumberResponse, error) {
	return i.service.ChangeObjectiveOrderNumber(ctx, req)
}

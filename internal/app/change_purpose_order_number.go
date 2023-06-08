package app

import (
	"context"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func (i *Implementation) ChangePurposeOrderNumber(ctx context.Context,
	req *pb.ChangePurposeOrderNumberRequest) (*pb.ChangePurposeOrderNumberResponse, error) {
	return i.service.ChangePurposeOrderNumber(ctx, req)
}

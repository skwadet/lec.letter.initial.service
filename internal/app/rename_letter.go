package app

import (
	"context"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func (i *Implementation) RenameLetter(ctx context.Context,
	req *pb.RenameLetterRequest) (*pb.RenameLetterResponse, error) {
	return i.service.RenameLetter(ctx, req)
}

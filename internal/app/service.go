package app

import (
	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"

	"github.com/ecosafety/lec.letter.initial.service/internal/service"
)

type Implementation struct {
	pb.UnimplementedLetterInitialServiceServer
	service *service.Service
}

func NewLecLetterInitialService(service *service.Service) *Implementation {
	return &Implementation{
		service: service,
	}
}

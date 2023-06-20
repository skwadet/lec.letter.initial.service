package service

import (
	"context"
	gateway "github.com/ecosafety/lec.letter.initial.service/internal/pb/document_generator_api"
	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"
)

func mapGenerateRequest(ctx context.Context, req *pb.GenerateLetterRequest,
	letter *gateway.FullLetterData) (*gateway.GenerateLetterRequest, error) {

	mappedOrganizationHeader, hErr := mapOrganizationHeader(ctx, req.GetOrganizationHeader())
	if hErr != nil {
		return nil, hErr
	}

	return &gateway.GenerateLetterRequest{
		Letter:             letter,
		OrganizationHeader: mappedOrganizationHeader,
	}, nil
}

func mapOrganizationHeader(_ context.Context, req *pb.OrganizationHeaderData) (*gateway.OrganizationHeaderData, error) {
	return &gateway.OrganizationHeaderData{
		Credits: req.GetCredits(),
		LogoUrl: req.GetLogoUrl(),
	}, nil
}

package service

import (
	"context"
	"github.com/ecosafety/lec.letter.initial.service/internal/utils"
	"time"

	gateway "github.com/ecosafety/lec.letter.initial.service/internal/pb/document_generator_api"
	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"

	"github.com/ecosafety/lec.letter.initial.service/pkg/domain"
)

func mapResearchInfoToReq(_ context.Context, info *domain.LetterResearchInfo,
	createdAt time.Time) (*gateway.LetterResearchInfoData, error) {
	mCreatedAt := utils.OuterFormatDate(createdAt)

	return &gateway.LetterResearchInfoData{
		Direction:            info.GetDirection(mCreatedAt),
		Greetings:            info.GetGreeting(),
		Welcome:              info.GetWelcomePart(),
		Sponsor:              info.GetSponsor(),
		ContractOrganization: info.GetContractOrganization(),
		DrugName:             info.GetDrug(),
		Manufacturer:         info.GetManufacturer(),
		MainResearcher:       info.GetMainResearcher(),
		Sign:                 info.GetSign(),
	}, nil
}

func mapResearchInfoToDomain(_ context.Context, req *pb.ResearchInfoData) (*domain.LetterResearchInfo, error) {
	return &domain.LetterResearchInfo{
		FullProtocol:            req.GetFullProtocol(),
		Sponsor:                 req.GetSponsor(),
		Contractor:              req.GetContractor(),
		DrugName:                req.GetDrugName(),
		Manufacturer:            req.GetManufacturer(),
		MainResearcher:          req.GetMainResearcher(),
		MainResearcherInitials:  req.GetMainResearcherInitials(),
		ChairmanWithoutSurname:  req.GetChairmanWithoutSurname(),
		ChairmanInitialsTitleTo: req.GetChairmanInitialsTitleTo(),
	}, nil
}

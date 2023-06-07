package domain

import (
	"fmt"
)

type LetterResearchInfo struct {
	ContractOrganization     *string
	Sponsor                  *string
	FullProtocol             *string
	DrugName                 *string
	MainResearcherInitials   *string
	MainResearcher           *string
	CommitteeChairman        *string
	CommitteeChairmanTitleTo *string
}

func (l *LetterResearchInfo) GetGreeting() *string {
	if l.CommitteeChairman == nil {
		return nil
	}

	compliment := "Глубокоуважаемый "

	welcomePart := fmt.Sprintf("%s %s!", compliment, *l.CommitteeChairman)

	return &welcomePart
}

func (l *LetterResearchInfo) GetWelcomePart() *string {
	if l.FullProtocol == nil {
		return nil
	}

	compliment := "Свидетельствую Вам свое почтение и прошу провести" +
		"этическую экспертизу документов и одобрить проведение нового клинического исследования по протоколу "

	welcomePart := fmt.Sprintf("%s %s", compliment, *l.FullProtocol)

	return &welcomePart
}

func (l *LetterResearchInfo) GetSponsor() *string {
	if l.Sponsor == nil {
		return nil
	}

	sponsor := fmt.Sprintf("<br>Спонсор:</br> %s", *l.Sponsor)
	return &sponsor
}

func (l *LetterResearchInfo) GetContractOrganization() *string {
	if l.ContractOrganization == nil {
		return nil
	}

	contrOrg := fmt.Sprintf("<br>Контрактно-исследовательская организация:</br> %s", *l.ContractOrganization)
	return &contrOrg
}

func (l *LetterResearchInfo) GetDrug() *string {
	if l.DrugName == nil {
		return nil
	}

	drugName := fmt.Sprintf("<br>Исследуемый препарат:</br> %s", *l.DrugName)

	return &drugName
}

func (l *LetterResearchInfo) GetSign() *string {
	if l.MainResearcherInitials == nil {
		return nil
	}

	mainResearcher := fmt.Sprintf("С уважением,\n Главный исследователь\n %s", *l.MainResearcherInitials)

	return &mainResearcher
}

package domain

import (
	"fmt"
)

type LetterResearchInfo struct {
	FullProtocol            *string
	Sponsor                 *string
	Contractor              *string
	Manufacturer            *string
	DrugName                *string
	MainResearcher          *string
	MainResearcherInitials  *string
	ChairmanWithoutSurname  *string
	ChairmanInitialsTitleTo *string
}

func (l *LetterResearchInfo) GetDirection(createdAt string) *string {
	if l.ChairmanInitialsTitleTo == nil || l.MainResearcherInitials == nil {
		return nil
	}

	direction := fmt.Sprintf("%s\n Главный исследователь\n %s\n %s",
		*l.ChairmanInitialsTitleTo,
		*l.MainResearcher,
		createdAt,
	)

	return &direction
}

func (l *LetterResearchInfo) GetGreeting() *string {
	if l.ChairmanWithoutSurname == nil {
		return nil
	}

	welcomePart := fmt.Sprintf("%s %s!", "Глубокоуважаемый ", *l.ChairmanWithoutSurname)
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
	if l.Contractor == nil {
		return nil
	}

	contractor := fmt.Sprintf("<br>Контрактно-исследовательская организация:</br> %s", *l.Contractor)
	return &contractor
}

func (l *LetterResearchInfo) GetManufacturer() *string {
	if l.Manufacturer == nil {
		return nil
	}

	manufacturer := fmt.Sprintf("<br>Производитель:</br> %s", *l.Manufacturer)
	return &manufacturer
}

func (l *LetterResearchInfo) GetDrug() *string {
	if l.DrugName == nil {
		return nil
	}

	drug := fmt.Sprintf("<br>Исследуемый препарат:</br> %s", *l.DrugName)
	return &drug
}

func (l *LetterResearchInfo) GetMainResearcher() *string {
	if l.MainResearcher == nil {
		return nil
	}

	mainResearcher := fmt.Sprintf("<br>Главный исследователь:</br> %s", *l.MainResearcher)
	return &mainResearcher
}

func (l *LetterResearchInfo) GetSign() *string {
	if l.MainResearcherInitials == nil {
		return nil
	}

	sign := fmt.Sprintf("С уважением,\n Главный исследователь\n %s", *l.MainResearcherInitials)
	return &sign
}

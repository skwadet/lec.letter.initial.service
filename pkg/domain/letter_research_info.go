package domain

import (
	"fmt"
)

type LetterResearchInfo struct {
	FullProtocol            string
	Sponsor                 string
	Contractor              string
	Manufacturer            string
	DrugName                string
	MainResearcher          string
	MainResearcherInitials  string
	ChairmanWithoutSurname  string
	ChairmanInitialsTitleTo string
}

func (l *LetterResearchInfo) GetDirection(createdAt string) *string {
	if l.ChairmanInitialsTitleTo == "" || l.MainResearcherInitials == "" {
		return nil
	}

	direction := fmt.Sprintf("<p style=\"text-align: end;\">%s</p> <p style=\"text-align: end;\">", l.ChairmanInitialsTitleTo) +
		fmt.Sprintf("Главный исследователь</p> ") +
		fmt.Sprintf("<p style=\"text-align: end;\">%s</p> ", l.MainResearcher) +
		fmt.Sprintf("<p style=\"text-align: end;\">%s</p>", createdAt)

	return &direction
}

func (l *LetterResearchInfo) GetGreeting() *string {
	if l.ChairmanWithoutSurname == "" {
		return nil
	}

	welcomePart := fmt.Sprintf("%s %s!", "Глубокоуважаемый(-ая) ", l.ChairmanWithoutSurname)
	return &welcomePart
}

func (l *LetterResearchInfo) GetWelcomePart() *string {
	if l.FullProtocol == "" {
		return nil
	}

	compliment := "Свидетельствую Вам свое почтение и прошу провести " +
		"этическую экспертизу документов и одобрить проведение нового клинического исследования по протоколу "

	welcomePart := fmt.Sprintf("%s %s", compliment, l.FullProtocol)
	return &welcomePart
}

func (l *LetterResearchInfo) GetSponsor() *string {
	if l.Sponsor == "" {
		return nil
	}

	sponsor := fmt.Sprintf("<span style=\"font-weight: 700;\">Спонсор:</span> %s", l.Sponsor)
	return &sponsor
}

func (l *LetterResearchInfo) GetContractOrganization() *string {
	if l.Contractor == "" {
		return nil
	}

	contractor := fmt.Sprintf("<span style=\"font-weight: 700;\">Контрактно-исследовательская организация:</span> %s", l.Contractor)
	return &contractor
}

func (l *LetterResearchInfo) GetManufacturer() *string {
	if l.Manufacturer == "" {
		return nil
	}

	manufacturer := fmt.Sprintf("<span style=\"font-weight: 700;\">Производитель:</span> %s", l.Manufacturer)
	return &manufacturer
}

func (l *LetterResearchInfo) GetDrug() *string {
	if l.DrugName == "" {
		return nil
	}

	drug := fmt.Sprintf("<span style=\"font-weight: 700;\">Исследуемый препарат:</span> %s", l.DrugName)
	return &drug
}

func (l *LetterResearchInfo) GetMainResearcher() *string {
	if l.MainResearcher == "" {
		return nil
	}

	mainResearcher := fmt.Sprintf("<span style=\"font-weight: 700;\">Главный исследователь:</span> %s", l.MainResearcher)
	return &mainResearcher
}

func (l *LetterResearchInfo) GetSign() *string {
	if l.MainResearcherInitials == "" {
		return nil
	}

	sign := fmt.Sprintf("<p>С уважением,</p> <p>Главный исследователь</p> <p>%s</p>", l.MainResearcherInitials)
	return &sign
}

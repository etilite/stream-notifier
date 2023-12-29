package dto

type PersonalCheckResultDTO struct {
	checkResult CheckResultDTO
	chatID      string
}

func NewPersonalCheckResultDTO(cr CheckResultDTO, chatID string) PersonalCheckResultDTO {
	return PersonalCheckResultDTO{
		checkResult: cr,
		chatID:      chatID,
	}
}

func (n PersonalCheckResultDTO) CheckResult() CheckResultDTO {
	return n.checkResult
}

func (n PersonalCheckResultDTO) ChatID() string {
	return n.chatID
}

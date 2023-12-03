package domain

type CheckResultDTO struct {
	id         string
	name       string
	previewUrl string
	category   string
	title      string
	online     bool
}

func NewCheckResult(
	id string,
	name string,
	previewUrl string,
	category string,
	title string,
	online bool,
) CheckResultDTO {
	return CheckResultDTO{
		id:         id,
		name:       name,
		previewUrl: previewUrl,
		category:   category,
		title:      title,
		online:     online,
	}
}

func (cr CheckResultDTO) ID() string {
	return cr.id
}

func (cr CheckResultDTO) Name() string {
	return cr.name
}

func (cr CheckResultDTO) PreviewUrl() string {
	return cr.previewUrl
}

func (cr CheckResultDTO) Category() string {
	return cr.category
}

func (cr CheckResultDTO) Title() string {
	return cr.title
}

func (cr CheckResultDTO) IsOnline() bool {
	return cr.online
}

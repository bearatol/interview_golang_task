package service

type Service struct {
	pdf PDF
}

func NewService(pdf PDF) *Service {
	return &Service{
		pdf: pdf,
	}
}

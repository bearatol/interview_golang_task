package service

type Repository interface{}

type PriceGenerator interface{}

type Service struct {
	repo     Repository
	priceGen PriceGenerator
}

func NewService(repo Repository, priceGen PriceGenerator) *Service {
	return &Service{
		repo:     repo,
		priceGen: priceGen,
	}
}

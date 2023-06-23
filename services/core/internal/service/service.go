package service

type Repository interface{}

type Service struct {
	repo     Repository
	priceGen PriceGenerator
	authGen  AuthGenerator
}

func NewService(
	repo Repository,
	priceGen PriceGenerator,
	authGen AuthGenerator,
) *Service {
	return &Service{
		repo:     repo,
		priceGen: priceGen,
		authGen:  authGen,
	}
}

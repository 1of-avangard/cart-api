package service

type Service struct {
	repo CartRepository
}

func NewService(repo CartRepository) *Service {
	return &Service{
		repo: repo,
	}
}

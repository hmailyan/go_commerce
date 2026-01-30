package brands

import "context"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, req CreateRequest) error {

	err := s.repo.Create(ctx, req)
	if err != nil {
		return err
	}
}

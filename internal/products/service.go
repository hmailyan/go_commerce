package products

import (
	"context"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Create(ctx context.Context, req CreateRequest) (*Product, error) {

	product := &Product{
		Name:  req.Name,
		Price: req.Price,
		Image: req.Image,
	}

	err := s.repo.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Service) List(ctx context.Context) ([]*Product, error) {
	products, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*Product, error) {
	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *Service) SearchByQuery(ctx context.Context, query string) ([]*Product, error) {
	products, err := s.repo.SearchByQuery(ctx, query)
	if err != nil {
		return nil, err
	}
	return products, nil
}

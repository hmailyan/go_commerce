package carts

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddItem(ctx context.Context, req AddItemRequest, userId string) error {
	var userCart *Cart

	uid, err := uuid.Parse(userId)
	if err != nil {
		return err
	}

	pid, err := uuid.Parse(req.ProductID)
	if err != nil {
		return err
	}

	userCart, err = s.repo.GetOrCreateUserCart(ctx, uid)

	if err != nil {
		return err
	}

	item, err := s.repo.GetItem(ctx, userCart.ID, pid)
	if err != nil {
		return err
	}

	if item != nil {
		return s.repo.UpdateQty(ctx, item.ID, item.Quantity+req.Quantity)
	}

	return s.repo.AddItem(ctx, userCart.ID, pid, req.Quantity)

}

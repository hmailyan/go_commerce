package carts

import (
	"context"
	"fmt"

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

func (s *Service) GetCart(ctx context.Context, userId string) (*Cart, error) {
	uid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	cart, err := s.repo.GetOrCreateUserCart(ctx, uid)
	if err != nil {
		return nil, err
	}

	return cart, nil

}

func (s *Service) RemoveItem(ctx context.Context, req RemoveItemRequest, userId string) error {
	pid, err := uuid.Parse(req.ProductID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	uid, err := uuid.Parse(userId)
	if err != nil {
		return err
	}

	cart, err := s.repo.GetOrCreateUserCart(ctx, uid)
	if err != nil {
		return err
	}

	item, err := s.repo.GetItem(ctx, cart.ID, pid)
	if err != nil {
		return err
	}

	if item.Quantity == req.Quantity {

		err = s.repo.RemoveItem(ctx, pid, req.Quantity, cart.ID)
		if err != nil {
			return err
		}
	} else {
		qty := item.Quantity - req.Quantity

		if qty < 0 {
			return ErrQuantityMinusable
		}
		err = s.repo.UpdateQty(ctx, item.ID, qty)
		if err != nil {
			return err
		}
	}

	return nil

}

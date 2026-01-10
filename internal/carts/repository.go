package carts

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r *Repository) GetOrCreateUserCart(ctx context.Context, userId uuid.UUID) (*Cart, error) {
	var cart Cart

	err := r.db.WithContext(ctx).Where("user_id = ?", userId).First(&cart).Error
	if err == nil {
		return &cart, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {

		cart = Cart{UserID: userId}
		if err := r.db.WithContext(ctx).Create(&cart).Error; err != nil {
			return nil, err
		}
		return &cart, nil
	}

	return nil, err

}

func (r *Repository) GetItem(ctx context.Context, cartID uuid.UUID, productID uuid.UUID) (*CartItem, error) {
	var item CartItem
	err := r.db.WithContext(ctx).
		Where("cart_id = ? AND product_id = ?", cartID, productID).
		First(&item).Error

	return &item, err
}

func (r *Repository) UpdateQty(ctx context.Context, itemID uuid.UUID, qty int) error {
	return r.db.WithContext(ctx).
		Model(&CartItem{}).
		Where("id = ?", itemID).
		Update("quantity", qty).Error
}

func (r *Repository) AddItem(ctx context.Context, cartID uuid.UUID, pID uuid.UUID, qty int) error {
	item := CartItem{
		CartID:    cartID,
		ProductID: pID,
		Quantity:  qty,
	}

	err := r.db.WithContext(ctx).Create(item).Error

	return err

}

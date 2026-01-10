package carts

import (
	"context"

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

	res := r.db.WithContext(ctx).Preload("Items.Product").Where("user_id = ?", userId).Limit(1).Find(&cart)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {

		cart = Cart{UserID: userId}
		if err := r.db.WithContext(ctx).Create(&cart).Error; err != nil {
			return nil, err
		}
	}

	return &cart, nil

}

func (r *Repository) GetItem(ctx context.Context, cartID uuid.UUID, productID uuid.UUID) (*CartItem, error) {
	var item CartItem
	res := r.db.WithContext(ctx).
		Where("cart_id = ? AND product_id = ?", cartID, productID).Limit(1).Find(&item)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, nil
	}
	return &item, nil
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

	err := r.db.WithContext(ctx).Create(&item).Error

	return err

}

func (r *Repository) RemoveItem(ctx context.Context, pid uuid.UUID, qty int, cartId uuid.UUID) error {
	var item CartItem
	err := r.db.WithContext(ctx).Where("cart_id = ? AND product_id = ?", cartId, pid).Delete(&item).Error

	if err != nil {
		return err
	}

	return nil

}

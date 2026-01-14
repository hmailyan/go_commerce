package carts

import "errors"

var (
	ErrQuantityMinusable = errors.New("Quantity can't be minus")
	ErrItemNotFound      = errors.New("item not found in cart")
)

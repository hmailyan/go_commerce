package routes

import (
	"github.com/hmailyan/go_ecommerce/internal/carts"
	"github.com/hmailyan/go_ecommerce/internal/products"
	"github.com/hmailyan/go_ecommerce/internal/users"
)

type Dependencies struct {
	UserHandler    *users.Handler
	ProductHandler *products.Handler
	CartHandler    *carts.Handler
}

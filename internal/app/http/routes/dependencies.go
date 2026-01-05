package routes

import (
	"github.com/hmailyan/go_ecommerce/internal/products"
	"github.com/hmailyan/go_ecommerce/internal/shared/middleware"
	"github.com/hmailyan/go_ecommerce/internal/users"
)

type Dependencies struct {
	UserHandler    *users.Handler
	ProductHandler *products.Handler
	AuthMiddleware *middleware.Auth()
	// OrderHandler   *orders.Handler
}

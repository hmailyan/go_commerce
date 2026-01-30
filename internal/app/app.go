package app

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/hmailyan/go_ecommerce/internal/app/http/routes"
	"github.com/hmailyan/go_ecommerce/internal/shared/mailer"
	"github.com/hmailyan/go_ecommerce/internal/shared/utils"

	"github.com/hmailyan/go_ecommerce/internal/brands"
	"github.com/hmailyan/go_ecommerce/internal/carts"
	"github.com/hmailyan/go_ecommerce/internal/products"
	"github.com/hmailyan/go_ecommerce/internal/shared/cache"
	"github.com/hmailyan/go_ecommerce/internal/users"
	"gorm.io/gorm"
)

func BuildRouter(db *gorm.DB, rds *cache.Redis) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	userRepo := users.NewRepository(db)
	productRepo := products.NewRepository(db)
	cartRepo := carts.NewRepository(db)
	brandRepo := brands.NewRepository(db)

	hasher := utils.NewPasswordUtils()
	tokenGen := utils.NewTokenUtils()

	smtpMailer := mailer.NewSMTPMailer(mailer.SMTPConfig{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMTP_PASSWORD"),
		From:     os.Getenv("SMTP_FROM"),
	})

	userService := users.NewService(userRepo, hasher, tokenGen, smtpMailer)
	productService := products.NewService(productRepo)
	cartService := carts.NewService(cartRepo)
	brandService := brands.NewService(brandRepo)

	userHandler := users.NewHandler(userService)
	productHandler := products.NewHandler(productService)
	CartHandler := carts.NewHandler(cartService)
	BrandHandler := brands.NewHandler(brandService)

	deps := &routes.Dependencies{
		UserHandler:    userHandler,
		ProductHandler: productHandler,
		CartHandler:    CartHandler,
		BrandHandler:   BrandHandler,
	}

	routes.RegisterRoutes(r, deps)

	return r
}

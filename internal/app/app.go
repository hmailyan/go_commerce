package app

import (
	"github.com/gin-gonic/gin"

	"github.com/hmailyan/go_ecommerce/internal/app/http/routes"
	"github.com/hmailyan/go_ecommerce/internal/shared/utils"
	"github.com/hmailyan/go_ecommerce/internal/users"
	"gorm.io/gorm"
)

func BuildRouter(db *gorm.DB) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	userRepo := users.NewRepository(db)
	hasher := utils.NewPasswordUtils()
	tokenGen := utils.NewTokenUtils()

	userService := users.NewService(userRepo, hasher, tokenGen)
	userHandler := users.NewHandler(userService)

	deps := &routes.Dependencies{
		UserHandler: userHandler,
		// ProductHandler: productHandler,
	}

	routes.RegisterRoutes(r, deps)

	return r
}

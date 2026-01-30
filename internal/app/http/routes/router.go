package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, deps *Dependencies) {
	api := r.Group("/api/v1")

	RegisterUserRoutes(api, deps)
	RegisterProductRoutes(api, deps)
	RegisterCartRoutes(api, deps)
	RegisterBrandRoutes(api, deps)

	// routes.RegisterOrderRoutes(api)
}

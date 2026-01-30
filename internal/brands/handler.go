package brands

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
			return
		}

		h.service.Create(c, req)
	}
}

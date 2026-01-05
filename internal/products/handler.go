package products

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

		out, err := h.service.Create(c, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, out)
	}

}

func (h *Handler) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		out, err := h.service.List(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, out)
	}
}

func (h *Handler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		productID := c.Param("id")
		if productID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "product id is required"})
			return
		}

		out, err := h.service.GetByID(c, productID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, out)
	}
}

func (h *Handler) SearchByQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		search := c.Query("name")
		if search == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "search query is required"})
			return
		}

		out, err := h.service.SearchByQuery(c, search)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, out)
	}
}

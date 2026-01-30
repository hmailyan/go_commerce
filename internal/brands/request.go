package brands

type CreateRequest struct {
	Name  string `json:"name" binding:"required"`
	Price uint64 `json:"price" binding:"required"`
	Image string `json:"image" binding:"required"`
}

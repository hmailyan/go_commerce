package carts

type AddItemRequest struct {
	ProductID string `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}

type CartResponse struct {
	TotalPrice float64
	Items      []CartItemResponse
}

type CartItemResponse struct {
	ProductName string
	Price       float64
	Image       string
	Quantity    int
}

func ToCartResponse(c *Cart) *CartResponse {
	var total float64
	var Items []CartItemResponse
	for _, item := range c.Items {
		Items = append(Items, ToCartItemResponse(item))
		// Multiply quantity by the preloaded product's price
		total += float64(item.Quantity) * float64(item.Product.Price)
	}

	return &CartResponse{
		TotalPrice: total,
		Items:      Items,
	}
}

func ToCartItemResponse(ci CartItem) CartItemResponse {
	return CartItemResponse{
		ProductName: ci.Product.Name,
		Price:       float64(ci.Product.Price),
		Image:       ci.Product.Image,
		Quantity:    ci.Quantity,
	}
}

package models

type CartItem struct {
	Product  Product `json:"product"`
	Quantity int64   `json:"quantity"`
}

//it is like our orderitem

type Cart struct {
	Items             []CartItem `json:"items"`
	OriginalPrice     float64    `json:"original_price"`
	CurrentTotalPrice float64    `json:"current_total_price"`
	LoyalityMemeber   bool       `json:"loyality_member"`
	PaymentBank       string     `json:"payment_bank"`
	DiscountAmmount   float64    `json:"discount_ammount"`
}

// this is like the orderRequest

func (c *Cart) AddProduct(product Product, quantity int64) {
	cartItem := CartItem{
		Product:  product,
		Quantity: quantity,
	}
	c.Items = append(c.Items, cartItem)
	price := cartItem.Product.Price * float64(cartItem.Quantity)
	c.OriginalPrice += price
	c.CurrentTotalPrice += price
}

func (c *Cart) ApplyDiscount(d float64) {
	c.CurrentTotalPrice = c.OriginalPrice - d
	if c.CurrentTotalPrice < 0 {
		c.CurrentTotalPrice = 0
	}
	c.DiscountAmmount = d
}

func (c *Cart) CalculateTotals() {
	c.OriginalPrice = 0
	c.CurrentTotalPrice = 0

	for _, item := range c.Items {
		price := item.Product.Price * float64(item.Quantity)
		c.OriginalPrice += price
		c.CurrentTotalPrice += price
	}
}

func NewCart() *Cart {
	return &Cart{}
}

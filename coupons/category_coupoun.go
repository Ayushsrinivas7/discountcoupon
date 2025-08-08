package coupons

import (
	"discount-coupon-engine/models"
	"discount-coupon-engine/strategy"
	"fmt"
)

type CategoryCoupon struct {
	BaseCoupon
	CategoryName string
	strategy     strategy.DiscountStrategy
}

func (cc *CategoryCoupon) Name() string {
	return fmt.Sprintf("coupoun of type %v \n", cc.CategoryName)
}
func (cc *CategoryCoupon) IsApplicable(cart *models.Cart) bool {
	for _, cartItem := range cart.Items {
		if cartItem.Product.Category == cc.CategoryName {
			return true
		}
	}

	return false
}
func (cc *CategoryCoupon) IsCombinable() bool {
	return true
}
func (cc *CategoryCoupon) GetDiscount(cart *models.Cart) float64 {
	return cc.strategy.Calculate(cart.OriginalPrice)
}
func (cc *CategoryCoupon) Apply(cart *models.Cart) {
	if cc.IsApplicable(cart) {
		discount := cc.GetDiscount(cart)
		cart.ApplyDiscount(discount)
		println(cc.Name(), "applied:", discount)
		if !cc.IsCombinable() {
			return
		}
	}
	if cc.GetNext() != nil {
		cc.GetNext().Apply(cart)
	}
}

func NewCategoryCoupon(categoryName string, strategy strategy.DiscountStrategy) *CategoryCoupon {
	return &CategoryCoupon{
		CategoryName: categoryName,
		strategy:     strategy,
	}
}

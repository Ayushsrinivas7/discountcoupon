package coupons

import (
	"discount-coupon-engine/models"
	"discount-coupon-engine/strategy"
)

type LoyaltyCoupon struct {
	BaseCoupon
	strategy strategy.DiscountStrategy
}

func (l *LoyaltyCoupon) IsApplicable(cart *models.Cart) bool {
	return cart.LoyalityMemeber
}

func (l *LoyaltyCoupon) GetDiscount(cart *models.Cart) float64 {
	return l.strategy.Calculate(cart.OriginalPrice)
}

func (l *LoyaltyCoupon) Apply(cart *models.Cart) {
	if l.IsApplicable(cart) {
		discount := l.GetDiscount(cart)
		cart.DiscountAmmount += discount
		cart.ApplyDiscount(discount)
		println(l.Name(), "applied:", cart.DiscountAmmount)
		if !l.IsCombinable() {
			return
		}
	}

	if l.GetNext() != nil {
		l.GetNext().Apply(cart)
	}
}
func (l *LoyaltyCoupon) Name() string {
	return "Loyalty Coupon"
}
func (l *LoyaltyCoupon) IsCombinable() bool {
	return false
}

func NewLoyaltyCoupon(s strategy.DiscountStrategy) *LoyaltyCoupon {
	return &LoyaltyCoupon{
		strategy: s,
	}
}

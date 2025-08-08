package coupons

import (
	"discount-coupon-engine/models"
	"discount-coupon-engine/strategy"
)

type BankCoupon struct {
	BaseCoupon
	strategy strategy.DiscountStrategy
	BankName string
}

func (b *BankCoupon) IsApplicable(cart *models.Cart) bool {
	if cart.PaymentBank == b.BankName {
		return true
	}
	return false
}

func (b *BankCoupon) IsCombinable() bool {
	return false
}
func (b *BankCoupon) GetDiscount(cart *models.Cart) float64 {
	return b.strategy.Calculate(cart.OriginalPrice)
}
func (b *BankCoupon) Apply(cart *models.Cart) {
	if b.IsApplicable(cart) {
		discount := b.GetDiscount(cart)
		cart.DiscountAmmount += discount
		cart.ApplyDiscount(discount)
		println(b.Name(), "applied:", discount)
		if !b.IsCombinable() {
			return
		}
	}
	if b.GetNext() != nil {
		b.GetNext().Apply(cart)
	}
}

func (b *BankCoupon) Name() string {
	return "Bank Coupon (" + b.BankName + ")"
}

func NewBankCoupon(bankNmae string, s strategy.DiscountStrategy) *BankCoupon {
	return &BankCoupon{
		BankName: bankNmae,
		strategy: s,
	}
}

package coupons

import (
	"discount-coupon-engine/models"
)

type Coupon interface {
	IsApplicable(cart *models.Cart) bool
	GetDiscount(cart *models.Cart) float64
	Name() string
	IsCombinable() bool
	Apply(cart *models.Cart)
	SetNext(next Coupon)
	GetNext() Coupon
}

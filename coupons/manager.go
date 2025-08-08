package coupons

import (
	"discount-coupon-engine/models"
)

type CouponManager struct {
	head Coupon
}

func (cm *CouponManager) AddCoupon(c Coupon) {
	if cm.head == nil {
		cm.head = c
	} else {
		current := cm.head
		for current.GetNext() != nil {
			current = current.GetNext()
		}
		current.SetNext(c)
	}
}

func (cm *CouponManager) Apply(cart *models.Cart) {
	if cm.head != nil {
		cm.head.Apply(cart)
	}
}

func NewCouponManager() *CouponManager {
	return &CouponManager{}
}

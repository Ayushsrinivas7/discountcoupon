package coupons

type BaseCoupon struct {
	next Coupon
}

func (b *BaseCoupon) SetNext(next Coupon) {
	b.next = next
}

func (b *BaseCoupon) GetNext() Coupon {
	return b.next
}

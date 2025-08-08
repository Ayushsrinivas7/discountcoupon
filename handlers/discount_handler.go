package handlers

import (
	"discount-coupon-engine/coupons"
	"discount-coupon-engine/models"
	"discount-coupon-engine/strategy"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApplyDiscountHandler(c *gin.Context) {
	var cart models.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cart.CalculateTotals()
	loyaltyStrat := strategy.GetStrategy(strategy.Percent, 10, 0)
	bankStrat := strategy.GetStrategy(strategy.Flat, 0, 300)
	categoryStrat := strategy.GetStrategy(strategy.PercentWithCap, 20, 200)

	manager := coupons.NewCouponManager()
	manager.AddCoupon(coupons.NewLoyaltyCoupon(loyaltyStrat))
	manager.AddCoupon(coupons.NewBankCoupon("HDFC", bankStrat))
	manager.AddCoupon(coupons.NewCategoryCoupon("Fashion", categoryStrat))

	manager.Apply(&cart)

	c.JSON(http.StatusOK, cart)
}

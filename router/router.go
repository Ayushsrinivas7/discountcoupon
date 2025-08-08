package router

import (
	"discount-coupon-engine/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/apply-discount", handlers.ApplyDiscountHandler)

	return r
}

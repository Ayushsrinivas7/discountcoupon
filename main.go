package main

import (
	"discount-coupon-engine/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}

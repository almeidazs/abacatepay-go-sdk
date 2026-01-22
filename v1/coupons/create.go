package coupons

import (
	"context"

	v1 "github.com/almeidazs/go-abacate-types/v1"
)

func (c *Coupons) Create(ctx context.Context, body CreateCouponBody) (*Coupon, error) {
	var coupon Coupon

	if err := c.http.Post(ctx, v1.RouteCreateCoupon, body, &coupon); err != nil {
		return nil, err
	}

	return &coupon, nil
}

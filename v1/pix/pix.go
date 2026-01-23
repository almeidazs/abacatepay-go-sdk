package pix

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type Pix struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *Pix {
	return &Pix{http}
}

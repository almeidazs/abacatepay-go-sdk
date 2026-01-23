package billings

import "github.com/AbacatePay/abacatepay-go-sdk/internal/pkg/fetch"

type Billings struct {
	http *fetch.Fetch
}

func New(http *fetch.Fetch) *Billings {
	return &Billings{http}
}

package pix

import (
	"context"

	v1 "github.com/almeidazs/go-abacate-types/v1"
)

func (p *Pix) Simulate(ctx context.Context, id string, metadata *map[string]any) (*SimulatePaymentData, error) {
	var data SimulatePaymentData

	// TODO: Fix this later
	if err := p.http.Post(ctx, v1.RouteSimulatePayment+"?"+"id"+id, metadata, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

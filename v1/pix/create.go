package pix

import (
	"context"

	v1 "github.com/almeidazs/go-abacate-types/v1"
)

func (p *Pix) Create(ctx context.Context, body CreateQrCodePIXBody) (*CreateQrCodePIXData, error) {
	var data CreateQrCodePIXData

	if err := p.http.Post(ctx, v1.RouteCreatePIXQRCode, body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

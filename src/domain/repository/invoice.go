package repository

import (
	"context"

	"github.com/kenta1234567893/upsider-coding-test/src/domain/model"
)

type InvoiceRepository interface {
	Save(ctx context.Context, invoice *model.Invoice) error
}

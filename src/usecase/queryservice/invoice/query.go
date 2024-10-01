package invoice

import (
	"context"
)

type InvoiceQueryService interface {
	FindByCondition(ctx context.Context, condition *QueryCondition) ([]*Invoice, error)
}

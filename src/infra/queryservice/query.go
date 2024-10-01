package queryservice

import (
	"context"
	"net/http"

	"github.com/kenta1234567893/upsider-coding-test/ent"
	"github.com/kenta1234567893/upsider-coding-test/ent/company"
	"github.com/kenta1234567893/upsider-coding-test/ent/invoice"
	"github.com/kenta1234567893/upsider-coding-test/src/shared"
	querySerivce "github.com/kenta1234567893/upsider-coding-test/src/usecase/queryservice/invoice"
)

type invoiceQueryService struct {
	client *ent.Client
}

func NewInvoiceQueryService(client *ent.Client) querySerivce.InvoiceQueryService {
	return &invoiceQueryService{
		client: client,
	}
}

func (s *invoiceQueryService) FindByCondition(ctx context.Context, condition *querySerivce.QueryCondition) ([]*querySerivce.Invoice, error) {
	query := s.client.Debug().Invoice.Query()
	query.Where(
		invoice.HasCompanyWith(company.ID(condition.CompanyID)),
	)
	if len(condition.Status) > 0 {
		query = query.Where(invoice.StatusIn(condition.Status...))
	}
	if condition.PaymentDueDateStart != nil {
		query = query.Where(invoice.PaymentDueDateGTE(*condition.PaymentDueDateStart))
	}
	if condition.PaymentDueDateEnd != nil {
		query = query.Where(invoice.PaymentDueDateLTE(*condition.PaymentDueDateEnd))
	}

	records, err := query.All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, shared.NewError(http.StatusInternalServerError, "請求書の取得に失敗しました。", err)
	}

	res := []*querySerivce.Invoice{}
	for _, record := range records {
		appendData := &querySerivce.Invoice{
			ID:             record.ID,
			IssueDate:      record.IssueDate,
			PaymentAmount:  record.PaymentAmount,
			FeeRate:        record.FeeRate,
			TaxRate:        record.TaxRate,
			PaymentDueDate: record.PaymentDueDate,
			Status:         record.Status,
		}
		res = append(res, appendData)
	}

	return res, nil
}

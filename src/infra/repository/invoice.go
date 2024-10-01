package repository

import (
	"context"
	"net/http"

	"github.com/kenta1234567893/upsider-coding-test/ent"
	"github.com/kenta1234567893/upsider-coding-test/src/domain/model"
	"github.com/kenta1234567893/upsider-coding-test/src/domain/repository"
	"github.com/kenta1234567893/upsider-coding-test/src/shared"
)

type invoiceRepository struct {
	client *ent.Client
}

func NewInvoiceRepository(client *ent.Client) repository.InvoiceRepository {
	return &invoiceRepository{
		client: client,
	}
}

func (r *invoiceRepository) Save(ctx context.Context, invoice *model.Invoice) error {
	mutater := func(m *ent.InvoiceMutation) {
		m.SetCompanyID(invoice.CompanyID())
		m.SetCustomerID(invoice.CustomerID())
		m.SetIssueDate(invoice.IssueDate())
		m.SetPaymentAmount(invoice.PaymentAmount())
		m.SetFee(invoice.Fee())
		m.SetFeeRate(invoice.FeeRate())
		m.SetTax(invoice.Tax())
		m.SetTaxRate(invoice.TaxRate())
		m.SetBillingAmount(invoice.BillingAmount())
		m.SetPaymentDueDate(invoice.PaymentDueDate())
		m.SetStatus(invoice.Status().Value())
	}

	if invoice.ID() == 0 {
		creater := r.client.Invoice.Create()
		mutater(creater.Mutation())
		record, err := creater.Save(ctx)
		if err != nil {
			return shared.NewError(http.StatusInternalServerError, "請求書の更新に失敗しました。", err)
		}
		invoice.SetID(record.ID)
	} else {
		updater := r.client.Invoice.UpdateOneID(invoice.ID())
		mutater(updater.Mutation())
		err := updater.Exec(ctx)
		if err != nil {
			return shared.NewError(http.StatusInternalServerError, "請求書の更新に失敗しました。", err)
		}
	}

	return nil
}

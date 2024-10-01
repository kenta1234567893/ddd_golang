package invoice

import (
	"context"

	"github.com/kenta1234567893/upsider-coding-test/src/domain/model"
	"github.com/kenta1234567893/upsider-coding-test/src/domain/repository"
)

type CreateInvoiceUsecase interface {
	Issue(ctx context.Context, input *InputIssue) (*OutputIssue, error)
}

type createInvoiceUsecase struct {
	invoiceRepository repository.InvoiceRepository
}

func NewCreateInvoiceUsecase(invoiceRepository repository.InvoiceRepository) CreateInvoiceUsecase {
	return &createInvoiceUsecase{
		invoiceRepository: invoiceRepository,
	}
}

func (i *createInvoiceUsecase) Issue(ctx context.Context, input *InputIssue) (*OutputIssue, error) {

	model, err := model.Issue(input.CompanyID, input.CustomerID, input.PaymentAmount, input.FeeRate, input.TaxRate, input.PaymentDueDate)
	if err != nil {
		return nil, err
	}

	err = i.invoiceRepository.Save(ctx, model)
	if err != nil {
		return nil, err
	}

	return &OutputIssue{
		InvoiceID:      model.ID(),
		CustomerID:     model.CustomerID(),
		IssueDate:      model.IssueDate(),
		PaymentAmount:  model.PaymentAmount(),
		Fee:            model.Fee(),
		FeeRate:        model.FeeRate(),
		Tax:            model.Tax(),
		TaxRate:        model.TaxRate(),
		BillingAmount:  model.BillingAmount(),
		PaymentDueDate: model.PaymentDueDate(),
		Status:         model.Status().Value(),
	}, nil

}

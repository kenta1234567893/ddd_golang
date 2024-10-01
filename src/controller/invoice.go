package controller

import (
	"net/http"

	"github.com/kenta1234567893/upsider-coding-test/src/shared"
	invoiceUC "github.com/kenta1234567893/upsider-coding-test/src/usecase/invoice"
	"github.com/labstack/echo/v4"
)

type InvoiceController struct {
	createInvoiceUC invoiceUC.CreateInvoiceUsecase
}

func NewInvoiceController(createInvoiceUC invoiceUC.CreateInvoiceUsecase) *InvoiceController {
	return &InvoiceController{
		createInvoiceUC: createInvoiceUC,
	}
}

func (i *InvoiceController) Issue(c echo.Context) error {
	req := new(IssueInvoiceRequest)
	if err := c.Bind(req); err != nil {
		return shared.NewError(http.StatusBadRequest, "リクエストデータが不正です。", err)
	}
	ctx := c.Request().Context()

	input := &invoiceUC.InputIssue{
		CompanyID:      req.CompanyID,
		CustomerID:     req.CustomerID,
		PaymentAmount:  req.PaymentAmount,
		PaymentDueDate: req.PaymentDueDate,
		FeeRate:        req.FeeRate,
		TaxRate:        req.TaxRate,
	}

	result, err := i.createInvoiceUC.Issue(ctx, input)
	if err != nil {
		return err
	}

	res := &IssueInvoiceResponse{
		InvoiceID:      result.InvoiceID,
		CustomerID:     result.CustomerID,
		IssueDate:      result.IssueDate,
		PaymentAmount:  result.PaymentAmount,
		Fee:            result.Fee,
		FeeRate:        result.FeeRate,
		Tax:            result.Tax,
		TaxRate:        result.TaxRate,
		BillingAmount:  result.BillingAmount,
		PaymentDueDate: result.PaymentDueDate,
		Status:         result.Status,
	}

	return c.JSON(http.StatusOK, res)

}

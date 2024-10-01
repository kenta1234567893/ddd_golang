package controller

import (
	"net/http"

	"github.com/kenta1234567893/upsider-coding-test/src/shared"
	invoiceUC "github.com/kenta1234567893/upsider-coding-test/src/usecase/invoice"
	invoiceQuery "github.com/kenta1234567893/upsider-coding-test/src/usecase/queryservice/invoice"
	"github.com/labstack/echo/v4"
)

type InvoiceController struct {
	createInvoiceUC     invoiceUC.CreateInvoiceUsecase
	invoiceQueryService invoiceQuery.InvoiceQueryService
}

func NewInvoiceController(
	createInvoiceUC invoiceUC.CreateInvoiceUsecase,
	invoiceQueryService invoiceQuery.InvoiceQueryService) *InvoiceController {
	return &InvoiceController{
		createInvoiceUC:     createInvoiceUC,
		invoiceQueryService: invoiceQueryService,
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

func (i *InvoiceController) Search(c echo.Context) error {
	req := new(SearchInvoiceRequest)
	if err := c.Bind(req); err != nil {
		return shared.NewError(http.StatusBadRequest, "リクエストデータが不正です。", err)
	}
	ctx := c.Request().Context()

	condition := &invoiceQuery.QueryCondition{
		CompanyID:           req.CompanyID,
		Status:              req.Status,
		PaymentDueDateStart: &req.PaymentDueDateStart,
		PaymentDueDateEnd:   &req.PaymentDueDateEnd,
	}
	result, err := i.invoiceQueryService.FindByCondition(ctx, condition)
	if err != nil {
		return err
	}

	res := &SearchInvoiceResponse{}
	for _, r := range result {
		res.Invoices = append(res.Invoices, &SearchInvoiceDetail{
			InvoiceID:      r.ID,
			IssueDate:      r.IssueDate,
			PaymentAmount:  r.PaymentAmount,
			Fee:            r.Fee,
			FeeRate:        r.FeeRate,
			Tax:            r.Tax,
			TaxRate:        r.TaxRate,
			BillingAmount:  r.BillingAmount,
			PaymentDueDate: r.PaymentDueDate,
			Status:         r.Status,
		})
	}

	return c.JSON(http.StatusOK, res)
}

package model

import (
	"net/http"
	"time"

	"github.com/kenta1234567893/upsider-coding-test/src/shared"
)

type Invoice struct {
	id             uint64
	companyID      uint64
	customerID     uint64
	issueDate      time.Time
	paymentAmount  float64
	feeRate        float64
	taxRate        float64
	paymentDueDate time.Time
	status         InvoiceStatus
}

func Issue(companyID, cutomerID uint64, paymentAmount, feeRate, taxRate float64, paymentDueDate time.Time) (*Invoice, error) {
	// TODO バリデーション
	if paymentAmount < 0 {
		return nil, shared.NewError(http.StatusBadRequest, "支払い金額が不正です。", nil)
	}

	t := time.Now()
	nowDate := t.Truncate(time.Hour).Add(-time.Duration(t.Hour()) * time.Hour)
	return &Invoice{
		companyID:      companyID,
		customerID:     cutomerID,
		issueDate:      nowDate,
		paymentAmount:  paymentAmount,
		feeRate:        feeRate,
		taxRate:        taxRate,
		paymentDueDate: paymentDueDate,
		status:         InvoiceStatusPending(),
	}, nil
}

func (i *Invoice) SetID(id uint64) {
	i.id = id
}

func (i *Invoice) ID() uint64 {
	return i.id
}

func (i *Invoice) CompanyID() uint64 {
	return i.companyID
}

func (i *Invoice) CustomerID() uint64 {
	return i.customerID
}

func (i *Invoice) IssueDate() time.Time {
	return i.issueDate
}

func (i *Invoice) PaymentAmount() float64 {
	return i.paymentAmount
}

func (i *Invoice) FeeRate() float64 {
	return i.feeRate
}

func (i *Invoice) TaxRate() float64 {
	return i.taxRate
}

func (i *Invoice) PaymentDueDate() time.Time {
	return i.paymentDueDate
}

func (i *Invoice) Status() InvoiceStatus {
	return i.status
}

func (i *Invoice) Fee() float64 {
	return i.paymentAmount * i.feeRate
}

func (i *Invoice) Tax() float64 {
	return i.Fee() * i.taxRate
}

func (i *Invoice) BillingAmount() float64 {
	return i.paymentAmount + i.Fee() + i.Tax()
}

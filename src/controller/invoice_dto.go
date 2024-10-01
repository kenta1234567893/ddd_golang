package controller

import (
	"time"
)

type IssueInvoiceRequest struct {
	CompanyID      uint64    `json:"company_id"`
	CustomerID     uint64    `json:"customer_id"`
	PaymentAmount  float64   `json:"payment_amount"`
	FeeRate        float64   `json:"fee_rate"`
	TaxRate        float64   `json:"tax_rate"`
	PaymentDueDate time.Time `json:"payment_due_date"`
}

type IssueInvoiceResponse struct {
	InvoiceID      uint64    `json:"invoice_id"`
	CustomerID     uint64    `json:"customer_id"`
	IssueDate      time.Time `json:"issue_date"`
	PaymentAmount  float64   `json:"payment_amount"`
	Fee            float64   `json:"fee"`
	FeeRate        float64   `json:"fee_rate"`
	Tax            float64   `json:"tax"`
	TaxRate        float64   `json:"tax_rate"`
	BillingAmount  float64   `json:"billing_amount"`
	PaymentDueDate time.Time `json:"payment_due_date"`
	Status         string    `json:"status"`
}

type SearchInvoiceRequest struct {
	CompanyID           uint64    `query:"company_id"`
	Status              []string  `query:"status"`
	PaymentDueDateStart time.Time `query:"payment_due_date_start"`
	PaymentDueDateEnd   time.Time `query:"payment_due_date_end"`
}
type SearchInvoiceResponse struct {
	Invoices []*SearchInvoiceDetail `json:"invoices"`
}

type SearchInvoiceDetail struct {
	InvoiceID      uint64    `json:"invoice_id"`
	IssueDate      time.Time `json:"issue_date"`
	PaymentAmount  float64   `json:"payment_amount"`
	Fee            float64   `json:"fee"`
	FeeRate        float64   `json:"fee_rate"`
	Tax            float64   `json:"tax"`
	TaxRate        float64   `json:"tax_rate"`
	BillingAmount  float64   `json:"billing_amount"`
	PaymentDueDate time.Time `json:"payment_due_date"`
	Status         string    `json:"status"`
}

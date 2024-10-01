package invoice

import "time"

type InputIssue struct {
	CompanyID      uint64
	CustomerID     uint64
	PaymentAmount  float64
	PaymentDueDate time.Time
	FeeRate        float64
	TaxRate        float64
}

type OutputIssue struct {
	InvoiceID      uint64
	CustomerID     uint64
	IssueDate      time.Time
	PaymentAmount  float64
	Fee            float64
	FeeRate        float64
	Tax            float64
	TaxRate        float64
	BillingAmount  float64
	PaymentDueDate time.Time
	Status         string
}
